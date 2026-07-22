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
		spec.ValueToResolve.Value,
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
	var title_prefix string=".SH",title_suffix string="\n"
	var description_prefix string="",description_suffix string="\n"
	var content_prefix string="", content_suffix string=""


	if(html){
		title_prefix="<h2>"
		title_suffix="</h2>"
		
		content_prefix="<section>"
		content_suffix="</section>"

		description_prefix="<p>"
		description_suffix="</p>"
	}

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

func GenerateName(spec parser.CommandSpec, html bool) string {

	var section_prefix string="",section_suffix string="\n",title_prefix string=".SH",title_suffix string="\n",paragraph_prefix string="",paragraph_suffix string="";
	var b strings.Builder

	if(html){
		section_prefix="<section>"
		section_suffix="</section>"
		title_prefix="<h2>"
		title_suffix="</h2>"
		paragraph_prefix="<p>"
		paragraph_suffix="</p>"
	}

	
	b.WriteString(section_prefix)
	b.WriteString(title_prefix)
	b.WriteString("NAME")
	b.WriteString(title_suffix)
	b.WriteString(paragraph_prefix)
	b.WriteString(spec.Command)
	b.WriteString(paragraph_suffix)
	b.WriteString(section_suffix)

	return b.String()
}

func GenerateParamsDescription(spec parser.CommandSpec,html bool) string {

	if(len(spec.Params) == 0){
		return "";
	}

	var section_prefix string="",section_suffix string="\n",title_prefix string=".SH",title_suffix string="\n";
	var descrption_prefix string="",description_suffix string="";
	var argument_prefix string=".TP\n.B",argument_suffix="\n",argument_description_prefix string ="", argument_description_prefix string="\n",required_prefix string="", required_suffix string="";
	
	var b strings.Builder

	if(html){
		section_prefix="<section>"
		section_suffix="</section>"
		title_prefix="<h2>"
		title_suffix="</h2>"
		descrption_prefix="<dl class=\"desc-columns\">"
		description_suffix="</dl>"
		argument_prefix="<dt>"
		argument_suffix="</dt>"
		argument_description_prefix="<dd>"
		argument_description_prefix="</dd>"
	}

	b.WriteString(section_prefix)
	b.WriteString(title_prefix)
	b.WriteString("ARGUMENTS")
	b.WriteString(title_suffix)

	for i, param := range spec.Params {
		b.WriteString(argument_prefix)
		b.WriteString(param.Name)
		b.WriteString(argument_suffix)

		b.WriteString(argument_description_prefix)
		if(param.Required){
			b.WriteString(required_prefix)
			b.WriteString("REQUIRED")
			b.WriteString(required_suffix)
			b.WriteString(" ")
		}
		b.WriteString(param.Description)
		b.WriteString(argument_description_prefix)
	}

	b.WriteString(argument_suffix)
	b.WriteString(section_suffix)

	return b.String()
}

func GenerateFieldDescription(spec parser.CommandSpec,html bool) string {

	if(len(spec.Fields) == 0){
		return "";
	}

	var section_prefix string="",section_suffix string="\n",title_prefix string=".SH",title_suffix string="\n";
	var descrption_prefix string="",description_suffix string="";
	var argument_prefix string=".TP\n.B",argument_suffix="\n",argument_description_prefix string ="", argument_description_prefix string="\n";
	
	var b strings.Builder

	if(html){
		section_prefix="<section>"
		section_suffix="</section>"
		title_prefix="<h2>"
		title_suffix="</h2>"
		descrption_prefix="<dl class=\"desc-columns\">"
		description_suffix="</dl>"
		argument_prefix="<dt>"
		argument_suffix="</dt>"
		argument_description_prefix="<dd>"
		argument_description_prefix="</dd>"
	}

	b.WriteString(section_prefix)
	b.WriteString(title_prefix)
	b.WriteString("FIELDS")
	b.WriteString(title_suffix)

	for i, param := range spec.Fields {
		b.WriteString(argument_prefix)
		b.WriteString(param.Name)
		b.WriteString(argument_suffix)

		b.WriteString(argument_description_prefix)
		b.WriteString(param.Description)
		b.WriteString(argument_description_prefix)
	}

	b.WriteString(argument_suffix)
	b.WriteString(section_suffix)

	return b.String()
}