package main

import (
    "os"
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
)

const VERSION = "0.0.1"

func printHelp() {
	printVersion()

	fmt.Println("\nUsage:\n\t"+os.Args[0]+" <variable_name> <variable_value> [--env-file <file_path>] [--no-create | -n | --n]\n")
	fmt.Println("Arguments:")
	fmt.Println("\tvariable_name\tREQUIRED The name of the variable")
	fmt.Println("\tvariable_value\tREQUIRED The value of the variable prtovided upon <variable_name>")
	fmt.Println("\nOptions:")
	fmt.Println("\t--env-file <file_path>\tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	fmt.Println("\t--no-create , -n , --n\tOPTIONAL Skip creating .env file if missing")
}

func printVersion(){
	fmt.Println("\nMkDotenv VERSION: ",VERSION)
	fmt.Println("Replace or add a variable into a .env file.")
}


func append_value_to_dotenv(file *os.File,output *bufio.Writer,variable_name string,variable_value string) (bool,error) {
	
	scanner := bufio.NewScanner(file)

	var variableFound bool = false

	variable_name=strings.TrimSpace(variable_name)

	if(variable_name == ""){
		return false,errors.New("Variable name is empty")
	}


	re, err := regexp.Compile(`^#?\s*`+variable_name+`\s*=.*`)
	if err != nil {
		return false,err
	}

	var newline string = fmt.Sprintf("%s=%s\n", variable_name, variable_value)
	
	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line

		if re.MatchString(line) {
			line_to_write = newline
			variableFound=true
		}
		output.WriteString(line_to_write+"\n")
	}

	if !variableFound {
		output.WriteString(newline+"\n")
	}

	return true,nil
}



func openFile(dotenv_filename string, create_dotenv bool)(*os.File, error){
	
	file_flags := os.O_RDWR|os.O_TRUNC

	if(create_dotenv){
		file_flags = file_flags|os.O_CREATE
	}

    // file, err := os.OpenFile(dotenv_filename,file_flags, 0644)
	file, err := os.Open(dotenv_filename)
	return file,err
}



func getParameters()(string,bool,string,string){
    
	if(len(os.Args) < 3){
        fmt.Fprintln(os.Stderr,"Not enough arguments provided")
		printHelp()
		os.Exit(1)
    }

    var create_dotenv bool =true
    var dotenv_filename string = ".env"
    var variable_name string = os.Args[1]
	var variable_value string = os.Args[2]

	if(strings.HasPrefix(variable_name,"-")){
		fmt.Fprintln(os.Stderr,"Variable Name should not start with - or --")
		printHelp()
		os.Exit(1)
	}

	for i, arg := range os.Args[3:] {

		switch arg {
			case "--env-file":
				// Arguments are parsed with an offset we get the next item + offset
				dotenv_filename = os.Args[i+3+1]
				
			case "--no-create","-n","--n":
				create_dotenv = false
		}
	}

	return dotenv_filename,create_dotenv,variable_name,variable_value
}

func printVersionOrHelp(){

	if(len(os.Args) > 2 ){
		return
	}

	switch(os.Args[1]){
		case "-h":
			fallthrough
		case "--help":
			printHelp()
			os.Exit(0)
		case "-v":
			fallthrough
		case "--version":
			printVersion()
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr,"Not enough arguments provided")
			printHelp()
			os.Exit(1)
	}
}

func main() {

	printVersionOrHelp()

	dotenv_filename,create_dotenv,variable_name,variable_value := getParameters()
    file, err := openFile(dotenv_filename,create_dotenv)

    if err != nil {
		fmt.Fprintln(os.Stderr,"Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

    _,err = append_value_to_dotenv(file,writer,variable_name,variable_value)


    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}