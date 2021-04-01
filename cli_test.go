package main

import "testing"

func TestConfiguration_ParseSuccess(test *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			test.Error("[FAIL] TestConfiguration_ParseSuccess: Expected no error, got error.")
		}
	}()

	LoadConfiguration()
}

func TestParamsValidation_FileExists(test *testing.T) {
	fileName := "main.go"

	result := FileExists(fileName)

	if !result {
		test.Error("[FAIL] TestParamsValidation_FileExists: Expected true value, got false.")
	}
}

func TestParamsValidation_FileIsNotDirectory(test *testing.T) {
	fileName := "main.go"

	result := !FileIsDirectory(fileName)

	if !result {
		test.Error("[FAIL] TestParamsValidation_FileIsNotDirectory: Expected true value, got false.")
	}
}

func TestParamsValidation_NotEnoughParameters(test *testing.T) {
	args := [2]string{"script", "run"}
	fileName := "main.go"

	result := IsRunCommandCorrect(args[:], fileName)

	if result {
		test.Error("[FAIL] TestParamsValidation_NotEnoughParameters: Expected false value, got true.")
	}
}

func TestParamsValidation_NoFileFlag(test *testing.T) {
	args := [4]string{"script", "run", "somearg", "filee"}
	fileName := "main.go"

	result := IsRunCommandCorrect(args[:], fileName)

	if result {
		test.Error("[FAIL] TestParamsValidation_NoFileFlag: Expected false value, got true.")
	}
}

func TestParamsValidation_FileDoesntExist(test *testing.T) {
	args := [4]string{"script", "run", "--file", "<FILE>"}
	fileName := "fail"

	result := IsRunCommandCorrect(args[:], fileName)

	if result {
		test.Error("[FAIL] TestParamsValidation_FileDoesntExist: Expected false value, got true.")
	}
}

func TestParamsValidation_ValidParams(test *testing.T) {
	args := [4]string{"script", "run", "--file", "main.go"}
	fileName := "main.go"

	result := IsRunCommandCorrect(args[:], fileName)

	if !result {
		test.Error("[FAIL] TestParamsValidation_ValidParams: Expected true value, got false.")
	}
}
