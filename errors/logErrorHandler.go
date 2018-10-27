package errors

import "log"

//LogErrorHandler is a simple implementation of IHandleError that writes to the console.
type LogErrorHandler struct {
}

//LogError writes a panic message to the console.
//		log.Panicf will print out the error to the console and terminate execution with a pnic.
func (LogErrorHandler) LogError(err error, message string) {
	if err != nil {
		log.Panicf("ERROR %s: %s", message, err)
	}
	log.Panicf("ERROR %s", message)
}

//LogWarning writes a fatal message to the console.
//		log.Fatalf will print out the error to the console but will not terminate execution.
func (LogErrorHandler) LogWarning(message string) {
	log.Fatal(message)
}

//LogInformation writes an information message to the console.
func (LogErrorHandler) LogInformation(message string) {
	log.Printf("INFORMATION %s", message)
}

//LogVerbose writes a verbose message to the console.
func (LogErrorHandler) LogVerbose(message string) {
	log.Printf("VERBOSE %s", message)
}
