package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
	"errors"
	"flag"
	"fmt"
)

var FLAG_ARGUMENTS = []string{"--env-file", "--input-file", "--output-file", "-v", "--version", "-h", "--h", "--help"}

type Arguments struct {
	DotenvFilename,VariableName,VariableValue,OutputFile string
	ParseComplete  bool
}

func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 3 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		DotenvFilename: ".env",
		VariableName:   osArguments[1],
		VariableValue:  osArguments[2],
		OutputFile: "",
		ParseComplete:  false,
	}

	if strings.HasPrefix(args.VariableName, "-") {
		return errors.New("variable name should not start with - or --"),args
	}

	if slices.Contains(FLAG_ARGUMENTS, args.VariableValue) {
		return errors.New("variable value should not contain reserved flag values"),args
	}

	var err error=nil
	var inputFileSet,outputFileSet bool=false,false
	
	flagSet := flag.NewFlagSet("params", flag.ContinueOnError)

	flagSet.String("env-file","",".env File to read upon")
	flagSet.String("input-file","",".env File to read upon")
	flagSet.String("output-file","",".env File to read upon")


	err=flagSet.Parse(osArguments[3:])
		

	if err != nil {
        return err, args
    }

	flagSet.Visit(func(f *flag.Flag){

		if(slices.Contains(FLAG_ARGUMENTS,f.Value.String())){
			err=fmt.Errorf("Flag %s should not contain a param value",f.Name)
			return
		}

		if(err !=nil){
			return
		}

		value:=f.Value.String()

		if(value == ""){
			err=fmt.Errorf("Value should not be empty for param %s",f.Name)
			return
		}

		switch (f.Name){

			case "input-file","env-file":

				if(inputFileSet){
					err=fmt.Errorf("Only One of `--env-file` and `--input-file` should be provided")
					return
				}
				
				if(value == ""){
					err=fmt.Errorf("Only One of `--env-file` and `--input-file` should be provided")
					return
				}

				args.DotenvFilename = value
				inputFileSet=true

			case "output-file":

				if(outputFileSet){
					err=fmt.Errorf("Output File has Already Been provided")
					return
				}

				args.OutputFile=value
				outputFileSet=true
		}

	})

	if(err!=nil){
		return err, args
	}

	args.ParseComplete = true
	return nil,args
}

func PrintVersionOrHelp(){

	if(len(os.Args) > 2 ){
		return
	}

	switch(os.Args[1]){
		case "-h":
			fallthrough
		case "--help":
			msg.PrintHelp()
			os.Exit(0)
		case "-v":
			fallthrough
		case "--version":
			msg.PrintVersion()
			os.Exit(0)
		default:
			msg.ExitError("Not enough arguments provided")
	}
}
