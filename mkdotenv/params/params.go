package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
	"errors"
)

var FLAG_ARGUMENTS=[8]string{"--env-file","--input-file","--output-file","-v","--version","-h","--h","--help"}

type Arguments struct {
	dotenv_filename,variable_name,variable_value,output_file string
	dislay_version,parse_complete bool
}

func GetParameters(osArguments []string)(error,Arguments){

	var err error=nil

	arguments:= Arguments{
		dotenv_filename:".env",
		variable_name:osArguments[1],
		variable_value:osArguments[2],
		dislay_version:false,
		parse_complete:false}


	if(len(osArguments) < 3){
		err = errors.New("Not enough arguments provided") 
		return err,arguments
    }


	if(strings.HasPrefix(arguments.variable_name,"-")){
		err = errors.New("Variable Name should not start with - or --")
		return err,arguments
	}

	if(slices.Contains(FLAG_ARGUMENTS[:],arguments.variable_value)){
		err = errors.New("\nVariable value should not contain any of the values:\n"+strings.Join(FLAG_ARGUMENTS[:],"\n"))
		return err,arguments
	}

	for i, arg := range osArguments[3:] {

		value:= ""
		arg,value = sliceArgument(arg)

		switch arg {
		 	case "--input-file":
				fallthrough;
			case "--env-file":
				err,arguments.dotenv_filename = getValue(value,i,3,osArguments)
				
			case "--output-file":
				err,arguments.output_file = getValue(value,i,3,osArguments)
		}

		if(err!=nil){
			// errorHandle("Unable to find the argumewnt value")
			return err,arguments
		}
	}

	arguments.parse_complete=true
	return err,arguments
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

