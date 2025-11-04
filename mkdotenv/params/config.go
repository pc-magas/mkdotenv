package params

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
    AllowMultiple bool
	Validator   func(value string) bool
}

var FLAG_ARGUMENTS = []string{}

var flagsMeta = []FlagMeta{
    {
        Name:     "help",
        Aliases:  []string{},
		Short: "h",
        Required: false,
        Usage:    "Display help message and exit",
        AllowMultiple: false,
		Type:  	BoolType,
        Order:    0,
    },
    {
        Name:     "version",
        Aliases:  []string{},
		Short: "v",
        Required: false,
        AllowMultiple: false,
		Type:  	BoolType,
        Usage:    "Display version and exit",
        Order:    0,
    },
    {
        Name:     "environment",
        Aliases:  []string{"env"},
        AllowMultiple: false,
		Short: "e",
		Type: StringType,
        Required: false,
		DefaultValue: ".env,dist",
        Usage:    "Environment in which secrets would be resolved",
        Order:    1,
		Validator: valiDateCommon,
    },
    {
        Name:     "template-file",
        Aliases:  []string{"template"},
		Short: "t",
        AllowMultiple: false,
		Type: StringType,
        Required: false,
		DefaultValue: ".env.dist",
        Usage:    "Template .env file containing commands on how .env file would be generated",
        Order:    2,
		Validator: valiDateCommon,
    },
    {
        Name:     "output-file",
        Aliases:  []string{},
        Short: "o",
		Type: StringType,
        AllowMultiple: false,
        Required: false,
		DefaultValue: ".env",
        Usage:    "File to write output to",
        Order:    2,
		Validator: valiDateCommon,
    },
    {
        Name:     "argument",
        Aliases:  []string{"arg"},
		Short: "a",
        AllowMultiple: true,
		Type: StringType,
        Required: false,
		DefaultValue: "",
        Usage:    "Template .env file containing commands on how .env file would be generated",
        Order:    3,
		Validator: valiDateCommon,
    },
}