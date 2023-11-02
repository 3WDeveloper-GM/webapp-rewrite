//This file is the helpers component of the config package. I'll code the most common error prompts
//that will arise when dealing with errors in the web application.

package config

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration/templating"
	"github.com/go-playground/form/v4"
)

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.Errorlog.Output(2, trace)
	if app.DebugMode {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) DecodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.FormDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}

func (app *Application) Render(w http.ResponseWriter, status int, page string, data *templating.TemplateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *Application) IsAuthenticated(r *http.Request) bool {
	isAutheticated, ok := r.Context().Value(IsAuthenticatedContextkey).(bool)
	if !ok {
		return false
	}
	return isAutheticated
}
