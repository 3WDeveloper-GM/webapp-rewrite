package main

import "log"

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func GetApp() *application {
	return &application{}
}
