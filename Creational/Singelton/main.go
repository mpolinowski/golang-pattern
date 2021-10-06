package main

func main() {
	// Instantiate logger
	log := getLoggerInstance()
	// Set loglevel and log message
	log.SetLogLevel(1)
	log.Log("Log level 1 message")
	// Set loglevel and log message
	log = getLoggerInstance()
	log.SetLogLevel(2)
	log.Log("Log level 2 message")
	// Set loglevel and log message
	log = getLoggerInstance()
	log.SetLogLevel(3)
	log.Log("Log level 3 message")
}