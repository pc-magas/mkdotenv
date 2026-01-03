package params

import(
	"errors"
	"github.com/pc-magas/mkdotenv/params/parser"
)

type Arguments struct {
	Environment string
	TemplateFile string 
	OutputFile string
	MiscArguments map[string]string
    DisplayHelp bool
	DisplayVersion bool
	ArgumentNum int
    ParseComplete bool
}

func getDefault(name string) any {
	for _, f := range flagsMeta {
		if f.Name == name {
			return f.DefaultValue
		}
	}
	return nil
}

func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 1 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		Environment:    getDefault("environment").(string),
		TemplateFile:   getDefault("template-file").(string),
		OutputFile:     getDefault("output-file").(string),
		ArgumentNum:    len(osArguments),
		MiscArguments:  make(map[string]string),
		DisplayHelp:    false,
		DisplayVersion: false,
		ParseComplete:  false,
	}

    paramParser := parser.NewParamParser[Arguments](flagsMeta)

	paramParser.OnAssign = func(meta parser.FlagMeta, value string, args *Arguments) error {
        switch meta.Name {
		case "help":
			args.DisplayHelp=true
		case "version":
			args.DisplayVersion=true
		case "environment":
			args.Environment=value
		case "template-file":
			args.TemplateFile = value
		case "output-file":
			args.OutputFile = value
		default:
			args.MiscArguments[meta.Name] = value
        }
        return nil
    }

	_,error:=paramParser.Parse(osArguments,&args)
	
	if(error!=nil){
		return error,args
	}

	return nil,args
}

func GetFlagsMeta() []parser.FlagMeta {
    out := make([]parser.FlagMeta, len(flagsMeta))
    copy(out, flagsMeta)
    return out
}