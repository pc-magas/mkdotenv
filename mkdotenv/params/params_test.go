package params

import "testing"
import "fmt"

func TestValidParams(t *testing.T){

	expected_template_file:="xxxx"
	expected_output_filename:="zzzz"
	expected_environment:="124"
	// Emulatins OS arguments first one is executable
	testCases := []struct {
		args []string
	}{
		{[]string{"exec","--environment",expected_environment,"--template-file", expected_template_file, "--output-file", expected_output_filename}},
		{[]string{"exec","-e",expected_environment,"-t", expected_template_file, "-o", expected_output_filename}},
		{[]string{"exec","--environment",expected_environment,"-t", expected_template_file, "-o", expected_output_filename}},
		{[]string{"exec","--environment",expected_environment,"-t", expected_template_file, "--output-file", expected_output_filename}},
		{[]string{"exec","--environment",expected_environment,"--template-file", expected_template_file, "-o", expected_output_filename}},
	}

	for _, tc := range testCases {

		t.Run(tc.args[1], func(t *testing.T) { // Creates subtests
		 			
			err,argumentStruct := GetParameters(tc.args)
			
			fmt.Println(err)

			if err != nil {
				t.Errorf("Error should be nil")
				t.Error(err)
			}

			if argumentStruct.Environment != expected_environment {
				t.Errorf("Expected Environment to be '%s', but got '%s'", expected_environment, argumentStruct.Environment)
			}

			if argumentStruct.TemplateFile != expected_template_file {
				t.Errorf("Expected Template to be '%s', but got '%s'", expected_template_file, argumentStruct.TemplateFile)
			}

			if argumentStruct.OutputFile != expected_output_filename {
				t.Errorf("Expected output_filename to be '%s', but got '%s'", expected_output_filename, argumentStruct.OutputFile)
			}

			if(argumentStruct.ParseComplete == false){
				t.Errorf("argument parsing is expected to be complete")
			}
		})
	}
}
