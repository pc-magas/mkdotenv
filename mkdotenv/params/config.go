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
	Validator    func(value string) bool
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
        Name:     "template-file",
        Aliases:  []string{"template"},
		Short: "t",
		Type: StringType,
        Required: false,
		DefaultValue: ".env",
        Usage:    "Template .env file containing commands on how .env file would be generated",
        Order:    2,
		Validator: valiDateCommon,
    },
    {
        Name:     "output-file",
        Aliases:  []string{},
		Type: StringType,
        Required: false,
		DefaultValue: ".env",
        Usage:    "File to write output to (`-` for stdout)",
        Order:    2,
		Validator: valiDateCommon,
    },
}