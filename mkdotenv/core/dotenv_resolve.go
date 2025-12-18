package core

import (
	"bufio"
	"io"
	"fmt"
	"regexp"
	"github.com/pc-magas/mkdotenv/core/parser"
	"github.com/pc-magas/mkdotenv/core/executor"
)

type DotenvManipulator struct{
	template io.Reader
	executor executor.Executor
}

type DotEnvSecretReplaceEngine  interface {
	Replace(output *bufio.Writer) error
}

func NewDotEnvManipulator(template io.Reader,) *DotenvManipulator {
	return &DotenvManipulator{
		template: template,
	}
}

func (manipulator *DotenvManipulator) extractVariableName(line string) (string, error) {
    re, err := regexp.Compile(`^\s*([A-Za-z_][A-Za-z0-9_]*)\s*=.*`)
    if err != nil {
        return "", err
    }

    matches := re.FindStringSubmatch(line)
    if len(matches) < 2 {
        return "", fmt.Errorf("no variable found in line: %q", line)
    }

    return matches[1], nil
}

func (manipulator *DotenvManipulator) Replace(output *bufio.Writer, environtment string) error {
	
	scanner := bufio.NewScanner(manipulator.template)

	var commandToExecute *parser.MkDotenvCommand = nil

	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line
		
		command := parser.ParseMkDotenvComment(line_to_write)

		if(command != nil){
			if(command.Environment == "*" || command.Environment == environtment){
				commandToExecute=command
			}
			output.WriteString(line_to_write)
			continue;
		}

		if(commandToExecute != nil){
			resolver,error := manipulator.executor.Execute(commandToExecute)
			
			if(error){
				return error
			}

			variable,error:=parser.ExtractVariableName(line_to_write)
			value = manipulator.secretResolve(commandToExecute)

			line_to_write = fmt.Sprintf("%s=%s",variable,value)
			
			commandToExecute = nil
			continue
		}

		output.WriteString(line_to_write)
	}


	return nil
}