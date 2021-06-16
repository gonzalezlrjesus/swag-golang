package main

import (
	"net/http"

	_ "gonzalezlrjesus/swag-golang/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Jesus Gonzalez's Swagger Example API
// @version 0.1
// @description This is a sample to understand how swag works

// @contact.name Jesus Gonzalez
// @contact.url https://github.com/gonzalezlrjesus
// @contact.email gonzalezlrjesus@gmail.com

// @host localhost:3000
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", welcome)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition"
	))
	http.ListenAndServe(":3000", r)
}

// GET welcome Message
// @tags Message
// @Summary show message "welcome"
// @Description show message "welcome"
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /  [get]
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
