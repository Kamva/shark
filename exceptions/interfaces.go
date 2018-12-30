package exceptions

import "github.com/Kamva/pantopoda/http"

// GenericException is an interface for any type of Exceptions
type GenericException interface {
	// GetCode returns exception code
	GetCode() string

	// GetStatus returns http status code
	GetStatus() http.StatusCode

	// GetMessage returns exception end user readable Message
	GetMessage() string

	// GetErrorMessage returns exception error Message
	GetErrorMessage() string

	// ShouldReport determine that whether the panic should be reported
	ShouldReport() bool

	// GetTags returns list of informational tags for reporting
	GetTags() map[string]string
}
