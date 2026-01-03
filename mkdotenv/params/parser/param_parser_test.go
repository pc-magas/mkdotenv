package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/pc-magas/mkdotenv/params/validate"
)

// test target struct that flags will populate
type testArgs struct {
	Name  string
	Flag bool
	ArgCount int
	Args []string
}

func TestParamParser_Parse(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","--name", expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}

func TestParamParser_ParseShort(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","-n", expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}

func TestParamParser_ParseValueSeperatedWithIson(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","-name="+expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}

func TestParamParser_ParseValueSeperatedWithIsonShort(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","-name="+expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}

func TestParamParser_ParseAlias(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{"onoma"},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","--onoma", expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}


func TestParamParser_ParseAliasIson(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{"onoma"},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "name":
				args.Name=value
		}
		return nil
	}

	complete,err:=parser.Parse([]string{"executable","--onoma="+expectedValue},&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,values.Name,expectedValue)
}

func TestParamParser_ParseFlag(t *testing.T) {

	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{},
			Short: "h",
			AllowMultiple: false,
			Type: NoValType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.Flag=true
		}
		return nil
	}

	args := []string{"executable","--help"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.True(t,values.Flag)
}

func TestParamParser_ParseFlagShort(t *testing.T) {

	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{},
			Short: "h",
			AllowMultiple: false,
			Type: NoValType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.Flag=true
		}
		return nil
	}

	args := []string{"executable","-h"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.True(t,values.Flag)
}

func TestParamParser_ParseFlagAlias(t *testing.T) {

	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{"voithia"},
			Short: "h",
			AllowMultiple: false,
			Type: NoValType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.Flag=true
		}
		return nil
	}

	args := []string{"executable","--voithia"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.True(t,values.Flag)
}

func TestParamParser_ParseFlagWithValueFails(t *testing.T) {

	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{"voithia"},
			Short: "h",
			AllowMultiple: false,
			Type: NoValType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.Flag=true
		}
		return nil
	}

	args := []string{"executable","--help=someval"}
	complete,err:=parser.Parse(args,&values);

	assert.Error(t,err)
	assert.False(t,complete)
	assert.False(t,values.Flag)
}

func TestParamParser_TestMultiple(t *testing.T) {
	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{"voithia"},
			Short: "h",
			AllowMultiple: true,
			Type: NoValType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		ArgCount:0,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.ArgCount=args.ArgCount+1
		}
		return nil
	}

	args := []string{"executable","--help","-h","--voithia","-h"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,4,values.ArgCount)
}

func TestParamParser_TestMultipleString(t *testing.T) {
	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{"voithia"},
			Short: "h",
			AllowMultiple: true,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Args:[]string{},
	}

	expectedArgs:=[]string{"1","2","3","4"}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "help":
				args.Args=append(args.Args,value)
		}
		return nil
	}

	args := []string{"executable","--help","1","-h","2","--voithia","3","-h","4"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,expectedArgs,values.Args)
}

func TestParamParser_TestMultipleStringWithOtherflags(t *testing.T) {
	flags := FlagList{
		{
			Name:     "argument",
			Aliases:  []string{"arg"},
			Short: "a",
			AllowMultiple: true,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
		{
			Name:     "name",
			Aliases:  []string{"onoma"},
			Short: "n",
			AllowMultiple: true,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: validate.ValidateCommon,
		},
	}

	values:=testArgs{
		Args:[]string{},
	}

	expectedArgs:=[]string{"1","2","3","4"}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "argument":
				args.Args=append(args.Args,value)
		}
		return nil
	}

	args := []string{"executable","--argument","1","--arg","2","-a","3","--name=alexandros","-a","4"}
	complete,err:=parser.Parse(args,&values);

	assert.NoError(t,err)
	assert.True(t,complete)
	assert.Equal(t,expectedArgs,values.Args)
}

func TestParamParser_ValidatorFails(t *testing.T) { 
	flags := FlagList{
		{
			Name:     "argument",
			Aliases:  []string{"arg"},
			Short: "a",
			AllowMultiple: false,
			Type: StringType,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator:  func(value string) bool {
				return false
			},
		},
	}

	values:=testArgs{
		Args:[]string{},
	}

	expectedArgs:=[]string{}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		switch meta.Name {
			case "argument":
				args.Args=append(args.Args,value)
		}
		return nil
	}

	args := []string{"executable","--argument","1"}
	complete,err:=parser.Parse(args,&values);

	assert.Error(t,err)
	assert.False(t,complete)
	assert.Equal(t,expectedArgs,values.Args)
}


