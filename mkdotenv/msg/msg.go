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

	fmt.Fprintln(os.Stderr,"\nUsage:\n\t"+os.Args[0]+" [-v|--version|-h|--help] --variable-name <variable_name> --variable-value <variable_value> [--env-file | --input-file <file_path>] [--output-file <file_path>] [--update-first] [--keep-first]\n")
	fmt.Fprintln(os.Stderr,"\nOptions:")
	fmt.Fprintln(os.Stderr,"\t--variable_name <variable_name>\tREQUIRED The name of the variable")
	fmt.Fprintln(os.Stderr,"\t--variable_value <variable_value>\tREQUIRED The value of the variable provided upon <variable_name>")
	fmt.Fprintln(os.Stderr,"\t-v, --version \tOPTIONAL Display Version Number.")
	fmt.Fprintln(os.Stderr,"\t-h, --help \tOPTIONAL Display the current message.")
	fmt.Fprintln(os.Stderr,"\t--env-file, --input-file <file_path>\tOPTIONAL Path to the .env file to modify. Default is `.env`.")
	fmt.Fprintln(os.Stderr,"\t--output-file <file_path>\tOPTIONAL Write the result to a file. Value `-` prints to console default is `.env`")
	fmt.Fprintln(os.Stderr,"\t--keep-first \tOPTIONAL Keep only the first occuirence and remove the rest occurences of the variable having <variable_name>")
}

func PrintVersion(){
	fmt.Fprintln(os.Stderr,"\nMkDotenv VERSION: ",version)
	fmt.Fprintln(os.Stderr,"Replace or add a variable into a .env file.")
}