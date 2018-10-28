package logs

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
