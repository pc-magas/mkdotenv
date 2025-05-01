package core

import (
	"io"
	"bytes"
	"testing"
	"fmt"
	"bufio"
	"strings"
)

func TestAppendValueToDotenvCreatesANewEnvFile(t *testing.T) {
	// Test input and expected output
	variable := "MYVAR"
	value := "MYVAL"
	expectedOutput := fmt.Sprintf("%s=%s\n", variable, value)

	var input io.Reader = nil

	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(input, writer, variable, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	writer.Flush()

	actualOutput := outputBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("expected output to be %q, but got %q", expectedOutput, actualOutput)
	}
}

func TestAppendValueToDotenvAppendsNewValueToFile(t *testing.T) {

	variable := "MYVAR"
	value := "MYVAL"

	dotenv:=`
VAR1=val
VAR2="val2"
`

	expectedOutput:=`
VAR1=val
VAR2="val2"
MYVAR=MYVAL
`
	var reader io.Reader = strings.NewReader(dotenv)

	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(reader, writer, variable, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	writer.Flush()

	actualOutput := outputBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("expected output to be %q, but got %q", expectedOutput, actualOutput)
	}
}


func TestAppendValueToDotenvReplacesNewValueToFile(t *testing.T) {

	variable := "MYVAR"
	value := "MYVAL"

	dotenv:=`
VAR1=val
VAR2="val2"
MYVAR=OLDVAL
`

	expectedOutput:=`
VAR1=val
VAR2="val2"
MYVAR=MYVAL
`

	var reader io.Reader = strings.NewReader(dotenv)

	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(reader, writer, variable, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	writer.Flush()

	actualOutput := outputBuffer.String()
	if actualOutput != expectedOutput {
		t.Errorf("expected output to be %q, but got %q", expectedOutput, actualOutput)
	}
}

func TestAppendValueToDotenvThrowsErrorUponInvalidVariableName(t *testing.T) {

	variable := "MY(V)AR"
	value := "MYVAL"

	dotenv:=`
VAR1=val
VAR2="val2"
MYVAR=OLDVAL
`
	var reader io.Reader = strings.NewReader(dotenv)
	
	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(reader, writer, variable, value)
	if err == nil {
		t.Fatalf("expected no error, got nil")
	}
}

func TestAppendValueToDotenvThrowsErrorUponVariableNameBeingSpaces(t *testing.T) {

	variable := "    "
	value := "MYVAL"

	dotenv:=`
VAR1=val
VAR2="val2"
MYVAR=OLDVAL
`
	var reader io.Reader = strings.NewReader(dotenv)
	
	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(reader, writer, variable, value)
	if err == nil {
		t.Fatalf("expected no error, got nil")
	}
}

func TestAppendValueToDotenvThrowsErrorUponVariableNameBeingEmtpyString(t *testing.T) {

	variable := "    "
	value := "MYVAL"

	dotenv:=`
VAR1=val
VAR2="val2"
MYVAR=OLDVAL
`
	var reader io.Reader = strings.NewReader(dotenv)
	
	var outputBuffer bytes.Buffer
	writer := bufio.NewWriter(&outputBuffer)

	_, err := AppendValueToDotenv(reader, writer, variable, value)
	if err == nil {
		t.Fatalf("expected no error, got nil")
	}
}
