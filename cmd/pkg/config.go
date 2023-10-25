package config

import (
	"log"

	loggers "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/loggers"
)

type Application struct {
	Infolog  *log.Logger
	Errorlog *log.Logger
}

func GetApp() *Application {
	return &Application{}
}

func (app *Application) defineErrorLog() {
	app.Errorlog = loggers.ErrorLog()
}

func (app *Application) defineInfoLog() {
	app.Infolog = loggers.Infolog()
}

func ExportLoggers() *Application {
	app := GetApp()
	app.defineErrorLog()
	app.defineInfoLog()
	return app
}
