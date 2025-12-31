package params

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// test target struct that flags will populate
type testArgs struct {
	Name  string
	Flag bool
}

func TestParamParser_Parse(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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

func TestParamParser_ParseAlias(t *testing.T) {
	flags := FlagList{
		{
			Name:     "name",
			Aliases:  []string{"onoma"},
			Short: "n",
			AllowMultiple: false,
			Type: StringType,
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Name:"wrong",
	}

	expectedValue:="David"

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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


func TestParamParser_ParseFlag(t *testing.T) {

	flags := FlagList{
		{
			Name:     "help",
			Aliases:  []string{},
			Short: "h",
			AllowMultiple: false,
			Type: BoolType,
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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
			Type: BoolType,
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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
			Type: BoolType,
			Required: false,
			DefaultValue: "",
			Usage:    "TestUsage",
			Order:    2,
			Validator: ValidateCommon,
		},
	}

	values:=testArgs{
		Flag:false,
	}

	parser := NewParamParser[testArgs](flags)
	parser.OnAssign = func(meta FlagMeta, value string, args *testArgs) error {
		fmt.Println(meta.Name,value)
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
