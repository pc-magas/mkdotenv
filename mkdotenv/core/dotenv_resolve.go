package core

import (
	"bufio"
	"io"
	"fmt"
	"regexp"
	"github.com/pc-magas/mkdotenv/core/parser"
	"github.com/pc-magas/mkdotenv/core/executor"
	"github.com/pc-magas/mkdotenv/core/context"
)

type DotenvManipulator struct{
	template io.Reader
	executor executor.Executor
}

type DotEnvSecretReplaceEngine  interface {
	Replace(output *bufio.Writer,environment string,arguments map[string]string) error
}

func NewDotEnvManipulator(template io.Reader, commandExecutor executor.Executor) *DotenvManipulator {
	return &DotenvManipulator{
		template: template,
		executor: commandExecutor,
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

func (manipulator *DotenvManipulator) Replace(output *bufio.Writer, environment string, ctx context.ResolutionContext) error {
	
	scanner := bufio.NewScanner(manipulator.template)

	var commandToExecute *parser.MkDotenvCommand = nil

	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line
		
		command := parser.ParseMkDotenvComment(line_to_write,ctx)

		if(command != nil){
			if(command.Environment == "*" || command.Environment == environment){
				commandToExecute=command
			}
			output.WriteString(line_to_write)
			output.WriteString("\n")
			continue;
		}

		if(commandToExecute != nil){

			variable,err:=parser.ExtractVariableName(line_to_write)
			if(err!=nil){
				return err
			}

			value,err := manipulator.executor.Execute(commandToExecute,ctx)

			// Unsure if return err
			if(err!=nil){
				return err
			}

			line_to_write = fmt.Sprintf("%s=%s",variable,value)
			
			commandToExecute = nil
		}

		output.WriteString(line_to_write)
		output.WriteString("\n")
	}


	return nil
}