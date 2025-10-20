package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMkdotenvCommentIsParsedCorrecly(t *testing.T) {
	line:="#mkdotenv()::plain(value=\"value\")"
	value,err:=ParseMkDotenvComment(line)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"",value.Item,"Item should be an empry String")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestMkdotenvCommentWithItem(t *testing.T){
	line:="#mkdotenv()::plain(value=\"value\").item"

	value,err:=ParseMkDotenvComment(line)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t,"default",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")
}

func TestMkDotenvExtractsEnvironment(t *testing.T){
	line:="#mkdotenv(prod)::plain(value=\"value\").item"

	value,err:=ParseMkDotenvComment(line)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\""},"Item should be an empry String")

}

func TestMkDotenvMultipleArguments(t *testing.T){
	line:="#mkdotenv(prod)::plain(value=\"value\",value1='value',value2=value).item"
	
	value,err:=ParseMkDotenvComment(line)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t,"prod",value.Environment,"Environment is not the expected one")
	assert.Equal(t,"plain",value.SecretResolverType,"Secret resolver is not the expected One")
	assert.Equal(t,"item",value.Item,"Item should equal with item")
	assert.Equal(t,value.Params, map[string]string{"value":"\"value\"","value1":"'value'","value2":"value"})
}