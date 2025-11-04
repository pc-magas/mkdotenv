package params

import "fmt"
import flag "github.com/spf13/pflag"

type OnAssignCallback[T any] func(meta FlagMeta, value string,target *T) error

type ParamParser[T any] struct {
    FlagsMeta []FlagMeta
    ParsedFlags map[string]int
    FlagSet *flag.FlagSet
    OnAssign OnAssignCallback[T]
}

func (p *ParamParser[T]) initFlags() (*flag.FlagSet) {

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

func NewParamParser[T any](flags []FlagMeta) *ParamParser[T] {
    p := &ParamParser[T]{
        FlagsMeta: flags,
        ParsedFlags: make(map[string]int),
        FlagSet: flag.NewFlagSet("params", flag.ContinueOnError),
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
        meta := SearchFlag(f.Name)
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

        // --- validate value ---
        if meta.Validator != nil && !meta.Validator(f.Value.String()) {
            parseErr = fmt.Errorf("invalid value for --%s", meta.Name)
            return
        }

        // --- assign to Arguments ---
        if p.OnAssign != nil {
            if err := p.OnAssign(*meta, f.Value.String(),target); err != nil {
                parseErr = err
                return
            }
        }

    })

    return true, parseErr
}

func valiDateCommon(value string) bool {
	if(value == ""){
		return false
	}

	return true
}