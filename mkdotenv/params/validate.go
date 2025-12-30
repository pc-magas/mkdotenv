package params

import (
    "os"
)

func ValidateCommon(value string) bool {
	if(value == ""){
		return false
	}

	return true
}

func ValidateExistingFile(value string) bool {
	validate := ValidateCommon(value)
	if(validate == false){
		return false
	}

	info, err := os.Stat(value)
    if err != nil {
        return false
    }

	return !info.IsDir()
}