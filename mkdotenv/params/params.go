package params

import(
	"os"
	"strings"
	"slices"
	"mkdotenv/msg"
	"errors"
	"flag"
	"fmt"
)

var FLAG_ARGUMENTS = []string{"--env-file", "--input-file", "--output-file", "-v", "--version", "-h", "--h", "--help"}

type Arguments struct {
	DotenvFilename,VariableName,VariableValue,OutputFile string
	ParseComplete  bool
}

func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 3 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		DotenvFilename: ".env",
		VariableName:   osArguments[1],
		VariableValue:  osArguments[2],
		OutputFile: "",
		ParseComplete:  false,
	}

	if strings.HasPrefix(args.VariableName, "-") {
		return errors.New("variable name should not start with - or --"),args
	}

	if slices.Contains(FLAG_ARGUMENTS, args.VariableValue) {
		return errors.New("variable value should not contain reserved flag values"),args
	}

	var err error=nil
	flagSet := flag.NewFlagSet("params", flag.ContinueOnError)


	var outputFile,inputFile,dotEnvFile string
	var inputFileSet bool=false
	var envFileSet bool=false

	flagSet.StringVar(&dotEnvFile,"env-file","",".env File to read upon")
	flagSet.StringVar(&inputFile,"input-file","",".env File to read upon")
	flagSet.StringVar(&outputFile,"output-file","",".env File to read upon")

	err=flagSet.Parse(osArguments[3:])
	
	if err != nil {
        return err, Arguments{}
    }

	flagSet.Visit(func(f *flag.Flag){
		inputFileSet=f.Name=="input-file"
		envFileSet=f.Name=="env-file"
	})


	fmt.Println("InputFileSet: ",inputFileSet,"envFileSet: ",envFileSet)
	fmt.Printf("outputFile: %s\ndotEnvfile:%s\n",outputFile,dotEnvFile)

	if(inputFileSet && envFileSet){
		return errors.New("Only One of `env-file` and `input-file` should be provided"),args
	}

	if(err!=nil){
		return err,args
	}


	if(inputFile!=""){
		args.DotenvFilename=inputFile
	} else if(dotEnvFile!=""){
		args.DotenvFilename=dotEnvFile
	}

	args.OutputFile=outputFile

	args.ParseComplete = true
	return nil,args
}

func PrintVersionOrHelp(){

	if(len(os.Args) > 2 ){
		return
	}

	switch(os.Args[1]){
		case "-h":
			fallthrough
		case "--help":
			msg.PrintHelp()
			os.Exit(0)
		case "-v":
			fallthrough
		case "--version":
			msg.PrintVersion()
			os.Exit(0)
		default:
			msg.ExitError("Not enough arguments provided")
	}
}
