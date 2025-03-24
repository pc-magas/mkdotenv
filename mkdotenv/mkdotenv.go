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
	"time"
	"strconv"
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

	dotenv_filename,output_file,variable_name,variable_value := params.GetParameters(os.Args,msg.ExitError)

	filenameToRead := dotenv_filename
	filenameCopy:=dotenv_filename+"."+strconv.FormatInt(time.Now().UnixMilli(),10)
	sameFileToReadAndWrite:=dotenv_filename == output_file 

	// If inputfile is same as Outputfile copy the input file to a temporary one
	if(sameFileToReadAndWrite){
		files.CopyFile(dotenv_filename,filenameCopy)
		filenameToRead=filenameCopy
	}


	file:=files.GetFileToRead(filenameToRead)
	defer file.Close()

	writer,outfile := files.CreateWriter(output_file)
	if(outfile!=nil){
		defer outfile.Close()
	}
	defer writer.Flush()

	_,err := files.AppendValueToDotenv(file,writer,variable_name,variable_value)

    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
		if(sameFileToReadAndWrite){
			files.CopyFile(filenameCopy,dotenv_filename)
		}
        os.Exit(1)
    }

	if sameFileToReadAndWrite {
		file.Close()
		err := os.Remove(filenameCopy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to remove temp file %s: %v\n", filenameCopy, err)
		}
	}
}