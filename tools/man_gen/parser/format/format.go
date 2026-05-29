package format

import(
	"fmt"
	"man_gen/parser"
)

func generateBaseCommand(spec parser.CommandSpec) string {
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

func GenerateUsage(spec parser.CommandSpec) []string {

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