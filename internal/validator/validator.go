package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// Validator holds a map of form fields to validation error messages.
type Validator struct {
	FieldErrors map[string]string
}

// Valid returns true if v contains no errors.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) addFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField adds an error message to a field if the provided `ok` validation check is false.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.addFieldError(key, message)
	}
}

// NotBlank returns true if a value is not an empty string.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars returns true if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue returns true if a value is in a list of specific permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
