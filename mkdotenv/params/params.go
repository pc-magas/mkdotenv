package params

import(
	"strings"
	"errors"
)

type Arguments struct {
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

	parser.Parse(osArguments,args)


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