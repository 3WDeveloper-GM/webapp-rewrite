package config

import (
	"net/http"
	"time"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/templating"
)

func (app *Application) CurrentYearTemplateData(r *http.Request) *templating.TemplateData {
	return &templating.TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
