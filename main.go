package main

import (
	"net/http"

	"github.com/akshitGoel09/goSimpleUserService/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
