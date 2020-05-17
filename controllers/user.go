package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller 🤑 😄"))
}

//constructor
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

/*
	package - like a namespace
	bind to userController line 7 (similar to this)
	write back in byte

	create consturtor using line 16 and point to stuct line 8
	line 20, defind patter /users/id

*/
