package core

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strings"
)
import "fmt"

type DotenvManipulator struct{
	template io.Reader
	logger   *log.Logger
}

type MkDotenvCommand struct {
	Environment string
	SecretResolverType string
	Params map[string]string
	Item string
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

func ParseMkDotenvComment(readline string) (*MkDotenvCommand,error) {

	re := regexp.MustCompile(
		`^#mkdotenv\(([^)]*)\)::([a-zA-Z0-9_]+)\(([^)]*)\)(?:\.([A-Za-z0-9_]+))?$`,
	)
	matches := re.FindStringSubmatch(readline)

	if len(matches) == 0 {
		return nil, nil // not a match
	}

	env := matches[1]
	resolver := matches[2]
	argString := matches[3]
	item := matches[4]

	// We assume empty environment is named default
	if(env == ""){
		env="default"
	}

	params := make(map[string]string)
	if argString != "" {
		for _, kv := range strings.Split(argString, ",") {
			pair := strings.SplitN(kv, "=", 2)
			if len(pair) == 2 {
				params[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
			}
		}
	}

	cmd := &MkDotenvCommand{
		Environment:        env,
		SecretResolverType: resolver,
		Params:             params,
	}

	// Optionally store the item if you want:
	if item != "" {
		cmd.Item = item
	}

	return cmd, nil
}

func (manipulator *DotenvManipulator) Replace(output *bufio.Writer, environtment string) error {
	
	scanner := bufio.NewScanner(manipulator.template)

	// lastMatch:= nil
	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line
		
		fmt.Print(line_to_write)
	
	}


	return nil
}