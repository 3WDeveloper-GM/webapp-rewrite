package templating

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/loggers"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
)

type TemplateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func HumanReadableDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humandate": HumanReadableDate,
}

//Some caching for more bien pinche fast

func NewTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		loggers.ErrorLog().Fatal(err)
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			loggers.ErrorLog().Print(err)
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			loggers.ErrorLog().Print(err)
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			loggers.ErrorLog().Print(err)
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
