package files

import (
    "os"
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
)


func AppendValueToDotenv(file *os.File,output *bufio.Writer,variable_name string,variable_value string) (bool,error) {
	
	var newline string = fmt.Sprintf("%s=%s", variable_name, variable_value)

	// If no .env exists then output the variable.
	if (file == nil){
		output.WriteString(newline+"\n")
		return true,nil
	}

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

	
	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line
		
		length_delta:=len(line)- len(newline)
		if(length_delta < 0){
			length_delta = -length_delta
		}

		if re.MatchString(line) {
			line_to_write = newline	
			variableFound=true
		}

		output.WriteString(line_to_write+"\n"+strings.Repeat(" ",length_delta))
	}

	if !variableFound {
		output.WriteString(newline+"\n")
	}

	return true,nil
}

func HandleFileError(err error, filename string) {
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: The file '%s' does not exist.\n", filename)
	} else if os.IsPermission(err) {
		fmt.Fprintf(os.Stderr, "Error: Permission denied for file '%s'.\n", filename)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Failed to open file '%s': %v\n", filename, err)
	}

	os.Exit(1)
}

func GetFileToRead(dotenv_filename string) *os.File {

	var file *os.File
	var err error

	stat, _ := os.Stdin.Stat()
	hasPipeInput := (stat.Mode() & os.ModeCharDevice) == 0

	if dotenv_filename != ".env" { 
		// User explicitly provided a file, use it
		file, err = os.Open(dotenv_filename)
		if err != nil {
			HandleFileError(err, dotenv_filename)
			os.Exit(1)
		}
	} else if hasPipeInput {
		// No --env-file, but we have pipe input
		file = os.Stdin
	} else {
		// Default to .env
		file, err = os.Open(".env")
		if err != nil {

			if(os.IsNotExist(err)){
				return nil
			}

			HandleFileError(err, ".env")
			os.Exit(1)
		}
	}

	return file
}