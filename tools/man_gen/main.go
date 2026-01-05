package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pc-magas/mkdotenv/params"
)

func main() {
	out := flag.String("o", "docs/mkdotenv.1", "Output manpage file")
	version := flag.String("v", "0.4.3", "Application version")
	flag.Parse()

	f, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := f

	// Header
	fmt.Fprintf(w, ".TH MKDOTENV 1 \"%s\" \"mkdotenv %s\"\n",
	time.Now().Format("January 2006"), *version)

	// NAME
	fmt.Fprintln(w, ".SH NAME")
	fmt.Fprintln(w, "mkdotenv \\- A command-line tool to add or update environment variables in a .env file.")

	// SYNOPSIS
	fmt.Fprintln(w, ".SH SYNOPSIS")
	fmt.Fprintln(w, ".B mkdotenv")
    
	for _, flagMeta := range params.GetFlagsMeta() {
		line := "\\fI--" + flagMeta.Name + "\\fR"
		if flagMeta.Type != "" {
			line = line+" <fI" + flagMeta.Type + "fR>"
		}
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)

	// DESCRIPTION
	fmt.Fprintln(w, ".SH DESCRIPTION")
	fmt.Fprintln(w, "The \\fBmkdotenv\\fR command allows users to add or update environment variables in a .env file.")
	fmt.Fprintln(w, "By default, it modifies the current \\fB.env\\fR file. You can optionally specify input and output files.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "It supports removing duplicate variable declarations while preserving the first occurrence.")

	// PIPE SUPPORT
	fmt.Fprintln(w, ".SH PIPE SUPPORT")
	fmt.Fprintln(w, "To apply multiple updates or transformations sequentially, \\fBmkdotenv\\fR supports reading from standard input and writing to standard output.")
	fmt.Fprintln(w, "Use \\fB--output-file=-\\fR to stream to stdout, and pipe the result into another \\fBmkdotenv\\fR invocation. This enables flexible, chainable modification workflows.")

	// REQUIRED ARGUMENTS
	fmt.Fprintln(w, ".SH REQUIRED ARGUMENTS")
	for _, flagMeta := range params.GetFlagsMeta() {
		if flagMeta.Required {
			fmt.Fprintln(w, ".TP")
			fmt.Fprintf(w, ".B --%s, -%s\n", flagMeta.Name, flagMeta.Name)
			fmt.Fprintf(w, "%s\n", flagMeta.Usage)
		}
	}

	// OPTIONAL OPTIONS
	fmt.Fprintln(w, ".SH OPTIONAL OPTIONS")
	
	for _, flagMeta := range  params.GetFlagsMeta() {
		if !flagMeta.Required {
			fmt.Fprintln(w, ".TP")
			fmt.Fprintf(w, ".B --%s", flagMeta.Name)
			if flagMeta.Type != "" {
				fmt.Fprintf(w, " <%s>", flagMeta.Type)
			}
			fmt.Fprintln(w, "")
			fmt.Fprintf(w, "%s\n", flagMeta.Usage)
		}
	}

	// EXAMPLES
	fmt.Fprintln(w, ".SH EXAMPLES")
	fmt.Fprintln(w, ".TP")
	fmt.Fprintln(w, "Add a variable to the default .env file:")
	fmt.Fprintln(w, ".RS")
	fmt.Fprintln(w, "$ mkdotenv --variable-name API_KEY --variable-value 123456")
	fmt.Fprintln(w, ".RE")

	fmt.Fprintln(w, ".TP")
	fmt.Fprintln(w, "Add a variable to a specific file:")
	fmt.Fprintln(w, ".RS")
	fmt.Fprintln(w, "$ mkdotenv --variable-name API_KEY --variable-value 123456 --env-file config.env")
	fmt.Fprintln(w, ".RE")

	fmt.Fprintln(w, ".TP")
	fmt.Fprintln(w, "Write the result to a different file:")
	fmt.Fprintln(w, ".RS")
	fmt.Fprintln(w, "$ mkdotenv --variable-name API_KEY --variable-value 123456 --output-file output.env")
	fmt.Fprintln(w, ".RE")

	fmt.Fprintln(w, ".TP")
	fmt.Fprintln(w, "Remove duplicates while writing:")
	fmt.Fprintln(w, ".RS")
	fmt.Fprintln(w, "$ mkdotenv --variable-name API_KEY --variable-value 123456 --remove-doubles")
	fmt.Fprintln(w, ".RE")

	fmt.Fprintln(w, ".TP")
	fmt.Fprintln(w, "Chain multiple updates using pipes (stream between commands with \\fB--output-file=-\\fR):")
	fmt.Fprintln(w, ".RS")
	fmt.Fprintln(w, ".nf")
	fmt.Fprintln(w, "$ mkdotenv --variable-name DB_HOST --variable-value 127.0.0.1 --output-file=- \\")
	fmt.Fprintln(w, "  | mkdotenv --variable-name DB_USER --variable-value maiuser --output-file=- \\")
	fmt.Fprintln(w, "  | mkdotenv --variable-name DB_PASSWORD --variable-value XXXX --output-file=.env.production")
	fmt.Fprintln(w, ".fi")
	fmt.Fprintln(w, ".RE")

	// AUTHOR
	fmt.Fprintln(w, ".SH AUTHOR")
	fmt.Fprintln(w, "Written by Desyllas Dimitrios.")

	// BUGS
	fmt.Fprintln(w, ".SH BUGS")
	fmt.Fprintln(w, "Report issues at https://github.com/pc-magas/mkdotenv/issues")

	// SEE ALSO
	fmt.Fprintln(w, ".SH SEE ALSO")
	fmt.Fprintln(w, ".BR dotenv (1)")
}
