package main

import (
	"encoding/json"
	"net/http"

	_ "gonzalezlrjesus/swag-golang/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Greetings struct
type Greetings struct {
	Message string `json:"message"`
}

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

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/", ShowMessage)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition"
	))
	http.ListenAndServe(":3000", r)
}

// ShowMessage POST
// @tags Show message
// @Summary Show message
// @Description Show message
// @Accept  json
// @Produce  json
// @Param data body Greetings true "Show Body message"
// @Success 200 {object} string
// @Router /  [post]
func ShowMessage(w http.ResponseWriter, r *http.Request) {
	data := &Greetings{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	json.NewEncoder(w).Encode(data.Message)
}
