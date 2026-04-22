package man

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"strings"
	"slices"
	"github.com/pc-magas/mkdotenv/params"
	"github.com/pc-magas/mkdotenv/params/parser"
)

func GetVersion(version_file string) string {
	file, err := os.Open(version_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return firstLine
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ""
}

func GenerateOptionExplanation(meta parser.FlagMeta ) string {

	explanation := "\n.TP\n.B "

	explanation+=fmt.Sprintf("--%s",meta.Name)

	if(meta.Short != ""){
		explanation+=fmt.Sprintf(",\\fI-%s\\fR",meta.Short)
	}

	for _,alias := range meta.Aliases {
		explanation+= fmt.Sprintf(" ,\\fI--%s\\fR",alias)
	}

	explanation+="\n"
	explanation+=meta.Usage
	return explanation
}

func GenerateSynopsisPart(meta parser.FlagMeta ) string {
	synopsis_part:=fmt.Sprintf("\\fI--%s\\fR",meta.Name)
			
	if(meta.Short != ""){
		synopsis_part+=fmt.Sprintf("| \\fI-%s\\fR",meta.Short)
	}

	for _,alias := range meta.Aliases {
		synopsis_part+= fmt.Sprintf(" | \\fI--%s\\fR",alias)
	}

	if(meta.Type != parser.NoValType ){
		synopsis_part+=" <\\fI"+strings.ToUpper(meta.Name)+"\\fR>"
	}

	if !meta.Required {
		synopsis_part="["+synopsis_part+"]"	
	}

	return synopsis_part
}

func MakeManpage(man_file string,version_file string) {
	version := GetVersion(version_file)
	year, month, _ := time.Now().Date()
	
	file, _ := os.Create(man_file)
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString(fmt.Sprintf(".TH MKDOTENV 1 \"%s %d\" \"mkdotenv %s\"\n",month,year,version))
	writer.WriteString(".SH NAME\nmkdotenv \\- A command-line tool that populates secrets upon a .env file from a template.\n")

	writer.WriteString(".SH DESCRIPTION\nThe \\fmkdotenv\\fR command allows users to poopulate environmental variables from a template file by placing appropriate markup upon the file.")

	
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

	synopsis_build.WriteString(".SH SYNOPSIS\n.B mkdotenv\n")
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

	writer.WriteString("\n")
	writer.WriteString(synopsis_build.String())
	writer.WriteString("\n")

	required:=required_build.String()
	
	if(required!=""){
		writer.WriteString(".SH REQUIRED ARGUMENTS\n")
		writer.WriteString(required)
		writer.WriteString("\n")
	}

	optional:=optional_build.String()

	if(optional!=""){
		writer.WriteString(".SH Optional ARGUMENTS\n")
		writer.WriteString(optional)
		writer.WriteString("\n")
	}

	writer.WriteString(".SH AUTHOR\nWritten by Desyllas Dimitrios.\n\n.SH BUGS\nReport issues at https://github.com/pc-magas/mkdotenv/issues\n\n.SH SEE ALSO\n.BR dotenv (1)\n")

	writer.Flush() 
}