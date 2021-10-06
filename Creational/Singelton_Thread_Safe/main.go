package main

import "fmt"

func main() {
	// Create several routines that all try to get an instance of logger to check concurret safety
	for i := 1; i < 10; i++ {
		getLoggerInstance()
	}
	fmt.Scanln()
}