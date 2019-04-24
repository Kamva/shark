package sentry

import (
	"errors"

	"github.com/Kamva/shark/exceptions"
	"github.com/getsentry/raven-go"
)

// CaptureRoutineException capture routines exceptions and report it to sentry.
func CaptureRoutineException(exception []exceptions.RoutineException) {
	errString := "routine exception"
	packet := raven.NewPacket(
		errString,
		raven.NewException(errors.New(errString), raven.NewStacktrace(2, 3, nil)),
	)

	raven.Capture(packet, getRoutineExceptionTags(exception))
}

// CaptureMessage formats and delivers a string message to the Sentry server.
func CaptureMessage(message string, tags ...map[string]string) {
	t := make(map[string]string)
	for _, ts := range tags {
		for key, value := range ts {
			t[key] = value
		}
	}
	raven.CaptureMessage(message, t)
}

func getRoutineExceptionTags(exceptions []exceptions.RoutineException) map[string]string {
	var tag = make(map[string]string)

	for _, value := range exceptions {
		tag[value.RoutineName] = value.Message
	}

	return tag
}
