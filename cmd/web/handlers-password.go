package main

import (
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	internal "github.com/3WDeveloper-GM/webapp-rewrite/internal/validator"
)

type userpasswordForm struct {
	CurrentPassword       string `form:"password"`
	NewPassWord           string `form:"password"`
	NewPassWordValidation string `form:"password"`
	internal.Validator
}

func UpdatePasswordPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var formdata userpasswordForm

		err := app.DecodePostForm(r, &formdata)
		if err != nil {
			app.ServerError(w, err)
		}

		formdata.CheckField(internal.NotBlank(formdata.CurrentPassword), "currentpassword", "This field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.NewPassWord), "newpassword", "This field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.NewPassWordValidation), "newpasswordvalidation", "This field cannot be blank")

		formdata.CheckField(internal.PasswordMatching(formdata.NewPassWord, formdata.NewPassWordValidation), "newpassword", "Both passwords must match")

		if !formdata.Valid() {
			data := app.GetTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusUnprocessableEntity, "update.tmpl", data)
		}

	}
}

func UpdatePasswordGet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.GetTemplateData(r)
		app.Render(w, http.StatusOK, "update.tmpl", data)
	}
}
