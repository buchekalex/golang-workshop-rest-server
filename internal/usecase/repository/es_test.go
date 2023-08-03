package repository

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go-mongo-crud-rest-api/internal/usecase"
	nethttp "net/http"
	"testing"
	"time"
)

func Test_ElasticSearch(t *testing.T) {
	searchIndex, err := NewElasticsearchIndex([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal(err)
	}

	usr, err := searchIndex.FindUser(context.Background(), "ypWYMtxKFm")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(usr)
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

	return esClient, nil
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository with its dependencies.
func NewElasticsearchIndex(addresses []string) (usecase.SearchIndex, error) {
	esClient, err := InitializeElasticsearchClient(addresses)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Elasticsearch client: %v", err)
	}

	index := "users"

	return NewElasticsearchRepository(esClient, index), nil
}
