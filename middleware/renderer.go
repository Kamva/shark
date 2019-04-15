package middleware

import (
	"fmt"

	"github.com/Kamva/pantopoda/http/api"
	"github.com/Kamva/shark/exceptions"
	"github.com/kataras/iris/context"
)

// ErrorRenderer renders errors and exception in json format
func ErrorRenderer() context.Handler {
	return func(context context.Context) {
		defer renderError(context)

		context.Next()
	}
}

func renderError(context context.Context) {
	response := api.NewResponse(context)
	err := context.Values().Get("exception")

	if err == nil {
		return
	}

	if clientErr, ok := err.(exceptions.ClientException); ok {
		response.Response(clientErr.GetCode(), clientErr.GetStatus(), clientErr.GetPayload(), clientErr.GetHeaders()...)
	} else if validation, ok := err.(exceptions.ValidationException); ok {
		response.UnprocessableEntity(validation.GetCode(), api.Payload{
			Message: validation.GetMessage(),
			Data:    validation.GetErrors(),
		})
	} else if e, ok := err.(exceptions.GenericException); ok {
		response.Response(e.GetCode(), e.GetStatus(), api.Payload{
			Message: e.GetMessage(),
		})
	} else {
		response.InternalServerError(exceptions.GENERAL, api.Payload{
			Message: fmt.Sprint(err),
		})
	}
}
