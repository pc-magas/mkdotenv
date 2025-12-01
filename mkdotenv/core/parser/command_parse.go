package parser

import( 
	"regexp"
	"strings"
)

type MkDotenvCommand struct {
	Environment string
	SecretResolverType string
	SecretPath string
	Params map[string]string
	Item string
}

func ParseMkDotenvComment(readline string) (*MkDotenvCommand) {

	re := regexp.MustCompile(
		`^#mkdotenv\(([^)]*)\):resolve\(([^)]*)\):([A-Za-z0-9_]+)\(([^)]*)\)(?:\.([A-Za-z0-9_]+))?$`,
	)
	matches := re.FindStringSubmatch(readline)

	if len(matches) == 0 {
		return nil
	}

	env := matches[1]
	secretPath:=matches[2]
	resolver := matches[3]
	argString := matches[4]
	item := matches[5]

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
		SecretPath : secretPath,
		SecretResolverType: resolver,
		Params:             params,
	}

	// Optionally store the item if you want:
	if item != "" {
		cmd.Item = item
	}

	return cmd
}