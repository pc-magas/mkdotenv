package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseMkdotenvCommentIsParsedCorrecly(t *testing.T) {
	line:="#mkdotenv():resolve(\"value\"):plain(value=\"value\")"
	value:=ParseMkDotenvComment(line, map[string]string{})

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"",value.Item,"Item should be an empry String")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestParseMkdotenvCommentWithItem(t *testing.T){
	line:="#mkdotenv():resolve(\"value\"):plain(value=\"value\").item"

	value:=ParseMkDotenvComment(line, map[string]string{})

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestParseMkdotenvExtractsEnvironment(t *testing.T){
	line:="#mkdotenv(prod):resolve(\"value\"):plain(value=\"value\").item"

	value:=ParseMkDotenvComment(line, map[string]string{})

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")

}

func TestParseMkdotenvMultipleArguments(t *testing.T){
	line:="#mkdotenv(prod):resolve(\"value\"):plain(value=\"value\",value1='value',value2=value).item"
	
	value:=ParseMkDotenvComment(line, map[string]string{})

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\"","value1":"'value'","value2":"value"})
}

func TestParseMkdotenvParseNormalLines(t *testing.T){
	line:="hello"
	value:=ParseMkDotenvComment(line, map[string]string{})

	if(value != nil){
		t.Fatalf("expected value to be nil, got %v", value)
	}
}

func TestParseMkdotenvParseNormalArg(t *testing.T){
	line:="# hello"
	value:=ParseMkDotenvComment(line, map[string]string{})

	if(value != nil){
		t.Fatalf("expected value to be nil, got %v", value)
	}
}

func TestParseInvalidMkdotenv(t *testing.T){
	// Emulatins OS arguments first one is executable
	testCases := []string{
		"#mkdotenv(sadsadsada)",
		"#mkdotenv(sadsadsada):dsadsadsadsa:dsadada.dsa",
		"",                                   // empty line
		"#mkdotenv",                         // missing parentheses
		"#mkdotenv()",                       // missing resolver (::...)
		"#mkdotenv():resolve(\"value\"):",                     // missing resolver name
		"#mkdotenv(prod):resolve(\"value\"):vault",            // missing parentheses after resolver
		"#mkdotenv(prod):resolve(\"value\"):vault(",           // unclosed parentheses
		"#mkdotenv(prod):resolve(\"value\"):vault)(",          // parentheses mismatch
		"#mkdotenv(prod):resolve(\"value\"):vault(access_key=foo", // unclosed arg list
		"mkdotenv(prod):resolve(\"value\"):vault()",           // missing leading '#'
		"#mkdotnev(prod):resolve(\"value\"):vault()",          // misspelled command
		"#MKDOTENV(prod):resolve(\"value\"):vault()",          // uppercase should fail (regex is case-sensitive)
		"#mkdotenv(prod):resolve(\"value\"):vault(access_key=foo).", // dot but no item
		"#mkdotenv(prod):resolve(\"value\"):vault(access_key=foo)..secret", // double dots
		"#mkdotenv(prod):resolve(\"value\"):vault(access_key=foo).secret.extra", // extra dot section
		"#mkdotenv(prod)::vault(access_key=foo).secret", //missing resolve
		"#mkdotenv(prod):resolve:vault(access_key=foo).secret", //Resolve doies not contain ()
		"#mkdotenv(prod):resolve(:vault(access_key=foo).secret", //Non Closing brackents upon resolve
		"#mkdotenv(prod):resolve(\"Value\":vault(access_key=foo).secret", //Non Closing brackents upon resolve
		"#mkdotenv(prod):resolve\"Value\"):vault(access_key=foo).secret", //Non Closing brackents upon resolve
	}


	for _,line := range testCases {
		t.Run(line, func(t *testing.T) {
			value:=ParseMkDotenvComment(line,map[string]string{})
			if(value != nil){
				t.Fatalf("expected value to be nil, got %v", value)
			}
		})
	}
}

func TestGetArg(t *testing.T){
	value:="$_ARG[test]"
	expected_value:="test"

	result:=GetArg(value)

	assert.Equal(t,expected_value,result)
}