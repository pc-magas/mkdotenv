package main

import (
    "os"
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
)


func append_value_to_dotenv(file *os.File,variable_name string,variable_value string) (bool,error) {
	
	scanner := bufio.NewScanner(file)

	var updatedContent strings.Builder
	var variableFound bool
	var line_to_write string

	variable_name=strings.TrimSpace(variable_name)

	if(variable_name == ""){
		return false,errors.New("Variable name is empty")
	}


	re, err := regexp.Compile(`^#?\s*`+variable_name+`\s*=\s*`)
	if err != nil {
		return false,err
	}

	var newline string = fmt.Sprintf("\n# Updated from mkdotenv\n%s=%s\n", variable_name, variable_value)

	for scanner.Scan() {
		var line = scanner.Text()
		line_to_write = line
		if re.MatchString(variable_name) {
			line_to_write = newline
			variableFound = true
		}

		updatedContent.WriteString(line_to_write + "\n")
	}

	if !variableFound {
		updatedContent.WriteString(newline)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return false,err
	}

	_, err = file.WriteString(updatedContent.String())
	if err != nil {
		return false,err
	}

	return true,nil
}

func printHelp(){

}

func printVersion(){

}


func openFile(dotenv_filename string, create_dotenv bool){
	
	file_flags = os.O_RDWR|os.O_TRUNC

	if(create_dotenv){
		file_flags = file_flags|os.O_CREATE
	}


    file, err := os.OpenFile(dotenv_filename,file_flags, 0644)

	defer file.Close()

	return file,err
}


func getParameters()(bool,string,string,string){
    
	if(len(os.Args) < 3){
        fmt.Fprintln(os.Stderr,"Not enough arguments provided")
        os.Exit(1)
    }

    var create_dotenv bool :=true
    var dotenv_filename string :=".env"
    var variable_name string := os.Args[1]
	var variable_value string := os.Args[2]

	if(len(os.Args) > 3){
		for i, arg := range os.Args[3:] {

			switch arg {
				case "--env-file":
					dotenv_filename = os.Args[i+1]
					break;
				
				case "--no-create":
				case "-n":
					create_dotenv = false
					break;
			}
		}
	}

	return (dotenv_filename,variable_name,variable_value,create_dotenv)
}

func main() {

	if(len(os.Args) == 2 ){
		switch(os.Args[1]){
			case "--h": 
			case "--help":
				printHelp()
				os.Exit(0)
			case "-v":
			case "--version":
				printVersion()
				os.Exit(0)
			default:
				fmt.Fprintln(os.Stderr,"Not enough arguments provided")
				os.Exit(1)
		}
	}

	dotenv_filename,variable_name,variable_value,create_dotenv := getParameters()

    file, err := openFile(dotenv_filename,create_dotenv)

    if err != nil {
		fmt.Fprintln(os.Stderr,"Error opening file:", err)
		os.Exit(1)
	}

    _,err = append_value_to_dotenv(file,variable_name,variable_value)    

    if(err!=nil){
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
}