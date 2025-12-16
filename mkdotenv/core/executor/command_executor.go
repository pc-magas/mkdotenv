package executor

import(
	"github.com/pc-magas/mkdotenv/core/parser"
)

type Executor interface {
	Execute(command *parser.MkDotenvCommand ) (string,error)
}

type Executor stuct {
}

func NewExecuter()(*Executor,error){
	return &Executor{}
}

func (executer *Executor) Execute(command *parser.MkDotenvCommand ) (string,error) {
	
	resolver:=nil
	switch command.SecretResolverType{
		case "keppassx":
			resolver = secret.KepassXResolver(command.Params["file"],command.Params["password"])
		case "plain":
			resolver = secret.PlaintextResolver(),nil
		default:
			return nil,fmt.Errorf("resolver %s not found", command.SecretResolverType)
	}

	if(command.Item != nil){
		return resolver.ResolveWithParam(command.SecretPath,command.Item)
	} 
		
	return resolver.Resolve(command.SecretPath)
}