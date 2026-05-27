package html

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"slices"
	"github.com/pc-magas/mkdotenv/params"
	"github.com/pc-magas/mkdotenv/params/parser"
)


func GenerateOptionExplanation(meta parser.FlagMeta ) string {

	explanation := "<dt>"

	explanation+=fmt.Sprintf("<code>--%s</code>",meta.Name)

	if(meta.Short != ""){
		explanation+=fmt.Sprintf(",<code>-%s</code>",meta.Short)
	}

	for _,alias := range meta.Aliases {
		explanation+= fmt.Sprintf(" ,<code>--%s</code>",alias)
	}

	explanation+="</dt>"
	explanation+="<dd>"+meta.Usage+"</dd>"
	
	return explanation
}

func GenerateSynopsisPart(meta parser.FlagMeta ) string {
	synopsis_part:=fmt.Sprintf("<var>--%s",meta.Name)
			
	if(meta.Short != ""){
		synopsis_part+=fmt.Sprintf("</var> | <var>-%s",meta.Short)
	}

	for _,alias := range meta.Aliases {
		synopsis_part+= fmt.Sprintf("</var> | <var>--%s",alias)
	}

	if(meta.Type != parser.NoValType ){
		synopsis_part+="&lt;"+strings.ToUpper(meta.Name)+"&gt;"
	}

	synopsis_part+="</var>";

	if !meta.Required {
		synopsis_part="["+synopsis_part+"]"	
	}

	return synopsis_part
}

func MakeHtml(man_file string) {
	
	file, _ := os.Create(man_file)
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("<article class=\"manpage\"><h1>mkdotenv</h1>")
	writer.WriteString("<section><h2>NAME</h2><p>mkdotenv - A command-line tool that populates secrets upon a .env file from a template.\n<p></section>")

	writer.WriteString("<section><h2>DESCRIPTION</h2><p>\nThe <strong>mkdotenv</strong> command allows users to populate environmental variables from a template file by placing appropriate markup upon the file.</p></section>")

	
	groups := make(map[int][]parser.FlagMeta)
	orders := []int{}

	for _, meta := range params.GetFlagsMeta() {
		order := meta.Order
		groups[order] = append(groups[order], meta)

		if !slices.Contains(orders, order) {
			orders = append(orders, order)
		}
	}

	slices.Sort(orders)

	var synopsis_build strings.Builder
	var required_build strings.Builder
	var optional_build strings.Builder

	synopsis_build.WriteString("\nmkdotenv\n")
	for _, order := range orders {
		flags := groups[order]
		for _, meta := range flags {
			synopsis_part:=GenerateSynopsisPart(meta)
			synopsis_build.WriteString(synopsis_part)
			explanation := GenerateOptionExplanation(meta)

			if(meta.Required){
				required_build.WriteString(explanation)
			} else {
				optional_build.WriteString(explanation)
			}
		}
		synopsis_build.WriteString("\n")
	}

	writer.WriteString("<section><h1>SYNOPSIS</h1><pre class=\"synopsis\">")
	writer.WriteString(synopsis_build.String())
	writer.WriteString("</pre></section>")

	required:=required_build.String()
	
	if(required!=""){
		writer.WriteString("<section>")
		writer.WriteString("<h2>REQUIRED ARGUMENTS<h2>")
		writer.WriteString("<dl>")
		writer.WriteString(required)
		writer.WriteString("</dl>")
		writer.WriteString("</section>")
	}

	optional:=optional_build.String()

	if(optional!=""){
		writer.WriteString("<section>")
		writer.WriteString("<h2>Optional ARGUMENTS</h2>")
		writer.WriteString(optional)
		writer.WriteString("</section>")
	}

	writer.WriteString("<section><h1>AUTHOR</h1><p>Written by Desyllas Dimitrios</p></section>")
	writer.WriteString("<section><h1>BUGS</h1><p>Report issues at <a href=\"https://github.com/pc-magas/mkdotenv/issues\">https://github.com/pc-magas/mkdotenv/issues</a></p></section>")
	writer.WriteString("</article>")

	writer.Flush() 
}