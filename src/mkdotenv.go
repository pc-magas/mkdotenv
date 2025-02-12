/**
*   MkDotenv the .env file manipulator
*    Copyright (C) 2024 Desyllas Dimitrios
*
*   This program is free software: you can redistribute it and/or modify
*   it under the terms of the GNU General Public License as published by
*   the Free Software Foundation, either version 3 of the License, or
*   (at your option) any later version.
*
*   This program is distributed in the hope that it will be useful,
*   but WITHOUT ANY WARRANTY; without even the implied warranty of
*   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*   GNU General Public License for more details.
*
*   You should have received a copy of the GNU General Public License
*   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
    "os"
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
	"slices"
)

const VERSION = "0.1.5"

func printHelp() {
	printVersion()

	fmt.Println("\nUsage:\n\t"+os.Args[0]+" <variable_name> <variable_value> [--env-file | --input-file <file_path>] [--output-file <file_path>]\n")
	fmt.Println("Arguments:")
	fmt.Println("\tvariable_name\tREQUIRED The name of the variable")
	fmt.Println("\tvariable_value\tREQUIRED The value of the variable prtovided upon <variable_name>")
	fmt.Println("\nOptions:")
	fmt.Println("\t--env-file (or --input-file) <file_path> \tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	fmt.Println("\t--output-file <file_path> \tOPTIONAL Instead of printing the result into console write it into a file. If value provided it will NOT output the contents of the .env file.")
}

func printVersion(){
	fmt.Println("\nMkDotenv VERSION: ",VERSION)
	fmt.Println("Replace or add a variable into a .env file.")
}


func append_value_to_dotenv(file *os.File,output *bufio.Writer,variable_name string,variable_value string) (bool,error) {
	
	var newline string = fmt.Sprintf("%s=%s", variable_name, variable_value)

	// If no .env exists then output the variable.
	if (file == nil){
		output.WriteString(newline+"\n")
		return true,nil
	}

	scanner := bufio.NewScanner(file)

	var variableFound bool = false

	variable_name=strings.TrimSpace(variable_name)

	if(variable_name == ""){
		return false,errors.New("Variable name is empty")
	}

	re, err := regexp.Compile(`^#?\s*`+variable_name+`\s*=.*`)
	if err != nil {
		return false,err
	}

	
	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line

		if re.MatchString(line) {
			line_to_write = newline
			variableFound=true
		}
		output.WriteString(line_to_write+"\n")
	}

	if !variableFound {
		output.WriteString(newline+"\n")
	}

	return true,nil
}


func getParameters()(string,string,string,string){
    
	if(len(os.Args) < 3){
        fmt.Fprintln(os.Stderr,"Not enough arguments provided")
		printHelp()
		os.Exit(1)
    }

    var dotenv_filename string = ".env"
    var variable_name string = os.Args[1]
	var variable_value string = os.Args[2]
	var output_file string = ""

	if(strings.HasPrefix(variable_name,"-")){
		printHelp()
		fmt.Fprintln(os.Stderr,"Variable Name should not start with - or --")
		os.Exit(1)
	}

	ARGUMENTS:= []string{"--env-file","--input-file","--output-file","-v","--version","-h","--h","--help"}

	if(slices.Contains(ARGUMENTS[:],variable_value)){
		printHelp()
		fmt.Fprintln(os.Stderr,"\nVariable value should not contain any of the values:\n"+strings.Join(ARGUMENTS[:],"\n"))
		os.Exit(1)
	}

	for i, arg := range os.Args[3:] {

		switch arg {
		 	case "--input-file":
				fallthrough;
			case "--env-file":
				// Arguments are parsed with an offset we get the next item + offset
				dotenv_filename = os.Args[i+3+1]
				
			case "--output-file":
				output_file = os.Args[i+3+1]
		}
	}
	
	return dotenv_filename,output_file,variable_name,variable_value
}

func printVersionOrHelp(){

	if(len(os.Args) > 2 ){
		return
	}


	switch(os.Args[1]){
		case "-h":
			fallthrough
		case "--help":
			printHelp()
			os.Exit(0)
		case "-v":
			fallthrough
		case "--version":
			printVersion()
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr,"Not enough arguments provided")
			printHelp()
			os.Exit(1)
	}
}

func HandleFileError(err error, filename string) {
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: The file '%s' does not exist.\n", filename)
	} else if os.IsPermission(err) {
		fmt.Fprintf(os.Stderr, "Error: Permission denied for file '%s'.\n", filename)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Failed to open file '%s': %v\n", filename, err)
	}

	os.Exit(1)
}

func getFileToRead(dotenv_filename string) *os.File {

	var file *os.File
	var err error

	stat, _ := os.Stdin.Stat()
	hasPipeInput := (stat.Mode() & os.ModeCharDevice) == 0

	if dotenv_filename != ".env" { 
		// User explicitly provided a file, use it
		file, err = os.Open(dotenv_filename)
		if err != nil {
			HandleFileError(err, dotenv_filename)
			os.Exit(1)
		}
	} else if hasPipeInput {
		// No --env-file, but we have pipe input
		file = os.Stdin
	} else {
		// Default to .env
		file, err = os.Open(".env")
		if err != nil {

			if(os.IsNotExist(err)){
				return nil
			}

			HandleFileError(err, ".env")
			os.Exit(1)
		}
	}

	return file
}

func main() {

	if (len(os.Args) == 1 ){
		printHelp()
		os.Exit(0)
	}

	printVersionOrHelp()

	dotenv_filename,output_file,variable_name,variable_value := getParameters()

	file:=getFileToRead(dotenv_filename)
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	if output_file != "" {
		outfile,err := os.OpenFile(output_file,os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			HandleFileError(err,output_file)
		}
		defer file.Close()
		writer = bufio.NewWriter(outfile)
	}
	defer writer.Flush()

    _,err := append_value_to_dotenv(file,writer,variable_name,variable_value)
    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}