package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gonzalezlrjesus/swag-golang/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/gonzalezlrjesus/swag-golang/docs"
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
	maxAge := 300
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Server port reading
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // localhost
	}

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/", handlers.ShowMessage)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("./swagger/doc.json"), // The url pointing to API definition"
	))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		os.Exit(1)

		return
	}

	log.Println("Ending main func")
}
