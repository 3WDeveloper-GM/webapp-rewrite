package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/database"
	loggers "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/loggers"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
)

type Application struct {
	Infolog  *log.Logger
	Errorlog *log.Logger
	Snippets *models.Snippetmodel
}

func GetApp() *Application {
	return &Application{}
}

func ExportConfig(dsn string) *Application {
	db, err := database.OpenDB(dsn)
	if err != nil {
		loggers.ErrorLog().Fatal(err)
	}
	app := &Application{
		Infolog:  loggers.Infolog(),
		Errorlog: loggers.ErrorLog(),
		Snippets: &models.Snippetmodel{DB: db},
	}
	return app
}
