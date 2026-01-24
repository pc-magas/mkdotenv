package core

import (
	"testing"
	"strings"
	"bufio"
	"bytes"

	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/pc-magas/mkdotenv/core/parser"
	"github.com/pc-magas/mkdotenv/core/executor"
	"github.com/pc-magas/mkdotenv/core/context"
)

func dummyResolutionContext() ResolutionContext {
	return ResolutionContext{
		TemplateDir: "/tmp/templates",
		CWD:         "/tmp/project",
		EnvVars: map[string]string{
			"ENV":      "test",
			"LOG_LEVEL": "debug",
		},
		Args: map[string]string{
			"service": "example",
			"version": "v1.0.0",
		},
	}
}

func TestReplace_Passthrough(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := "# Hello\nFOO=bar\nBAZ=qux\n"
	var output bytes.Buffer

	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input),mockExec)
	
	mockExec.EXPECT().Execute(gomock.Any()).Times(0)

	err := m.Replace(writer, "dev",dummyResolutionContext())
	writer.Flush()
	assert.NoError(t, err)

	assert.Equal(t, input, output.String())
}

func TestReplace_IncalidMkdotencCommand_Passthrough(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `# mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	var output bytes.Buffer

	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input),mockExec)
	
	mockExec.EXPECT().Execute(gomock.Any()).Times(0)

	err := m.Replace(writer, "dev",dummyResolutionContext())
	writer.Flush()
	assert.NoError(t, err)

	assert.Equal(t, input, output.String())
}

func TestValueIsReplaced_UponExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
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

	err := m.Replace(writer, "dev",dummyResolutionContext())
	writer.Flush()

	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}

func TestValueIsReplaced_UponExecutionWithMultipleEnvironments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=default_secret
`

	// EXPECTATION
	mockExec.
		EXPECT().
		Execute(gomock.AssignableToTypeOf(&parser.MkDotenvCommand{})).
		DoAndReturn(func(cmd *parser.MkDotenvCommand) (string, error) {
			// Here you can assert details about the command
			assert.Equal(t, "prod", cmd.Environment)
			assert.Equal(t, "\"path_to_secret_prod\"", cmd.SecretPath)
			return "default_secret", nil
		}).
		Times(1)

	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input), mockExec)

	err := m.Replace(writer, "prod",dummyResolutionContext())
	writer.Flush()

	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}

func TestValueIsReplaced_UponExecutionWithDefaultEnvironment(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default1"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default1"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=default_secret
`
	mockExec.
		EXPECT().
		Execute(gomock.AssignableToTypeOf(&parser.MkDotenvCommand{})).
		DoAndReturn(func(cmd *parser.MkDotenvCommand) (string, error) {
			// Here you can assert details about the command
			assert.Equal(t, "default", cmd.Environment)
			assert.Equal(t, "\"path_to_secret_default1\"", cmd.SecretPath)
			return "default_secret", nil
		}).
		Times(1)

	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input), mockExec)

	err := m.Replace(writer, "default",dummyResolutionContext())
	writer.Flush()

	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}

func TestValueIsReplaced_UponExecutionWithWildcardEnvironment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExec := executor.NewMockExecutor(ctrl)

	input := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=old
`
	expected := `#mkdotenv(*):resolve("path_to_secret"):secret_resolver()
#mkdotenv(default):resolve("path_to_secret_default"):secret_resolver()
#mkdotenv(*):resolve("path_to_secret_wild"):secret_resolver()
#mkdotenv(prod):resolve("path_to_secret_prod"):secret_resolver()
#mkdotenv(dev):resolve("path_to_secret_dev"):secret_resolver()
#mkdotenv(test):resolve("path_to_secret"):secret_resolver()
API_KEY=default_secret
`
	mockExec.
		EXPECT().
		Execute(gomock.AssignableToTypeOf(&parser.MkDotenvCommand{})).
		DoAndReturn(func(cmd *parser.MkDotenvCommand) (string, error) {
			// Here you can assert details about the command
			assert.Equal(t, "*", cmd.Environment)
			assert.Equal(t, "\"path_to_secret_wild\"", cmd.SecretPath)
			return "default_secret", nil
		}).
		Times(1)

	var output bytes.Buffer
	writer := bufio.NewWriter(&output)

	m := NewDotEnvManipulator(strings.NewReader(input), mockExec)

	err := m.Replace(writer, "default",dummyResolutionContext())
	writer.Flush()

	assert.NoError(t, err)
	assert.Equal(t, expected, output.String())
}