package main

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"jjgdevelopment.com/go/rest-ws/handlers"
	"jjgdevelopment.com/go/rest-ws/middleware"
	"jjgdevelopment.com/go/rest-ws/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}

	server.Start(BindRoutes)
}

func BindRoutes(server server.Server, router *mux.Router) {
	router.Use(middleware.CheckAuthMiddleware(server))

	router.HandleFunc("/", handlers.HomeHandler(server)).Methods("GET")
	router.HandleFunc("/signup", handlers.SignUpHandler(server)).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler(server)).Methods("POST")
	router.HandleFunc("/me", handlers.MeHandler(server)).Methods("GET")
	router.HandleFunc("/posts", handlers.InsertPostHandler(server)).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(server)).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(server)).Methods("PUT")
	router.HandleFunc("/posts/{id}", handlers.DeletePostHandler(server)).Methods("DELETE")
	router.HandleFunc("/posts", handlers.ListPostHandler(server)).Methods("GET")
}
