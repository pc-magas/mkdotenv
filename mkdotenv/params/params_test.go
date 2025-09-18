package params

import "testing"
import "fmt"

func TestValidParams(t *testing.T){

	expected_dotenv_filename:="xxxx"
	expected_output_filename:="zzzz"
	expected_variable_name:="123"
	expected_variable_value:="XXXX"

	// Emulatins OS arguments first one is executable
	testCases := []struct {
		args []string
	}{
		{[]string{"exec", "--variable-name","123","--variable-value", "XXXX", "--input-file", "xxxx", "--output-file", "zzzz"}},
		{[]string{"exec", "--variable-name", "123","--variable-value", "XXXX", "--env-file", "xxxx", "--output-file", "zzzz"}},
		{[]string{"exec", "--variable-name", "123","--variable-value", "XXXX", "--input-file=xxxx", "--output-file", "zzzz"}},
		{[]string{"exec", "--variable-name", "123","--variable-value=XXXX", "--input-file", "xxxx", "--output-file=zzzz"}},
		{[]string{"exec", "--variable-name=123","--variable-value", "XXXX", "--input-file=xxxx", "--output-file=zzzz"}},
	}

	for _, tc := range testCases {

		t.Run(tc.args[1], func(t *testing.T) { // Creates subtests
		 			
			err,argumentStruct := GetParameters(tc.args)
			
			fmt.Println(err)

			if err != nil {
				t.Errorf("Error should be nil")
				t.Error(err)
			}

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
		})
	}
}

func TestMissingInputFileAndOutputFile(t *testing.T){
	arguments:=[]string{"exec","--variable-name","123","--variable-value", "XXXX"}

	err,argumentStruct := GetParameters(arguments)

	if err != nil {
		t.Errorf("Error should be nil")
		t.Error(err)
	}

	if(argumentStruct.ParseComplete == false){
		t.Errorf("argument parsing is expected to be complete")
	}

	expected_dotenv_filename:=".env"
	var expected_output_filename string = ".env"
	expected_variable_name:="123"
	expected_variable_value:="XXXX"

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
		t.Errorf("argument parsin is expected to be complete")
	}

}

func TestMissingParams(t *testing.T){
	
	testCases := []struct {
		args []string
	}{
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file","--output-file"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=","--output-file="}},
	}

	for _, tc := range testCases {

		t.Run(tc.args[1], func(t *testing.T) { 

			err,_ := GetParameters(tc.args)
			if err == nil {
				fmt.Println(err)
				t.Errorf("Error should not be nil")
			}
		})
	}

}
func TestDuplicateArguments(t *testing.T){
	testCases := []struct {
		args []string
	}{
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--env-file","EEEE","--output-file=SSS","--output-file","SSS"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--env-file=EEEE","--output-file=SSS","--output-file","SSS"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--env-file","EEEE","--output-file","SSS","--output-file","SSS"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--env-file","EEEE","--output-file","SSS","--output-file","SSS","--output-file=SSS"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--env-file","EEEE","--output-file","SSS"}},
		{[]string{"exec","--variable-name","123","--variable-value", "XXXX","--input-file=XXXX","--input-file","XXXX","--env-file=QQQQ","--env-file","EEEE","--output-file","SSS"}},
	}

	for _, tc := range testCases {

		t.Run(tc.args[1], func(t *testing.T) { 

			err,_ := GetParameters(tc.args)
			
			if err == nil {
				t.Errorf("Error should not be nil")
			}
		})
	}

}
