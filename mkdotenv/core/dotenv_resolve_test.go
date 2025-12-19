package core

import (
	"testing"
	"strings"
	"bufio"
	"bytes"
	
	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/pc-magas/mkdotenv/core/parser"
)

func TestReplace_Passthrough(t *testing.T) {
	ctrl := gomock.NewController(t)

  	m := NewMockExecutor(ctrl)

	input := "FOO=bar\nBAZ=qux\n"
	var output bytes.Buffer

	m := NewDotEnvManipulator(strings.NewReader(input))

	err := m.Replace(bufio.NewWriter(&output), "dev")
	assert.NoError(t, err)

	assert.Equal(t, input, output.String())
}

func TestValueIsReplaced_UponExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=default_secret
`

	// EXPECTATION
	mockExec.
		EXPECT().
		Execute(gomock.AssignableToTypeOf(&parser.MkDotenvCommand{})).
		Return("default_secret", nil).
		Times(1)

	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input), mockExec)

	err := m.Replace(writer, "dev")
	writer.Flush()

	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}