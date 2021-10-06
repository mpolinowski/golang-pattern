package main

import (
	"fmt"
	"sync"
)

// Create a struct to be used as a singleton
type AppLogger struct {
	loglevel int
}

// Use logger to print a log message
func (l *AppLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

// Set the loglevel
func (l *AppLogger) SetLogLevel(level int) {
	l.loglevel = level
}

// Instantiate logger
var logger *AppLogger

// Sync package enforces concurrent goroutine safety
var once sync.Once


// Provide global access to logger instance
func  getLoggerInstance() *AppLogger {
	once.Do(func ()  {
		// Lazy create logger on first call
		fmt.Println("Creating logger instance")
		logger = &AppLogger{}	
	})
		
	fmt.Println("Providing logger instance")
	return logger
}