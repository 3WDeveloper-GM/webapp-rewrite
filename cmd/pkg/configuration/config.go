package config

import (
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/database"
	loggers "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/loggers"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
)

type Application struct {
	Infolog       *log.Logger
	Errorlog      *log.Logger
	Snippets      *models.Snippetmodel
	TemplateCache map[string]*template.Template
}

func GetApp() *Application {
	return &Application{}
}

func ExportConfig(dsn string, templates map[string]*template.Template) *Application {
	db, err := database.OpenDB(dsn)
	if err != nil {
		loggers.ErrorLog().Fatal(err)
	}
	app := &Application{
		Infolog:       loggers.Infolog(),
		Errorlog:      loggers.ErrorLog(),
		Snippets:      &models.Snippetmodel{DB: db},
		TemplateCache: templates,
	}
	return app
}
