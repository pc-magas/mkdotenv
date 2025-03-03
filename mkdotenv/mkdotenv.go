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
	"mkdotenv/params"
	"mkdotenv/msg"

)

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
		msg.PrintHelp()
		os.Exit(0)
	}

	params.PrintVersionOrHelp()

	dotenv_filename,output_file,variable_name,variable_value := params.GetParameters(os.Args)

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