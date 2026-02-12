package context

import (
	"os"
	"path/filepath"
	"strings"
)

type ResolutionContext struct {
    TemplateDir string // Directory where template file if located upon
    CWD         string // Current Working Directory
	EnvVars 	map[string]string // SYSTEM environmental variables
    Args        map[string]string // User provided arguments
}

func NewResolutionContext(templatePath string, userArgs map[string]string) (ResolutionContext, error) {
    
    cwd, err := os.Getwd()
	if err != nil {
		return ResolutionContext{}, err
	}

	var templateDir string
	if templatePath != "" && templatePath != "-" { 
		absPath, err := filepath.Abs(templatePath)
		if err == nil {
			templateDir = filepath.Dir(absPath)
		}
	} else {
		templateDir = "" 
	}

    envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}

	return ResolutionContext{
		TemplateDir: templateDir,
		CWD:         cwd,
		EnvVars:     envVars,
		Args:        userArgs,
	}, nil
}