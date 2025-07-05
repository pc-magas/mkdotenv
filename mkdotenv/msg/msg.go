package msg

import(
	"os"
	"fmt"
)

// This is changed upon runtime.
var version = "dev"

func ExitError(msg string){
	fmt.Fprintln(os.Stderr,msg)
	PrintHelp()
	os.Exit(1)
}

func PrintHelp() {
	PrintVersion()

	fmt.Fprintln(os.Stderr,"\nUsage:\n\t"+os.Args[0]+" [-v|--version|-h|--help] --variable-name <variable_name> --variable-value <variable_value> [--env-file | --input-file <file_path>] [--output-file <file_path>]\n")
	fmt.Fprintln(os.Stderr,"\nOptions:")
	fmt.Fprintln(os.Stderr,"\t--variable_name <variable_name>\tREQUIRED The name of the variable")
	fmt.Fprintln(os.Stderr,"\t--variable_value <variable_value>\tREQUIRED The value of the variable provided upon <variable_name>")
	fmt.Fprintln(os.Stderr,"\t-v (or --version)\tOPTIONAL Display Version Number. If provided any other argument is ignored.")
	fmt.Fprintln(os.Stderr,"\t-h (or --help)\tOPTIONAL Display the current message. If provided any other argument is ignored.")
	fmt.Fprintln(os.Stderr,"\t--env-file (or --input-file) <file_path>\tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	fmt.Fprintln(os.Stderr,"\t--output-file <file_path>\tOPTIONAL Instead of printing the result into console write it into a file.")
}

func PrintVersion(){
	fmt.Fprintln(os.Stderr,"\nMkDotenv VERSION: ",version)
	fmt.Fprintln(os.Stderr,"Replace or add a variable into a .env file.")
}