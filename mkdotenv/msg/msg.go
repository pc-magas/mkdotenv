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

	fmt.Println("\nUsage:\n\t"+os.Args[0]+" [-v|--version|-h|--help] <variable_name> <variable_value> [--env-file | --input-file <file_path>] [--output-file <file_path>]\n")
	fmt.Println("Arguments:")
	fmt.Println("\tvariable_name\tREQUIRED The name of the variable")
	fmt.Println("\tvariable_value\tREQUIRED The value of the variable provided upon <variable_name>")
	fmt.Println("\nOptions:")
	fmt.Println("\t-v (or --version)\tOPTIONAL Display Version Number. If provided any other argument is ignored.")
	fmt.Println("\t-h (or --help)\tOPTIONAL Display the current message. If provided any other argument is ignored.")
	fmt.Println("\t--env-file (or --input-file) <file_path>\tOPTIONAL The .env file path in <file_path> that will be manipulated. Default value .env")
	fmt.Println("\t--output-file <file_path>\tOPTIONAL Instead of printing the result into console write it into a file.")
}

func PrintVersion(){
	fmt.Println("\nMkDotenv VERSION: ",version)
	fmt.Println("Replace or add a variable into a .env file.")
}