package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"path"
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

	explanation = ".TP.B "

	explanation+=fmt.Sprintf("--%s",meta.Name)

	

}

func GenerateSynopsisPart(meta parser.FlagMeta ) string {
	synopsis_part:=fmt.Sprintf("\\fI--%s\\fR",meta.Name)
			
	if(meta.Short != ""){
		synopsis_part+=fmt.Sprintf("|\\fI-%s\\fR",meta.Short)
	}

	for _,alias := range meta.Aliases {
		synopsis_part+= fmt.Sprintf(" |\\fI--%s\\fR",alias)
	}

	if !meta.Required {
		synopsis_part="["+synopsis_part+"]"	
	}

	return synopsis_part
}

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get caller info")
	}

	man_file := path.Join(filepath.Dir(filename),"..","..","man","mkdotenv.1")
	version_file := path.Join(filepath.Dir(filename),"..","..","VERSION")

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
	// var required_build strings.Builder
	// var optional_build strings.Builder

	synopsis_build.WriteString(".SH SYNOPSIS\n.B mkdotenv\n")
	for _, order := range orders {
		flags := groups[order]
		for _, meta := range flags {
			synopsis_part:=GenerateSynopsisPart(meta)
			synopsis_build.WriteString(synopsis_part)
		}
		synopsis_build.WriteString("\n")
	}

	writer.WriteString(synopsis_build.String())
	writer.WriteString(".SH AUTHOR\nWritten by Desyllas Dimitrios.\n\n.SH BUGS\nReport issues at https://github.com/pc-magas/mkdotenv/issues\n\n.SH SEE ALSO\n.BR dotenv (1)\n")

	writer.Flush() 
}
