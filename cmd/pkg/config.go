package config

import (
	"log"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/loggers"
)

type application struct {
	Infolog  *log.Logger
	Errorlog *log.Logger
}

func GetApp() *application {
	return &application{}
}

func (app *application) defineErrorLog() {
	app.Errorlog = loggers.ErrorLog()
}

func (app *application) defineInfoLog() {
	app.Infolog = loggers.Infolog()
}

func ExportLoggers() *application {
	app := GetApp()
	app.defineErrorLog()
	app.defineInfoLog()
	return app
}
