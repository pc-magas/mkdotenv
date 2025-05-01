package core

import (
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
	"io"
)

func AppendValueToDotenv(input io.Reader,output *bufio.Writer,variable_name string,variable_value string) (bool,error) {
	
	var newline string = fmt.Sprintf("%s=%s", variable_name, variable_value)

	// If no .env exists then output the variable.
	if (input == nil){
		output.WriteString(newline+"\n")
		return true,nil
	}

	scanner := bufio.NewScanner(input)

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