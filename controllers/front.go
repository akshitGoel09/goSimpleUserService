package controllers

import "net/http"

//RegisterControllers route the request to correct http request handler
func RegisterControllers() {

	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
