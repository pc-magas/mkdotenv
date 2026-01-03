package params

import (
    "fmt"
    "strings"
)

type CLIArgType string

type FlagList []FlagMeta

const (
	StringType CLIArgType = "string"
	NoValType   CLIArgType = "noval"
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
    OnAssign OnAssignCallback[T]
}

func NewParamParser[T any](flags FlagList) *ParamParser[T] {
    p := &ParamParser[T]{
        FlagsMeta: flags,
        ParsedFlags: make(map[string]int),
    }

    return p
}

func (p *ParamParser[T]) extractArgumentAndValue(argument string) (string,string) {

    var value string

    if strings.Contains(argument, "=") {
        parts := strings.SplitN(argument, "=", 2)
        argument = parts[0]
        value = parts[1]
    }

    argument = strings.TrimLeft(argument, "-")

    return argument,value

}

func (p *ParamParser[T]) extractValueFromNext(i int, osArgs []string) string {

    length := len(osArgs)
    next := i+1;

    if(next > length){
        return ""
    }

    return osArgs[next]

}

func (p *ParamParser[T]) Parse(osArgs []string,target *T) (bool,error) {
    
    if(p.OnAssign == nil){
        return false,fmt.Errorf("No callback has been provided to assign the values")
    }

    var parseErr error

    for i, arg := range osArgs {
        
        if(i==0){
            continue;
        }

        if !strings.HasPrefix(arg, "-") {
            continue
        }

        name,value := p.extractArgumentAndValue(arg)
        meta := p.SearchFlag(name)

        if meta == nil {
            parseErr = fmt.Errorf("unknown flag: %s", arg)
            return false, parseErr
        }

        p.ParsedFlags[meta.Name]++
        if p.ParsedFlags[meta.Name] > 1 && !meta.AllowMultiple {
            parseErr = fmt.Errorf("flag --%s provided multiple times", meta.Name)
            return false, parseErr
        }

        if(meta.Type == StringType){
            if(value == ""){
                value = p.extractValueFromNext(i,osArgs)
            }
        } else if (meta.Type == NoValType){
            
            if(value != ""){
               parseErr = fmt.Errorf("flag --%s does not require value to be provided upon. %s provided", meta.Name,value)
                return false, parseErr
            }

            value="true"
        }
        
        if(meta.Validator!=nil && !meta.Validator(value)){
            parseErr = fmt.Errorf("flag --%s contains Invalid value. %s provided", meta.Name,value)
            return false,parseErr
        }

        if err := p.OnAssign(*meta, value, target); err != nil {
            return false,err
        }
    }

    return true, parseErr
}

func (p *ParamParser[T])SearchFlag(name string) *FlagMeta {
	name = strings.Trim(name,"-")
    for i := range p.FlagsMeta {
        if p.FlagsMeta[i].Name == name || p.FlagsMeta[i].Short == name  {
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