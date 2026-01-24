package parser

import( 
	"regexp"
	"strings"
	"github.com/pc-magas/mkdotenv/core/context"
)

type MkDotenvCommand struct {
	Environment string
	SecretResolverType string
	SecretPath string
	UserParams map[string]string
	Item string
}

func GetArg(value string) string {
	re := regexp.MustCompile(`\$_ARG\[([\w]+)\]`)
	matches :=re.FindStringSubmatch(value)

	if(len(matches) == 0){
		return ""
	}

	return matches[1]
}

func ParseMkDotenvComment(readline string, arguments map[string]string,ctx context.ResolutionContext) *MkDotenvCommand {

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
				value:=strings.TrimSpace(pair[1])
				arg:=GetArg(value)
				val, ok := arguments[arg]
				if arg != "" && ok {
					value = val
				}
				params[strings.TrimSpace(pair[0])] = value
			}
		}
	}

	cmd := &MkDotenvCommand{
		Environment: env,
		SecretPath : secretPath,
		SecretResolverType: resolver,
		UserParams: params,
	}

	if item != "" {
		cmd.Item = item
	}

	return cmd
}
