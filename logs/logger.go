package logs

import "log"

//ILogger provides a contract for dealing with errors and warnings.
//By using an interface, the underlying implementation can be changed and injected to any users of ILogger.
//		LogError abstracts the details of logging an error message.
//		LogWarning abstracts the details of logging a warning message.
//		LogInformation abstracts the details of logging an information message.
//		LogVerbose abstracts the details of logging a verbose message.
type ILogger interface {
	LogError(err error, message string)
	LogWarning(message string)
	LogInformation(message string)
	LogVerbose(message string)
}

//Logger is a simple implementation of ILogger that writes to the console.
type Logger struct {
}

//LogError writes a panic message to the console.
//		log.Panicf will print out the error to the console and terminate execution with a pnic.
func (Logger) LogError(err error, message string) {
	if err != nil {
		log.Panicf("ERROR %s: %s", message, err)
	}
	log.Panicf("ERROR %s", message)
}

//LogWarning writes a fatal message to the console.
//		log.Fatalf will print out the error to the console but will not terminate execution.
func (Logger) LogWarning(message string) {
	log.Fatal(message)
}

//LogInformation writes an information message to the console.
func (Logger) LogInformation(message string) {
	log.Printf("INFORMATION %s", message)
}

//LogVerbose writes a verbose message to the console.
func (Logger) LogVerbose(message string) {
	log.Printf("VERBOSE %s", message)
}
