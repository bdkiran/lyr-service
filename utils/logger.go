package utils

import (
	"io"
	"log"
	"os"
)

//Logger is the logger structure that will provide logging fetures for the project
type Logger struct {
	//Info level logging
	Info *log.Logger
	//Warning level logging
	Warning *log.Logger
	//Error level logging
	Error *log.Logger
}

//defaultLogger is the logger that will be used across all packages
var defaultLogger *Logger

func init() {
	fileName := "app.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file : " + fileName)
	}

	multi := io.MultiWriter(file, os.Stdout)

	defaultLogger = new(Logger)
	defaultLogger.Info = log.New(multi,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	defaultLogger.Warning = log.New(multi,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	defaultLogger.Error = log.New(multi,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//NewLogger provides access to the gloabal logger to be used across the application
func NewLogger() *Logger {
	return defaultLogger
}
