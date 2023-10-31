package config

import (
	"crypto/tls"
	"database/sql"
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
	TLSconfig      *tls.Config
}

func GetApp() *Application {
	return &Application{}
}

// This function initalizes a database configuration, I do the error handling
// "inside" of the function, I don't really care for the error in the case the
// function works, so at the compile time, the ExportConfig function just uses
// the first value of databaseConfig.
func databaseConfig(dsn string) (*sql.DB, error) {
	db, err := database.OpenDB(dsn)
	if err != nil {
		loggers.ErrorLog().Fatal(err)
		return nil, err
	}
	return db, nil
}

// This little function just initializes a *scs.SessionManager object to be
// used by the database and the application. It just saves me the effort to
// change some things in the session manager if i need.
func sessionManager(db *sql.DB) *scs.SessionManager {
	sman := scs.New()
	sman.Store = mysqlstore.New(db)
	sman.Lifetime = 12 * time.Hour
	return sman
}

// Establishes the preferences for the elliptic curves used by the app, the book says
// that this curves are good enough to get security and still have acceptable performance.
func tlsConfig() *tls.Config {
	return &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
	}
}

func ExportConfig(dsn string, templates map[string]*template.Template) *Application {

	// Database configuration at compile time
	db, _ := databaseConfig(dsn)

	//Form decoder configuration
	formDecoder := form.NewDecoder()

	appSessionManager := sessionManager(db)

	//Generates an application object that is exported to the main function
	app := &Application{
		Infolog:        loggers.Infolog(),
		Errorlog:       loggers.ErrorLog(),
		Snippets:       &models.Snippetmodel{DB: db},
		TemplateCache:  templates,
		FormDecoder:    formDecoder,
		SessionManager: appSessionManager,
		TLSconfig:      tlsConfig(),
	}
	return app
}
