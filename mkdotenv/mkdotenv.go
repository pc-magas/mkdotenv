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
	"mkdotenv/params"
	"mkdotenv/msg"
	"mkdotenv/files"
)

func main() {

	if (len(os.Args) == 1 ){
		msg.PrintHelp()
		os.Exit(0)
	}

	params.PrintVersionOrHelp()

	dotenv_filename,output_file,variable_name,variable_value := params.GetParameters(os.Args)

	file:=files.GetFileToRead(dotenv_filename)
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	if output_file != "" {
		outfile,err := os.OpenFile(output_file,os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			files.HandleFileError(err,output_file)
		}
		defer file.Close()
		writer = bufio.NewWriter(outfile)
	}
	defer writer.Flush()

    _,err := files.AppendValueToDotenv(file,writer,variable_name,variable_value)
    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}