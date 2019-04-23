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

	raven.Capture(packet, getCaptureTags(exception))
}

func getCaptureTags(exceptions []exceptions.RoutineException) map[string]string {
	var tag = make(map[string]string)

	for _, value := range exceptions {
		tag[value.RoutineName] = value.Message
	}

	return tag
}
