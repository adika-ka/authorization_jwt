package main

import (
	"fmt"
	"log"
	"net/http"

	_ "authorization_jwt/docs"
	"authorization_jwt/internal/handlers"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Geo Service API
// @version 1.0
// @description	API для работы с адресами и геокодингом
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	var userStore = handlers.NewUserStore()

	r := chi.NewRouter()

	r.Route("/api/address", func(r chi.Router) {
		r.Use(handlers.JWTMiddleware)
		r.Post("/search", handlers.SearchHandler)
		r.Post("/geocode", handlers.GeocodeHandler)
	})

	r.Post("/api/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, userStore)
	})
	r.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userStore)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
