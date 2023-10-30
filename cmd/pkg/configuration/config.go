package config

import (
	"html/template"
	"log"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/database"
	loggers "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/loggers"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
)

type Application struct {
	Infolog        *log.Logger
	Errorlog       *log.Logger
	Snippets       *models.Snippetmodel
	TemplateCache  map[string]*template.Template
	FormDecoder    *form.Decoder
	SessionManager *scs.SessionManager
}

func GetApp() *Application {
	return &Application{}
}

func ExportConfig(dsn string, templates map[string]*template.Template) *Application {

	// Database configuration at compile time
	db, err := database.OpenDB(dsn)
	if err != nil {
		loggers.ErrorLog().Fatal(err)
	}

	//Form decoder configuration
	formDecoder := form.NewDecoder()

	sman := scs.New()
	sman.Store = mysqlstore.New(db)
	sman.Lifetime = 12 * time.Hour

	//Generates an application object that is exported to the main function
	app := &Application{
		Infolog:        loggers.Infolog(),
		Errorlog:       loggers.ErrorLog(),
		Snippets:       &models.Snippetmodel{DB: db},
		TemplateCache:  templates,
		FormDecoder:    formDecoder,
		SessionManager: sman,
	}
	return app
}
