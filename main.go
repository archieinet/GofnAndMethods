package main

import (
	"net/http"

	"github.com/archieinet/GofnAndMethods/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)	
}
