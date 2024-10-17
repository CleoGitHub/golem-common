package mvalidator

import (
	"errors"
	"testing"
)

func TestIsValidatorErrorsCorrect(t *testing.T) {
	validorError := make(ApiErrors, 1)
	validorError[0] = ApiError{
		Msg:   "Message for test validator",
		Field: "Field",
	}
	err := errors.New(validorError.ToJson())
	if !IsValidatorErrors(err) {
		t.Errorf("IsValidatorError should return true")
	}
}

func TestIsValidatorErrorsNotCorrect(t *testing.T) {
	err := errors.New("test normal error")
	if IsValidatorErrors(err) {
		t.Errorf("IsValidatorError should return true")
	}
}
