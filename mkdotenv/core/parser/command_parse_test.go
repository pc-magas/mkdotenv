package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseMkdotenvCommentIsParsedCorrecly(t *testing.T) {
	line:="#mkdotenv()::plain(value=\"value\")"
	value:=ParseMkDotenvComment(line)

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"",value.Item,"Item should be an empry String")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestParseMkdotenvCommentWithItem(t *testing.T){
	line:="#mkdotenv()::plain(value=\"value\").item"

	value:=ParseMkDotenvComment(line)

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestParseMkdotenvExtractsEnvironment(t *testing.T){
	line:="#mkdotenv(prod)::plain(value=\"value\").item"

	value:=ParseMkDotenvComment(line)

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")

}

func TestParseMkdotenvMultipleArguments(t *testing.T){
	line:="#mkdotenv(prod)::plain(value=\"value\",value1='value',value2=value).item"
	
	value:=ParseMkDotenvComment(line)

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\"","value1":"'value'","value2":"value"})
}

func TestParseMkdotenvParseNormalLines(t *testing.T){
	line:="hello"
	value:=ParseMkDotenvComment(line)

	if(value != nil){
		t.Fatalf("expected value to be nil, got %v", value)
	}
}

func TestParseMkdotenvParseNormalArg(t *testing.T){
	line:="# hello"
	value:=ParseMkDotenvComment(line)

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
		"#mkdotenv()::",                     // missing resolver name
		"#mkdotenv(prod)::vault",            // missing parentheses after resolver
		"#mkdotenv(prod)::vault(",           // unclosed parentheses
		"#mkdotenv(prod)::vault)(",          // parentheses mismatch
		"#mkdotenv(prod)::vault(access_key=foo", // unclosed arg list
		"mkdotenv(prod)::vault()",           // missing leading '#'
		"#mkdotnev(prod)::vault()",          // misspelled command
		"#MKDOTENV(prod)::vault()",          // uppercase should fail (regex is case-sensitive)
		"#mkdotenv(prod)::vault(access_key=foo).", // dot but no item
		"#mkdotenv(prod)::vault(access_key=foo)..secret", // double dots
		"#mkdotenv(prod)::vault(access_key=foo).secret.extra", // extra dot section
	}


	for _,line := range testCases {
		t.Run(line, func(t *testing.T) {
			value:=ParseMkDotenvComment(line)
			if(value != nil){
				t.Fatalf("expected value to be nil, got %v", value)
			}
		})
	}
}