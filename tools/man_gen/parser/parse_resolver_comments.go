package parser

import (
	"fmt"
	"strings"
)

// CommandSpec represents a parsed @command block.
type CommandSpec struct {
	// Command name, e.g. "keepassx"
	Command string

	// Short description from @short-description
	ShortDescription string

	// Full description block from @description
	Description string

	// Resolution path from @resolves (e.g. parent/child/entry)
	ResolvesPath string

	// Input parameters (@param)
	Params map[string]Param

	// Available entry fields (@field)
	Fields map[string]Field
}

// Param represents a @param definition.
type Param struct {
	Name        string
	Description string
	DummyVal string
	Required    bool
}

// Field represents a @field definition.
type Field struct {
	Name        string
	Description string
}

func generateBaseCommand(spec CommandSpec) string {
	var b strings.Builder

	fmt.Fprintf(
		&b,
		"mkdotenv()::resolve(%s)::%s(",
		spec.ResolvesPath,
		spec.Command,
	)

	// append params
	for i, param := range spec.Params {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(param.Name)
		b.WriteString("=")
		b.WriteString(param.DummyVal)
	}

	b.WriteString(")")

	return b.String()
}

func GenerateUsage(spec CommandSpec) []string {

	baseCommand := generateBaseCommand(spec)

	if len(spec.Fields) == 0 {
		return []string{baseCommand}
	}

	var results []string

	for _, field := range spec.Fields {
		baseCommand := fmt.Sprintf("%s.%s", base, field.Name)
		results = append(results, line)
	}

	return results
}