package params

import(
	"errors"
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


func GetParameters(osArguments []string) (error,Arguments) {
	
	if len(osArguments) < 1 {
		return errors.New("not enough arguments provided"),Arguments{}
	}

	args := Arguments{
		Environment: "default",
		TemplateFile: ".env.dist",
		OutputFile: ".env",
		ArgumentNum: len(osArguments),
		MiscArguments: make(map[string]string),
		ParseComplete: false,
	}

    parser := NewParamParser[Arguments](flagsMeta)

	parser.OnAssign = func(meta FlagMeta, value string, args *Arguments) error {
        switch meta.Name {
		case "help":
			args.DisplayHelp=true
		case "version":
			args.DisplayVersion=true
		case "template-file":
			args.TemplateFile = value
		case "output-file":
			args.OutputFile = value
		default:
			args.MiscArguments[meta.Name] = value
        }
        return nil
    }

	parser.Parse(osArguments,&args)


	return nil,args
}

func GetFlagsMeta() []FlagMeta {
    out := make([]FlagMeta, len(flagsMeta))
    copy(out, flagsMeta)
    return out
}