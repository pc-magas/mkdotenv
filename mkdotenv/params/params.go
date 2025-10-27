package params

import(
	"strings"
	"errors"
	"fmt"
)

import flag "github.com/spf13/pflag"


func valiDateCommon(value string) bool {
	if(value == ""){
		return false
	}

	return true
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

		meta := SearchFlag(f.Name)
		value:=f.Value.String()

		if meta != nil && meta.Validator != nil {
			fmt.Println("Validate", meta.Name)
			if ! meta.Validator(value) {
				// stop early with validation error
				err = fmt.Errorf("invalid value for --%s", f.Name)
				return
			}
		}
		
		// Argument Value also should not be a flag name as well
		if(SearchFlag(value) != nil){
			fmt.Println("Hello")
			err = fmt.Errorf("Value should not be an argument value")
			return
		}

		switch (f.Name){

			case "template-file","template","t":
				
				if(inputFileSet){
					err=fmt.Errorf("Only one of `--tempalte-file`, `--template`,`-t` values should be provided")
					return
				}
				
				if(value == ""){
					err=fmt.Errorf("`--tempalte-file`, `--template`,`-t` should contain the template file that would be processed")
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


func SearchFlag(name string) *FlagMeta {
	name = strings.Trim(name,"-")
    for i := range flagsMeta {
        if flagsMeta[i].Name == name {
            return &flagsMeta[i]
        }
        for _, alias := range flagsMeta[i].Aliases {
            if alias == name {
                return &flagsMeta[i]
            }
        }
    }
    return nil
}