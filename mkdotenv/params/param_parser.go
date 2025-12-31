package params

import "fmt"
import "strings"
import flag "github.com/spf13/pflag"

type CLIArgType string

type FlagList []FlagMeta

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
    AllowMultiple bool
	Validator   func(value string) bool
}

type OnAssignCallback[T any] func(meta FlagMeta, value string,target *T) error

type ParamParser[T any] struct {
    FlagsMeta FlagList
    ParsedFlags map[string]int
    FlagSet *flag.FlagSet
    OnAssign OnAssignCallback[T]
}

func (p *ParamParser[T]) initFlags() {
    p.FlagSet=flag.NewFlagSet("params", flag.ContinueOnError)
	for _, meta := range p.FlagsMeta {
        switch meta.Type {
			case StringType:
				
				if(meta.Short == ""){
					p.FlagSet.String(meta.Name, meta.DefaultValue, meta.Usage)
				} else {
					p.FlagSet.StringP(meta.Name, meta.Short, meta.DefaultValue, meta.Usage)
				}

				for _, alias := range meta.Aliases {
					p.FlagSet.String(alias, meta.DefaultValue, "(alias of --"+meta.Name+") "+meta.Usage)
				}

			case BoolType:
				def := meta.DefaultValue == "true"

				if(meta.Short == ""){
					p.FlagSet.Bool(meta.Name, def, meta.Usage)
				} else {
					p.FlagSet.BoolP(meta.Name,meta.Short, def, meta.Usage)
				}
				
				for _, alias := range meta.Aliases {
					p.FlagSet.Bool(alias, def, "(alias of --"+meta.Name+") "+meta.Usage)
				}
        }
    }
}

func NewParamParser[T any](flags FlagList) *ParamParser[T] {
    p := &ParamParser[T]{
        FlagsMeta: flags,
        ParsedFlags: make(map[string]int),
    }

    p.initFlags()
    return p
}

func (p *ParamParser[T]) Parse(osArgs []string,target *T) (bool,error) {
    err := p.FlagSet.Parse(osArgs)
    if err != nil {
        return false, err
    }

    var parseErr error

    p.FlagSet.Visit(func(f *flag.Flag) {
        meta := p.SearchFlag(f.Name)
        if meta == nil {
            parseErr = fmt.Errorf("unknown flag: %s", f.Name)
            return
        }

        // --- count occurrences ---
        p.ParsedFlags[meta.Name]++
        if p.ParsedFlags[meta.Name] > 1 && !meta.AllowMultiple {
            parseErr = fmt.Errorf("flag --%s provided multiple times", meta.Name)
            return 
        }

        value := f.Value.String()
        
        if(value == ""){
            value = meta.DefaultValue
        }

        if meta.Validator != nil && !meta.Validator(value) {
            parseErr = fmt.Errorf("invalid value for --%s", meta.Name)
            return
        }

        // --- assign to Arguments ---
        if p.OnAssign != nil {
            if err := p.OnAssign(*meta, value,target); err != nil {
                parseErr = err
                return
            }
        }

    })

    return true, parseErr
}

func (p *ParamParser[T])SearchFlag(name string) *FlagMeta {
	name = strings.Trim(name,"-")
    for i := range p.FlagsMeta {
        if p.FlagsMeta[i].Name == name {
            return &p.FlagsMeta[i]
        }
        for _, alias := range p.FlagsMeta[i].Aliases {
            if alias == name {
                return &p.FlagsMeta[i]
            }
        }
    }
    return nil
}