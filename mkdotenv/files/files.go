package files

import (
    "os"
    "fmt"
    "bufio"
	"strings"
	"regexp"
	"errors"
	"io"
)


func AppendValueToDotenv(input *io.Reader,output *bufio.Writer,variable_name string,variable_value string) (bool,error) {
	
	var newline string = fmt.Sprintf("%s=%s", variable_name, variable_value)

	// If no .env exists then output the variable.
	if (input == nil){
		output.WriteString(newline+"\n")
		return true,nil
	}

	scanner := bufio.NewScanner(*input)

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

func HandleFileError(err error, filename string) {

	if (err == nil){
		return;
	}

	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: The file '%s' does not exist.\n", filename)
	} else if os.IsPermission(err) {
		fmt.Fprintf(os.Stderr, "Error: Permission denied for file '%s'.\n", filename)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Failed to open file '%s': %v\n", filename, err)
	}

	os.Exit(1)
}


func CopyFile(dotenv_filename string,dest_filename string) {

	file, err:= os.Open(dotenv_filename)
	HandleFileError(err, dotenv_filename)

	newFile,err:= os.Create(dest_filename)
	HandleFileError(err, dest_filename)

	_,err=io.Copy(newFile,file)
	HandleFileError(err,dest_filename)
	
	err = newFile.Sync()
	HandleFileError(err,dest_filename)

	file.Close()
	newFile.Close()
}

func GetFileToRead(dotenv_filename string) *os.File {

	var file *os.File
	var err error

	stat, _ := os.Stdin.Stat()
	hasPipeInput := (stat.Mode() & os.ModeCharDevice) == 0

	// Input is piped through STDIN
	if(hasPipeInput){
		return os.Stdin
	}

	if(dotenv_filename == ""){
		dotenv_filename = ".env"
	}
	
	HandleFileError(err, dotenv_filename)
	
	file,err = os.Open(dotenv_filename)
	HandleFileError(err,dotenv_filename)

	return file
}


func CreateWriter(filename string) (*bufio.Writer,*os.File) {
	
	if(filename == "-"){
		return bufio.NewWriter(os.Stdout),nil
	}

	if(filename == ""){
		filename=".env"
	}

	outfile,err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		HandleFileError(err,filename)
	}
		
	return bufio.NewWriter(outfile),outfile
}
