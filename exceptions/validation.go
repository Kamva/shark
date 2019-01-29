package exceptions

import (
	"github.com/Kamva/shark"
)

// ValidationException is an exception for validation errors
type ValidationException struct {
	Exception
	errorBag shark.ErrorBag
}

// GetErrors return validation faults map
func (e ValidationException) GetErrors() map[string][]string {
	return e.errorBag.GetErrors()
}

// ThrowValidationException panics a validation exception with given data
func ThrowValidationException(code string, message string, errorBag shark.ErrorBag) {
	exception := Exception{code: code, message: message}
	panic(ValidationException{Exception: exception, errorBag: errorBag})
}
