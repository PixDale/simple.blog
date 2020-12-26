package controllers

import (
	"net/http"

	"simple.blog/api/responses"
)

// Home is a simple response to serve the request
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
