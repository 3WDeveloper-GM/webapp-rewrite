package main

import (
	"log"
	"os"
)

func Infolog() *log.Logger {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	return logger
}

func ErrorLog() *log.Logger {
	logger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
