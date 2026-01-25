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
	"bufio"
	"github.com/pc-magas/mkdotenv/params"
	"github.com/pc-magas/mkdotenv/msg"
	"github.com/pc-magas/mkdotenv/core"
	"github.com/pc-magas/mkdotenv/core/executor"
	"github.com/pc-magas/mkdotenv/core/context"
)

func readTemplateFile(dotenv_filename string) *os.File {

	var file *os.File
	var err error

	stat, _ := os.Stdin.Stat()
	hasPipeInput := (stat.Mode() & os.ModeCharDevice) == 0

	// Input is piped through STDIN
	if(hasPipeInput){
		return os.Stdin
	}

	// TODO Raise Error
	// if(dotenv_filename == ""){
	// 	dotenv_filename = ".env"
	// }
	
	msg.HandleFileError(err, dotenv_filename)
	
	file,err = os.Open(dotenv_filename)
	msg.HandleFileError(err,dotenv_filename)

	return file
}


func createWriter(filename string) (*bufio.Writer,*os.File) {
	
	if(filename == "-"){
		return bufio.NewWriter(os.Stdout),nil
	}

	// TODO Raise error initialization happens outside
	// if(filename == ""){
	// 	filename=".env"
	// }

	outfile,err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		msg.HandleFileError(err,filename)
	}
		
	return bufio.NewWriter(outfile),outfile
}


func main() {

	paramErr,paramStruct := params.GetParameters(os.Args)

	if(paramErr != nil){
		msg.ExitError(paramErr.Error(),true)
	}

	if(paramStruct.DisplayHelp){
		msg.PrintHelp()
		os.Exit(0)
	}

	if(paramStruct.DisplayVersion){
		msg.PrintVersion()
		os.Exit(0)
	}

	if(paramStruct.TemplateFile == paramStruct.OutputFile){
		msg.ExitError("Template file and output file should not be the same.",false)
	}

	templateFile:=readTemplateFile(paramStruct.TemplateFile)
	defer templateFile.Close()

	writer,outfile := createWriter(paramStruct.OutputFile)
	if(outfile!=nil){
		defer outfile.Close()
	}
	defer writer.Flush()

	manipulator:= core.NewDotEnvManipulator(templateFile,executor.NewExecutor())

	resolutionContext,err:=context.NewResolutionContext(paramStruct.TemplateFile,paramStruct.MiscArguments)

	if(err!=nil){
		msg.ExitError(err.Error(),false)
	}

	err = manipulator.Replace(writer,paramStruct.Environment,resolutionContext)

    if(err!=nil){
       msg.ExitError(err.Error(),false)
    }
}