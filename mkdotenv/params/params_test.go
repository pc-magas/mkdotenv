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

		dotenv_filename,output_file,variable_name,variable_value := GetParameters(args)
		
		if dotenv_filename != expected_dotenv_filename {
			t.Errorf("Expected dotenv_filename to be '%s', but got '%s'", expected_dotenv_filename, dotenv_filename)
		}

		if output_file != expected_output_filename {
			t.Errorf("Expected output_filename to be '%s', but got '%s'", expected_output_filename, output_file)
		}

		if variable_name != expected_variable_name {
			t.Errorf("Expected variable_name to be '%s', but got '%s'", expected_variable_name, variable_name)
		}

		if variable_value != expected_variable_value {
			t.Errorf("Expected variable_value to be '%s', but got '%s'", expected_variable_value, variable_value)
		}
	}

}