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
	// "time"
	// "strconv"
	"github.com/pc-magas/mkdotenv/params"
	"github.com/pc-magas/mkdotenv/msg"
	"github.com/pc-magas/mkdotenv/files"
	"github.com/pc-magas/mkdotenv/core"
	"github.com/pc-magas/mkdotenv/core/executor"
)

func displayVersionOrHelp(paramStruct params.Arguments){

	if(paramStruct.DisplayHelp){
		msg.PrintHelp()
		os.Exit(0)
	}

	if(paramStruct.DisplayVersion){
		msg.PrintVersion()
		os.Exit(0)
	}
}

func main() {

	paramErr,paramStruct := params.GetParameters(os.Args)

	if (paramStruct.ArgumentNum == 1 ){
		msg.PrintHelp()
		os.Exit(0)
	}

	if(paramErr != nil){
		msg.ExitError(paramErr.Error())
	}

	displayVersionOrHelp(paramStruct)

	filenameToRead := paramStruct.TemplateFile

	if(paramStruct.TemplateFile == paramStruct.OutputFile){
		msg.ExitError("Template file and output file should not be the same.")
	}

	file:=files.GetFileToRead(filenameToRead)
	defer file.Close()

	writer,outfile := files.CreateWriter(paramStruct.OutputFile)
	if(outfile!=nil){
		defer outfile.Close()
	}
	defer writer.Flush()

	manipulator:= core.NewDotEnvManipulator(file,executor.NewExecutor())

	err := manipulator.Replace(writer,paramStruct.Environment)

    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}