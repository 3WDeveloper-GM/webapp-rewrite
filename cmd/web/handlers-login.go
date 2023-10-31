package main

import (
	"fmt"
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
)

func userSignup(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some web form for signup")
	}
}

func userSignupPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some method limited for posting")
	}
}

func userLogin(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some web form for login")
	}
}

func userLoginPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some method for login, limited to posting")
	}
}

func userLogoutPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "some method for logout, limited to post")
	}
}
