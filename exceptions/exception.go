package exceptions

import "github.com/Kamva/pantopoda/http"

// GENERAL is default status code
const GENERAL string = "GENERAL_ERROR"

// Exception is an exception for fatal error happened in runtime
type Exception struct {
	code         string
	status       http.StatusCode
	message      string
	errorMessage string
}

// GetCode returns exception code
func (e Exception) GetCode() string {
	if e.code == "" {
		return "GENERAL_ERROR"
	}

	return e.code
}

// GetStatus returns http status code
func (e Exception) GetStatus() http.StatusCode {
	return e.status
}

// GetMessage returns exception end user readable Message
func (e Exception) GetMessage() string {
	return e.message
}

// GetErrorMessage returns exception error Message
func (e Exception) GetErrorMessage() string {
	return e.errorMessage
}

// ShouldReport determine that whether the panic should be reported
// By default all exceptions returns true for ShouldReport method
func (e Exception) ShouldReport() bool {
	return true
}

// GetTags returns list of informational tags for reporting
func (e Exception) GetTags() map[string]string {
	return map[string]string{
		"exception": "global",
		"type":      "Exception",
		"code":      e.code,
	}
}

// ThrowException panic with an exception object filled with given data
func ThrowException(code string, status http.StatusCode, message string, errorMessage string) {
	panic(Exception{code: code, status: status, message: message, errorMessage: errorMessage})
}
