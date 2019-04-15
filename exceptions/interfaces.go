package exceptions

import (
	"github.com/Kamva/pantopoda/http"
	"github.com/Kamva/pantopoda/http/api"
)

// GlobalException is an interface for any type of Exceptions
type GlobalException interface {
	// GetCode returns exception code
	GetCode() string

	// GetStatus returns http status code
	GetStatus() http.StatusCode

	// GetMessage returns exception end user readable Message
	GetMessage() string
}

// GenericException is an interface for any server Exceptions
type GenericException interface {
	GlobalException

	// GetErrorMessage returns exception error Message
	GetErrorMessage() string

	// ShouldReport determine that whether the panic should be reported
	ShouldReport() bool

	// GetTags returns list of informational tags for reporting
	GetTags() map[string]string
}

// ClientException is an interface for any client Exceptions
type ClientException interface {
	GlobalException

	// GetPayload returns payload of client error response
	GetPayload() api.Payload

	// GetHeaders returns headers for client error response
	GetHeaders() []api.ResponseHeader
}
