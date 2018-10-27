package errors

//IHandleError provides a contract for dealing with errors and warnings.
//By using an interface, the underlying implementation can be changed and injected to any users of IHandleError.
//		LogError abstracts the details of logging an error message.
//		LogWarning abstracts the details of logging a warning message. An error can be passed through if the user feels the error can be regarded as warning.
type IHandleError interface {
	LogError(err error, message string)
	LogWarning(message string)
	LogInformation(message string)
	LogVerbose(message string)
}
