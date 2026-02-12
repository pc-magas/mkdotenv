package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func testExtractVariableNameExtractsVariableName(t *testing.T) {
	expected_variable_name:="VARIABLE"
	line:=expected_variable_name+"=VALUE"

	variable_name,err:=ExtractVariableName(line)
	
	if(err != nil){
		t.Errorf(err.Error())
	}

	assert.EqualValues(t,expected_variable_name,variable_name)

}
