package middleware

import (
	"fmt"
	"os"

	"github.com/Kamva/pantopoda/http/api"
	"github.com/Kamva/shark/exceptions"
	"github.com/kataras/iris/v12"
)

// ErrorRenderer renders errors and exception in json format
func ErrorRenderer() iris.Handler {
	return func(context iris.Context) {
		defer renderError(context)

		context.Next()
	}
}

func renderError(context iris.Context) {
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
		payload := api.Payload{}
		if os.Getenv("APP_ENV") != "production" {
			payload.Message = e.GetMessage()
		} else {
			payload.Message = "INTERNAL ERROR"
		}

		response.Response(e.GetCode(), e.GetStatus(), payload)
	} else {
		payload := api.Payload{}
		if os.Getenv("APP_ENV") != "production" {
			payload.Message = fmt.Sprint(err)
		} else {
			payload.Message = "INTERNAL ERROR"
		}

		response.InternalServerError(exceptions.GENERAL, payload)
	}
}
