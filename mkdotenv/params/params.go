package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
)


func GetParameters(osArguments []string,errorHandle func(msg string))(string,string,string,string){

	if(len(osArguments) < 3){
		errorHandle("Not enough arguments provided")
    }

    var dotenv_filename string = ".env"
    var variable_name string = osArguments[1]
	var variable_value string = osArguments[2]
	var output_file string = ""

	if(strings.HasPrefix(variable_name,"-")){
		errorHandle("Variable Name should not start with - or --")
	}

	ARGUMENTS:= []string{"--env-file","--input-file","--output-file","-v","--version","-h","--h","--help"}

	if(slices.Contains(ARGUMENTS[:],variable_value)){
		errorHandle("\nVariable value should not contain any of the values:\n"+strings.Join(ARGUMENTS[:],"\n"))
	}

	for i, arg := range osArguments[3:] {

		value:= ""
		arg,value = sliceArgument(arg)

		switch arg {
		 	case "--input-file":
				fallthrough;
			case "--env-file":
				dotenv_filename = getValue(value,i,3,osArguments)
				
			case "--output-file":
				output_file = getValue(value,i,3,osArguments)
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

func getValue(value string,i int,offset int,arguments []string)(string){

	if(value == ""){
		// Arguments are parsed with an offset we get the next item + offset
		return arguments[i+offset+1]
	}

	return value
}

