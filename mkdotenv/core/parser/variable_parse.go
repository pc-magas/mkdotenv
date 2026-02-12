package parser

import (
	"regexp" 
	"fmt"
)

func ExtractVariableName(line string) (string, error) {
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