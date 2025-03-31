package params

import "testing"

func TestValidParams(t *testing.T){

	expected_dotenv_filename:="xxxx"
	expected_output_filename:="zzzz"
	expected_variable_name:="123"
	expected_variable_value:="XXXX"

	// Emulatins OS arguments first one is executable
	arguments:=[][]string{
		{"exec","123","XXXX","--input-file","xxxx","--output-file","zzzz"},
		{"exec","123","XXXX","--input-file=xxxx","--output-file","zzzz"},
		{"exec","123","XXXX","--input-file","xxxx","--output-file=zzzz"},
		{"exec","123","XXXX","--input-file=xxxx","--output-file=zzzz"},
	}

	for _, args := range arguments {

		err,argumentStruct := GetParameters(args)
		
		if argumentStruct.DotenvFilename != expected_dotenv_filename {
			t.Errorf("Expected DotenvFilename to be '%s', but got '%s'", expected_dotenv_filename, argumentStruct.DotenvFilename)
		}

		if argumentStruct.OutputFile != expected_output_filename {
			t.Errorf("Expected output_filename to be '%s', but got '%s'", expected_output_filename, argumentStruct.OutputFile)
		}

		if argumentStruct.VariableName != expected_variable_name {
			t.Errorf("Expected variable_name to be '%s', but got '%s'", expected_variable_name, argumentStruct.VariableName)
		}

		if argumentStruct.VariableValue != expected_variable_value {
			t.Errorf("Expected variable_value to be '%s', but got '%s'", expected_variable_value, argumentStruct.VariableValue)
		}

		if(argumentStruct.ParseComplete == false){
			t.Errorf("argument parsing is expected to be complete")
		}

		if err != nil {
			t.Errorf("Error should be nil")
		}
	}
}

func TestMissingParams(t *testing.T){
	
	arguments:=[][]string{
		{"exec","123","XXXX","--input-file","--output-file"},
		{"exec","123","XXXX","--input-file","xxxx","--output-file="},
		{"exec","123","XXXX","--input-file=","--output-file="},
	}

	

	for _, args := range arguments {

		err,_ := GetParameters(args,)

		if err == nil {
			t.Errorf("Error should not be nil")
		}
	}

}

func TestMissingInputFile(t *testing.T){
	arguments:=[][]string {
		{"exec","123","XXXX","--input-file=","--output-file","zzzz"},
		{"exec","123","XXXX","--input-file","--output-file","zzzz"},
	}

	for _, args := range arguments {

		err,_:= GetParameters(args)

		if err == nil {
			t.Errorf("Error should not be nil")
		}
	}
}