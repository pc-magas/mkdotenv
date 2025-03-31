package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
	"errors"
)

var FLAG_ARGUMENTS = []string{"--env-file", "--input-file", "--output-file", "-v", "--version", "-h", "--h", "--help"}

type Arguments struct {
	DotenvFilename,VariableName,VariableValue,OutputFile string
	DisplayVersion,ParseComplete  bool
}

func GetParameters(osArguments []string) (error,Arguments) {
	if len(osArguments) < 3 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		DotenvFilename: ".env",
		VariableName:   osArguments[1],
		VariableValue:  osArguments[2],
		ParseComplete:  false,
	}

	var err error=nil

	if strings.HasPrefix(args.VariableName, "-") {
		return errors.New("variable name should not start with - or --"),Arguments{}
	}

	if slices.Contains(FLAG_ARGUMENTS, args.VariableValue) {
		return errors.New("variable value should not contain reserved flag values"),Arguments{}
	}

	for i := 3; i < len(osArguments); i++ {
		arg, value := sliceArgument(osArguments[i])

		switch arg {
		case "--input-file", "--env-file":
			
			err, args.DotenvFilename = getValue(value, i,3, osArguments)
			if err != nil {
				return err,Arguments{}
			}
		case "--output-file":
			err, args.OutputFile = getValue(value, i,3, osArguments)
			if err != nil {
				return err,Arguments{}
			}
		}
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


func sliceArgument(argument string) (string, string) {
	arguments := strings.Split(argument, "=")

	if len(arguments) > 1 {
		value := strings.TrimSpace(arguments[1])
		if value == "" || value == " " {
			return arguments[0], ""
		}
		return arguments[0], value
	}

	return arguments[0], ""
}

func getValue(value string,i int,offset int,arguments []string)(error,string){

	if(value == ""){
		index:= i+offset+1
		if(index >= len(arguments) ){
			return errors.New("Index out of bounds"),value
		}
		// Arguments are parsed with an offset we get the next item + offset
		value=arguments[index]

	}

	if(slices.Contains(FLAG_ARGUMENTS[:],value)){
		return errors.New("Value contains argumwent value"),value
	}

	return nil,value
}

