package config

import (
	"net/http"
	"time"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/templating"
)

func (app *Application) GetTemplateData(r *http.Request) *templating.TemplateData {
	return &templating.TemplateData{
		CurrentYear: time.Now().Year(),

		Flash: app.SessionManager.PopString(r.Context(), "flash"),
	}
}
