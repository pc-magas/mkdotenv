package parser

import( 
	"regexp"
	"strings"
)

type MkDotenvCommand struct {
	Environment string
	SecretResolverType string
	Params map[string]string
	Item string
}

func ParseMkDotenvComment(readline string) (*MkDotenvCommand) {

	re := regexp.MustCompile(
		`^#mkdotenv\(([^)]*)\)::([a-zA-Z0-9_]+)\(([^)]*)\)(?:\.([A-Za-z0-9_]+))?$`,
	)
	matches := re.FindStringSubmatch(readline)

	if len(matches) == 0 {
		return nil
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

	return cmd
}