package hw09structvalidator

import (
	"fmt"
	"strings"
)

type IllegalArgumentError struct {
	Message string
}

func (i IllegalArgumentError) Error() string {
	return fmt.Sprintf("illegal argument: %s", i.Message)
}

func NewIllegalArgumentError(message string) IllegalArgumentError {
	return IllegalArgumentError{
		Message: message,
	}
}

type FieldValidationError struct {
	Message string
}

func (f FieldValidationError) Error() string {
	return f.Message
}

func NewFieldValidationError(message string) FieldValidationError {
	return FieldValidationError{
		Message: message,
	}
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	sb.Write([]byte("Validation Errors:\n"))
	for _, err := range v {
		sb.Write([]byte(fmt.Sprintf("Field: %s, Err: %s\n", err.Field, err.Err)))
	}
	return sb.String()
}

func NewValidationErrors(validationErrors ...ValidationError) ValidationErrors {
	return validationErrors
}
