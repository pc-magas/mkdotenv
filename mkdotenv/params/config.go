package params

import (
    "github.com/pc-magas/mkdotenv/params/parser"
    "github.com/pc-magas/mkdotenv/params/validate"
)

var flagsMeta = parser.FlagList{
    {
        Name:     "help",
        Aliases:  []string{},
		Short: "h",
        Usage:    "Display help message and exit",
        AllowMultiple: false,
		Type:  	parser.NoValType,
        Required: false,
        Order:    0,
    },
    {
        Name:     "version",
        Aliases:  []string{},
		Short: "v",
        AllowMultiple: false,
		Type:  	parser.NoValType,
        Usage:    "Display version and exit",
        Required: false,
        Order:    0,
    },
    {
        Name:     "environment",
        Aliases:  []string{"env"},
        AllowMultiple: false,
		Short: "e",
		Type: parser.StringType,
		DefaultValue: "default",
        Usage:    "Environment in which secrets would be resolved",
        Order:    1,
        Required: false,
		Validator: validate.ValidateCommon,
    },
    {
        Name:     "template-file",
        Aliases:  []string{"template"},
		Short: "t",
        AllowMultiple: false,
		Type: parser.StringType,
		DefaultValue: ".env.dist",
        Usage:    "Template .env file containing commands on how .env file would be generated. If no vanue provided .env.dist assumed.",
        Order:    2,
        Required: false,
		Validator: validate.ValidateExistingFile,
    },
    {
        Name:     "output-file",
        Aliases:  []string{},
        Short: "o",
		Type: parser.StringType,
        AllowMultiple: false,
		DefaultValue: ".env",
        Usage:    "File to write output to",
        Order:    2,
        Required: false,
		Validator: validate.ValidateCommon,
    },
    {
        Name:     "argument",
        Aliases:  []string{"arg"},
		Short: "a",
        AllowMultiple: true,
		Type: parser.StringType,
		DefaultValue: "",
        Usage:    "Argument provided as $_ARG upon template file.",
        Order:    3,
        Required: false,
		Validator: validate.ValidateCommon,
    },
}