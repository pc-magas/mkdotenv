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

		err,arguments := GetParameters(args)
		
		if arguments.dotenv_filename != expected_dotenv_filename {
			t.Errorf("Expected dotenv_filename to be '%s', but got '%s'", expected_dotenv_filename, arguments.dotenv_filename)
		}

		if arguments.output_file != expected_output_filename {
			t.Errorf("Expected output_filename to be '%s', but got '%s'", expected_output_filename, arguments.output_file)
		}

		if arguments.variable_name != expected_variable_name {
			t.Errorf("Expected variable_name to be '%s', but got '%s'", expected_variable_name, arguments.variable_name)
		}

		if arguments.variable_value != expected_variable_value {
			t.Errorf("Expected variable_value to be '%s', but got '%s'", expected_variable_value, arguments.variable_value)
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

		callbackCalled := false // Flag to check if callback was called

		emptyCallback := func(msg string) {
			callbackCalled = true // Set flag when callback is called
		}

		_, _, _, _ = GetParameters(args, emptyCallback)

		// Ensure callback was called, indicating invalid parameters
		if !callbackCalled {
			t.Errorf("Expected emptyCallback to be called for args: %v", args)
		}
	}

}

func TestMissingInputFile(t *testing.T){
	arguments:=[][]string {
		{"exec","123","XXXX","--input-file=","--output-file","zzzz"},
		{"exec","123","XXXX","--input-file","--output-file","zzzz"},
	}

	for _, args := range arguments {

		callbackCalled := false // Flag to check if callback was called

		emptyCallback := func(msg string) {
			callbackCalled = true // Set flag when callback is called
		}

		_, _, _, _ = GetParameters(args, emptyCallback)

		// Ensure callback was called, indicating invalid parameters
		if !callbackCalled {
			t.Errorf("Expected emptyCallback to be called for args: %v", args)
		}
	}
}