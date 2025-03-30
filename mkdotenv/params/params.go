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
	DotenvFilename  string
	VariableName    string
	VariableValue   string
	OutputFile      string
	DisplayVersion  bool
	ParseComplete   bool
}

func GetParameters(osArguments []string) (Arguments, error) {
	if len(osArguments) < 3 {
		return Arguments{}, errors.New("not enough arguments provided")
	}

	args := Arguments{
		DotenvFilename: ".env",
		VariableName:   osArguments[1],
		VariableValue:  osArguments[2],
		ParseComplete:  false,
	}

	if strings.HasPrefix(args.VariableName, "-") {
		return Arguments{}, errors.New("variable name should not start with - or --")
	}

	if slices.Contains(FLAG_ARGUMENTS, args.VariableValue) {
		return Arguments{}, errors.New("variable value should not contain reserved flag values")
	}

	for i := 3; i < len(osArguments); i++ {
		arg, value := parseArgument(osArguments[i])

		switch arg {
		case "--input-file", "--env-file":
			var err error
			args.DotenvFilename, err = getArgumentValue(value, i, osArguments)
			if err != nil {
				return Arguments{}, err
			}
		case "--output-file":
			var err error
			args.OutputFile, err = getArgumentValue(value, i, osArguments)
			if err != nil {
				return Arguments{}, err
			}
		}
	}

	args.ParseComplete = true
	return args, nil
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

