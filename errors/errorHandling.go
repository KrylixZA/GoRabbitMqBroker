package errors

import "log"

//IHandleError provides a contract for dealing with errors and warnings.
//By using an interface, the underlying implementation can be changed and injected to any users of IHandleError.
//		LogError abstracts the details of logging an error message.
//		LogWarning abstracts the details of logging a warning message. An error can be passed through if the user feels the error can be regarded as warning.
type IHandleError interface {
	LogError(err error, message string)
	LogWarning(err error, message string)
}

//LogErrorHandler is a simple implementation of IHandleError that writes to the console.
type LogErrorHandler struct {
}

//LogError writes a panic message to the console.
//		log.Panicf will print out the error to the console and terminate execution with a pnic.
func (LogErrorHandler) LogError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}

	log.Panic(msg)
}

//LogWarning writes a fatal message to the console.
//		log.Fatalf will print out the error to the console but will not terminate execution.
func (LogErrorHandler) LogWarning(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

	log.Fatal(msg)
}
