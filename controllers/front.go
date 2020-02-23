package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers route the request to correct http request handler
func RegisterControllers() {

	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
