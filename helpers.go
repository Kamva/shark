package shark

// PanicIfError panic exception if the given error is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// PanicIfErrCustomMsg panic with custom message if given error is not nil
func PanicIfErrorWithMessage(err error, message string) {
	if err != nil {
		panic(message)
	}
}
