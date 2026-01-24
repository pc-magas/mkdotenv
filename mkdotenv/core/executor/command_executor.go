package executor

import (
	"fmt"
	"github.com/pc-magas/mkdotenv/core/context"
	"github.com/pc-magas/mkdotenv/core/parser"
	"github.com/pc-magas/mkdotenv/secret"
)

type Executor interface {
	Execute(command *parser.MkDotenvCommand, ctx context.ResolutionContext) (string, error)
}

type CommandExecutor struct{}

func NewExecutor() Executor {
	return &CommandExecutor{}
}

func (executer *CommandExecutor) Execute(command *parser.MkDotenvCommand, ctx context.ResolutionContext) (string, error) {

	var resolver secret.Resolver
	var err error

	switch command.SecretResolverType {
	case "keppassx":
		path := context.NewContextPath(command.Params["file"], ctx.TemplateDir)
		resolver, err = secret.NewKeepassXResolver(path, command.Params["password"])
	case "plain":
		resolver = secret.NewPlaintextResolver()
	default:
		return "", fmt.Errorf("resolver %s not found", command.SecretResolverType)
	}

	if err != nil {
		return "", err
	}

	if command.Item != "" {
		return resolver.ResolveWithParam(command.SecretPath, command.Item)
	}

	return resolver.Resolve(command.SecretPath)
}
