package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"jjgdevelopment.com/go/rest-ws/database"
	"jjgdevelopment.com/go/rest-ws/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (broker *Broker) Config() *Config {
	return broker.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required.")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required.")
	}

	if config.DatabaseURL == "" {
		return nil, errors.New("A database URL is required.")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (broker *Broker) Start(binder func(server Server, router *mux.Router)) {
	broker.router = mux.NewRouter()
	binder(broker, broker.router)

	repo, err := database.NewPostgresRepository(broker.config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	log.Println("Starting server on port", broker.Config().Port)

	if err := http.ListenAndServe(broker.config.Port, broker.router); err != nil {
		log.Fatal("[ERROR] Error counted on ", err)
	}
}
