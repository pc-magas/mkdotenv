package core

import (
	"testing"
	"strings"
	"bufio"
	"bytes"
	"io"
	"github.com/stretchr/testify/assert"
	"github.com/pc-magas/mkdotenv/core/parser"
)

type fakeExecutor struct {
	value string
	err   error
}

func (f *fakeExecutor) Execute(cmd *parser.MkDotenvCommand) (string, error) {

	switch(cmd.Environment){
		case "prod":
			return "prod_secret"
		case "dev":
			return "dev_secret"
		default:
			return "default_secret"
	}

	return f.value, f.err
}

func TestReplace_Passthrough(t *testing.T) {
	input := "FOO=bar\nBAZ=qux\n"
	var output bytes.Buffer

	m := NewDotEnvManipulator(strings.NewReader(input))
	m.executor = &fakeExecutor{}

	err := m.Replace(bufio.NewWriter(&output), "dev")
	assert.NoError(t, err)

	assert.Equal(t, input, output.String())
}

func TestValueIsReplaced_UponExecution(t *testing.T){
		input := `
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=default_secret
`
	var output bytes.Buffer

	m := NewDotEnvManipulator(strings.NewReader(input), log.Default())
	m.executor = &fakeExecutor{
		value:"key",
		error:nil,
	}

	err := m.Replace(bufio.NewWriter(&output), "dev")
	assert.NoError(t, err)

	assert.Equal(t, input, output.String())
}