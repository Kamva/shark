package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/Kamva/shark/exceptions"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// ErrorReporter creates and returns a new error reporter middleware.
func ErrorReporter() context.Handler {
	return func(context context.Context) {
		request := context.Request()

		defer func() {
			// Check if any panic happened during the process
			if err := recover(); err != nil {
				// Set err value in context for error handling in next middleware
				context.Values().Set("exception", err)

				var reportMessage string
				var reportTags map[string]string

				if e, ok := err.(exceptions.GenericException); ok {
					if !e.ShouldReport() {
						return
					}
					reportMessage = e.GetErrorMessage()
					reportTags = e.GetTags()
				} else {
					reportMessage = fmt.Sprint(err)
					reportTags = map[string]string{"exceptions": "unknown", "type": fmt.Sprintf("%T", err)}
				}

				reportErr(reportMessage, request, reportTags)

				logWarning(context, reportMessage)
			}
		}()

		context.Next()
	}
}

func reportErr(reportMessage string, request *http.Request, reportTags map[string]string) {
	packet := raven.NewPacket(
		reportMessage,
		raven.NewException(errors.New(reportMessage), raven.NewStacktrace(2, 3, nil)),
		raven.NewHttp(request),
	)
	raven.Capture(packet, reportTags)
}

func logWarning(context iris.Context, err string) {
	logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", context.HandlerName())
	logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(context))
	logMessage += fmt.Sprintf("Trace: %s\n", err)
	logMessage += fmt.Sprintf("\n%s", getStacktrace())
	context.Application().Logger().Warn(logMessage)
}

func getStacktrace() string {
	var stacktrace string
	for i := 1; ; i++ {
		_, f, l, got := runtime.Caller(i)
		if !got {
			break

		}

		stacktrace += fmt.Sprintf("%s:%d\n", f, l)
	}

	return stacktrace
}

func getRequestLogs(context context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(context.GetStatusCode())
	path = context.Path()
	method = context.Method()
	ip = context.RemoteAddr()

	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}
