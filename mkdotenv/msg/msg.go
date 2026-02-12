package msg

import(
	"os"
	"fmt"
	"github.com/pc-magas/mkdotenv/params/usage"
)

// value changed upon compile do not remove this.
var version = "dev"

func ExitError(msg string, displayUsage bool){
	fmt.Fprintln(os.Stderr,msg)
	os.Exit(1)
}

func PrintHelp() {
	PrintVersion()
	executable:=os.Args[0]
	fmt.Fprintln(os.Stderr,"\nUsage:\n\t"+executable+" \\"+usage.BuildCommandUsage()+"\n")
	fmt.Fprint(os.Stderr,"Options:\n\n")
	fmt.Fprintln(os.Stderr,usage.BuildArgumentUsage())
}

func PrintVersion(){
	fmt.Fprintln(os.Stderr,"\nMkDotenv VERSION: ",version)
	fmt.Fprintln(os.Stderr,"Replace or add a variable into a .env file.")
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

