package format

import(
	"fmt"
	"man_gen/parser"
)

func generateBaseCommand(spec parser.CommandSpec, prefix string, suffix string) string {
	var b strings.Builder
	
	fmt.Fprintf(
		&b,
		"mkdotenv()::resolve(%s%s%s)::%s(",
		prefix,
		spec.ResolvesPath,
		suffix,
		spec.Command,
	)

	// append params
	for i, param := range spec.Params {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(prefix)
		b.WriteString(param.Name)
		b.WriteString(suffix)
		b.WriteString("=")
		b.WriteString(param.DummyVal)
	}

	b.WriteString(")")

	return b.String()
}

func GenerateUsage(spec parser.CommandSpec, html bool) []string {

	var prefix string="",suffix string=""
	
	if(html){
		prefix="<var>"
		suffix="</var>"
	}

	baseCommand := generateBaseCommand(spec,prefix,suffix)

	if len(spec.Fields) == 0 {
		return []string{baseCommand}
	}

	var results []string

	for _, field := range spec.Fields {
		baseCommand := fmt.Sprintf("%s.%s%s%s", base, prefix, field.Name, suffix)
		results = append(results, line)
	}

	return results
}

func GenerateDescription(spec parser.CommandSpec, html bool) string {
	var title_prefix string=".SH",title_suffix string=""
	var content_prefix string="", content_suffix string=""
	var description_prefix string="",description_suffix string=""

	var b strings.Builder

	b.WriteString(content_prefix)
	b.WriteString(title_prefix)
	b.WriteString(DESCRIPTION)
	b.WriteString(title_suffix)
	b.WriteString(description_prefix)
	b.WriteString(spec.Description)
	b.WriteString(description_suffix)
	b.WriteString(content_suffix)

	return b.String()
}