package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
	"errors"
)

var FLAG_ARGUMENTS=[8]string{"--env-file","--input-file","--output-file","-v","--version","-h","--h","--help"}


func GetParameters(osArguments []string,errorHandle func(msg string))(string,string,string,string){

	if(len(osArguments) < 3){
		errorHandle("Not enough arguments provided")
    }

    var dotenv_filename string = ".env"
    var variable_name string = osArguments[1]
	var variable_value string = osArguments[2]
	var output_file string = ""
	var err error=nil

	if(strings.HasPrefix(variable_name,"-")){
		errorHandle("Variable Name should not start with - or --")
		return dotenv_filename,output_file,variable_name,variable_value
	}

	if(slices.Contains(FLAG_ARGUMENTS[:],variable_value)){
		errorHandle("\nVariable value should not contain any of the values:\n"+strings.Join(FLAG_ARGUMENTS[:],"\n"))
		return dotenv_filename,output_file,variable_name,variable_value
	}

	for i, arg := range osArguments[3:] {

		value:= ""
		arg,value = sliceArgument(arg)

		switch arg {
		 	case "--input-file":
				fallthrough;
			case "--env-file":
				err,dotenv_filename = getValue(value,i,3,osArguments)
				
			case "--output-file":
				err,output_file = getValue(value,i,3,osArguments)
		}

		if(err!=nil){
			errorHandle("Unable to find the argumewnt value")
			return dotenv_filename,output_file,variable_name,variable_value
		}
	}
	
	return dotenv_filename,output_file,variable_name,variable_value
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
	if(slices.Contains(FLAG_ARGUMENTS[:],value)){
		return errors.New("Value contains argumwent value"),value
	}

	if(value == ""){
		index:= i+offset+1
		if(index >= len(arguments) ){
			return errors.New("Index out of bounds"),value
		}
		// Arguments are parsed with an offset we get the next item + offset
		return nil,arguments[index]

	}


	return nil,value
}

