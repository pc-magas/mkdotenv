package params

import(
	"slices"
	"errors"
	"fmt"
)
import flag "github.com/spf13/pflag"

type CLIArgType string

const (
	StringType CLIArgType = "string"
	BoolType   CLIArgType = "bool"
	IntType    CLIArgType = "int"
)

type FlagMeta struct {
    Name        string   // canonical flag name
	Type 		CLIArgType
	DefaultValue string
	Short 		string   // short value of the flag
    Aliases     []string // e.g., "h" is alias for "help"
    Required    bool     // whether the flag is required
    Usage       string   // help message
    Order       int      // display order
}

type Arguments struct {
	DotenvFilename string
	OutputFile string
	VariableName string
	VariableValue string
	KeepFirst bool
	DisplayHelp bool
	DisplayVersion bool
	ParseComplete  bool
	RemoveDoubles bool
	ArgumentNum int
}

var FLAG_ARGUMENTS = []string{}

var flagsMeta = []FlagMeta{
    {
        Name:     "help",
        Aliases:  []string{},
		Short: "h",
        Required: false,
        Usage:    "Display help message and exit",
		Type:  	BoolType,
        Order:    0,
    },
    {
        Name:     "version",
        Aliases:  []string{},
		Short: "v",
        Required: false,
		Type:  	BoolType,
        Usage:    "Display version and exit",
        Order:    0,
    },
    {
        Name:     "variable-name",
        Aliases:  []string{},
		Type: StringType,
        Required: true,
        Usage:    "Name of the variable to be set",
        Order:    1,
    },
    {
        Name:     "variable-value",
        Aliases:  []string{},
		Type: StringType,
        Required: true,
        Usage:    "Value to assign to the variable",
        Order:    1,
    },
    {
        Name:     "env-file",
        Aliases:  []string{"input-file"},
		Type: StringType,
        Required: false,
		DefaultValue: ".env",
        Usage:    "Input .env file path (default .env)",
        Order:    2,
    },
    {
        Name:     "output-file",
        Aliases:  []string{},
		Type: StringType,
        Required: false,
		DefaultValue: ".env",
        Usage:    "File to write output to (`-` for stdout)",
        Order:    2,
    },
    {
        Name:     "remove-doubles",
		Type: BoolType,
        Aliases:  []string{},
        Required: false,
        Usage:    "Remove duplicate variable entries, keeping the first",
        Order:    3,
    },
}


func initFlags() (*flag.FlagSet) {

	flagSet := flag.NewFlagSet("params", flag.ContinueOnError)

	for _, meta := range flagsMeta {

		FLAG_ARGUMENTS=append(FLAG_ARGUMENTS,"--"+meta.Name)
		FLAG_ARGUMENTS=append(FLAG_ARGUMENTS,"-"+meta.Name)

        switch meta.Type {
			case StringType:
				
				if(meta.Short == ""){
					flagSet.String(meta.Name, meta.DefaultValue, meta.Usage)
				} else {
					flagSet.StringP(meta.Name, meta.Short, meta.DefaultValue, meta.Usage)
				}

				for _, alias := range meta.Aliases {
					FLAG_ARGUMENTS=append(FLAG_ARGUMENTS,"--"+alias)
					flagSet.String(alias, meta.DefaultValue, "(alias of --"+meta.Name+") "+meta.Usage)
				}

			case BoolType:
				def := meta.DefaultValue == "true"

				if(meta.Short == ""){
					flagSet.Bool(meta.Name, def, meta.Usage)
				} else {
					flagSet.BoolP(meta.Name,meta.Short, def, meta.Usage)
				}
				
				for _, alias := range meta.Aliases {
					flagSet.Bool(alias, def, "(alias of --"+meta.Name+") "+meta.Usage)
				}
        }
    }

	return flagSet
}

func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 1 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		VariableName:   "",
		VariableValue:  "",
		OutputFile: ".env",
		DotenvFilename: ".env",
		RemoveDoubles: false,
		DisplayHelp: false,
		DisplayVersion: false,
		ArgumentNum: len(osArguments),
		ParseComplete:  false,
	}


	var err error=nil
	var inputFileSet,outputFileSet bool=false,false
	
	var flagSet *flag.FlagSet = initFlags()
	err=flagSet.Parse(osArguments[1:])
	if err != nil {
        return err, args
    }

	flagSet.Visit(func(f *flag.Flag){

		if (slices.Contains(FLAG_ARGUMENTS,f.Value.String())){
			err=fmt.Errorf("Flag %s should not contain a param value",f.Name)
			return
		}

		if(err !=nil){
			return
		}
		
		value:=f.Value.String()

		if(value == ""){
			err=fmt.Errorf("Value should not be empty for param %s",f.Name)
			return
		}
		
		switch (f.Name){

			case "input-file","env-file":
				
				if(inputFileSet){
					err=fmt.Errorf("Only One of `--env-file` and `--input-file` should be provided")
					return
				}
				
				if(value == ""){
					err=fmt.Errorf("Only One of `--env-file` and `--input-file` should be provided")
					return
				}

				args.DotenvFilename = value
				inputFileSet=true

			case "output-file":

				if(outputFileSet){
					err=fmt.Errorf("Output File has Already Been provided")
					return
				}

				if(value == ""){
					err=fmt.Errorf(f.Name+" should not be empty")
					return
				}

				args.OutputFile=value
				outputFileSet=true

			case "variable-name":
				args.VariableName=value

			case "variable-value":
				args.VariableValue=value
			case "remove-doubles":
				args.RemoveDoubles = value=="true"
			case "h","help":
				args.DisplayHelp = value=="true"
			case "v","version":
				args.DisplayVersion = value=="true"
		}

	})

	if(err!=nil){
		return err, args
	}

	args.ParseComplete = true
	return nil,args
}

func GetFlagsMeta() []FlagMeta {
    out := make([]FlagMeta, len(flagsMeta))
    copy(out, flagsMeta)
    return out
}
