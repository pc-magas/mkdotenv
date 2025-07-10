package params

import(
	"os"
	"slices"
	"github.com/pc-magas/mkdotenv/msg"
	"errors"
	"flag"
	"fmt"
)

var FLAG_ARGUMENTS = []string{"--env-file", "--input-file", "--output-file", "-v", "--version", "-h", "--h", "--help","--variable-name"}

type Arguments struct {
	DotenvFilename,OutputFile,VariableName,VariableValue string
	KeepFirst bool
	DisplayHelp bool
	DisplayVersion bool
	ParseComplete  bool
}

func initFlags() (*flag.FlagSet) {

	flagSet := flag.NewFlagSet("params", flag.ContinueOnError)

	flagSet.String("variable-name", "", "REQUIRED The name of the variable")
	flagSet.String("variable-value", "", "REQUIRED The value of the variable provided upon <variable_name>")
	flagSet.String("env-file", "", "<file_path>\tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	flagSet.String("input-file", "", "<file_path>\tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	flagSet.String("output-file", "", "<file_path>\tOPTIONAL Instead of printing the result into console write it into a file.")
	flagSet.Bool("keep-first",false,"\tOPTIONAL Keep only the first occuirence and remove the rest occurences of the variable having <variable_name>")
	flagSet.Bool("h",false,"\tOPTIONAL Display the current message.")
	flagSet.Bool("help",false,"\tOPTIONAL Display the current message.")
	flagSet.Bool("v",false,"\tOPTIONAL Display Version Number.")
	flagSet.Bool("version",false,"tOPTIONAL Display Version Number.")

	return flagSet
}

func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 3 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		VariableName:   "",
		VariableValue:  "",
		OutputFile: ".env",
		DotenvFilename: ".env",
		KeepFirst: false,
		DisplayHelp: false,
		DisplayVersion: false,
		ParseComplete:  false,
	}


	var err error=nil
	var inputFileSet,outputFileSet bool=false,false
	
	var flagSet *flag.FlagSet = initFlags()
	err=flagSet.Parse(osArguments[1:])
	if err != nil {
        return err, args
    }

	flagSet.Visit(func(f *flag.Flag){

		if (slices.Contains(FLAG_ARGUMENTS,f.Value.String())){
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

				if(value == ""){
					err=fmt.Errorf(f.Name+" should not be empty")
					return
				}

				args.OutputFile=value
				outputFileSet=true

			case "variable-name":
				args.VariableName=value

			case "variable-value":
				args.VariableValue=value
			case "keep-first":
				args.KeepFirst = value=="true"
			case "h","help":
				args.DisplayHelp = value=="true"
			case "v","version":
				args.DisplayVersion = value=="true"
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
			flag.Usage()
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
