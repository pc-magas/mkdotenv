package core

import (
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
	"io"
)

func validateVarName(variable_name string) (error){

	if(variable_name == ""){
		return errors.New("Variable Name should not be an empty string")
	}

	re, _ := regexp.Compile(`^[A-Za-z\d+_]+$`)
	if(!re.MatchString(variable_name)){
		return errors.New("Variable name is Invalid string")
	}

	return nil
}

func AppendValueToDotenv(input io.Reader,output *bufio.Writer,variable_name string,variable_value string,removeDoubles bool) (bool,error) {
	
	variable_name=strings.TrimSpace(variable_name)
	validationError:=validateVarName(variable_name)
	if validationError!=nil{
		return false, validationError
	}

	var newline string = fmt.Sprintf("%s=%s", variable_name, variable_value)

	// If no .env exists then output the variable.
	if (input == nil){
		output.WriteString(newline+"\n")
		return true,nil
	}

	scanner := bufio.NewScanner(input)

	var variableFound bool = false

	re, err := regexp.Compile(`^#?\s*`+variable_name+`\s*=.*`)
	if err != nil {
		return false,err
	}
	
	for scanner.Scan() {
		line:=scanner.Text()
		line_to_write:=line
		
		if re.MatchString(line) {

			if(variableFound && removeDoubles){
				continue
			}

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