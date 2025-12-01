package core

import (
	"bufio"
	"io"
	"log"
	"github.com/pc-magas/mkdotenv/core/parser"
	"github.com/pc-magas/mkdotenv/secret"
)

type DotenvManipulator struct{
	template io.Reader
	logger   *log.Logger
}

type DotEnvSecretReplaceEngine  interface {
	Replace(output *bufio.Writer) error
}

func NewDotEnvManipulator(template io.Reader, logger *log.Logger) *DotenvManipulator {
	return &DotenvManipulator{
		template: template,
		logger:   logger,
	}
}

func (manipulator *DotenvManipulator) secretResolver(command *parser.MkDotenvCommand ) (string,error) {
	
	secret_val:=""

	switch command.SecretResolverType{
		case "keppassx":
			resolver:=secret.KepassXResolver(command.Params["file"],command.Params["password"])
			secret_val=resolver['']
		case "plain":
			return secret.PlaintextResolver(),nil
		default:
			return nil,"Resolver "+command.SecretResolverType+"is Not found"
	}

}

func(manipulator *DotenvManipulator) extractVariableName

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
			// TODO: Resolve Secret from command
			// TODO: write line with secret



			resolver,error := manipulator.secretResolverFactory(commandToExecute)
			
			if(error){
				return error
			}

			if(command.Item != nil){
				value = resolver.ResolveWithParam()
			} else {
				value = resolver.Resolve()
			}

			commandToExecute = nil
			continue
		}

		output.WriteString(line_to_write)
	}


	return nil
}