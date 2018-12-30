package exceptions

import "github.com/Kamva/pantopoda/http"

// RoutineException is an exception that happened in routines
type RoutineException struct {
	Message     string
	RoutineName string
	Critical    bool
}

// AggregatedRoutineException is aggregate of routine exceptions of a single event
type AggregatedRoutineException struct {
	message      string
	errorMessage string
	code         string
	status       http.StatusCode
	errors       []RoutineException
}

// NewAggregatedRoutineException generate an aggregate routine exception
func NewAggregatedRoutineException(
	message string,
	errorMessage string,
	code string, status http.StatusCode,
	errors []RoutineException,
) AggregatedRoutineException {
	return AggregatedRoutineException{
		message:      message,
		errorMessage: errorMessage,
		code:         code,
		status:       status,
		errors:       errors,
	}
}

// GetCode returns exception code
func (e AggregatedRoutineException) GetCode() string {
	return e.code
}

// GetStatus returns http status code
func (e AggregatedRoutineException) GetStatus() http.StatusCode {
	return e.status
}

// GetMessage returns exception end user readable Message
func (e AggregatedRoutineException) GetMessage() string {
	return e.message
}

// GetErrorMessage returns exception error Message
func (e AggregatedRoutineException) GetErrorMessage() string {
	return e.errorMessage
}

// ShouldReport determine that whether the panic should be reported
func (e AggregatedRoutineException) ShouldReport() bool {
	return true
}

// GetTags returns list of informational tags for reporting
func (e AggregatedRoutineException) GetTags() map[string]string {
	var tags = make(map[string]string)

	for _, exc := range e.errors {
		tags[exc.RoutineName] = exc.Message
	}

	return tags
}
