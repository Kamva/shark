package exceptions

import (
	"github.com/Kamva/pantopoda/http"
	"github.com/Kamva/pantopoda/http/api"
)

// Exception is an exception for fatal error happened in runtime
type ClientErrorException struct {
	code    string
	status  http.StatusCode
	payload api.Payload
	headers []api.ResponseHeader
}

// GetCode returns exception code
func (e ClientErrorException) GetCode() string {
	return e.code
}

// GetStatus returns http status code
func (e ClientErrorException) GetStatus() http.StatusCode {
	return e.status
}

// GetMessage returns exception end user readable Message
func (e ClientErrorException) GetMessage() string {
	return e.payload.Message
}

// GetPayload returns payload of client error response
func (e ClientErrorException) GetPayload() api.Payload {
	return e.payload
}

// GetHeaders returns headers for client error response
func (e ClientErrorException) GetHeaders() []api.ResponseHeader {
	return e.headers
}

// ThrowClientErrorException panic with an client error exception object filled
// with given data
func ThrowClientErrorException(code string, status http.StatusCode, payload api.Payload, headers ...api.ResponseHeader) {
	panic(ClientErrorException{code: code, status: status, payload: payload, headers: headers})
}
