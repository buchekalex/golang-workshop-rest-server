// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"go-mongo-crud-rest-api/internal/config"
	"go-mongo-crud-rest-api/internal/usecase"
	"go-mongo-crud-rest-api/internal/usecase/repository"
	"go-mongo-crud-rest-api/pkg/httpserver"
	"go-mongo-crud-rest-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	v1 "go-mongo-crud-rest-api/internal/controller/http/v1"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// create a database connection
	client, err := mongo.Connect(context.TODO(), &options.ClientOptions{Hosts: []string{cfg.Mongo.URI}})
	if err != nil {
		log.Fatal(err)
	}

	// create a repository
	repo := repository.NewRepository(client.Database("users"))
	elasticRepository, err := NewElasticsearchRepository([]string{cfg.Elastic.Address})
	if err != nil {
		log.Fatal(err)
	}

	// HTTP Server
	userUseCase := usecase.NewUserUseCase(repo, elasticRepository)
	server := v1.NewServer(userUseCase)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/users/:email", server.GetUser)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)
		router.GET("/health", server.HealthCheck)
		router.GET("/users/find/:email", server.FindUser)
	}

	httpServer := httpserver.New(router.Handler(), httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

// InitializeElasticsearchClient initializes and returns an Elasticsearch client.
func InitializeElasticsearchClient(addresses []string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
		Transport: &nethttp.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DisableCompression:    true,
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Elasticsearch client: %v", err)
	}

	// Check if the client is up and running
	res, err := esClient.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to ping Elasticsearch: %v", err)
	}
	defer func() {
		err = res.Body.Close()
	}()

	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch returned an error response: %s", res.String())
	}

	log.Printf("Elasticsearch cluster info: %s", res.String())

	return esClient, nil
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository with its dependencies.
func NewElasticsearchRepository(addresses []string) (usecase.SearchIndex, error) {
	esClient, err := InitializeElasticsearchClient(addresses)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Elasticsearch client: %v", err)
	}

	index := "users"

	return repository.NewElasticsearchRepository(esClient, index), nil
}
