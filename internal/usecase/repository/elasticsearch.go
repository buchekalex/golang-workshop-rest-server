package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"go-mongo-crud-rest-api/internal/entitiy"
	"net/http"
	"strings"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// SearchResults wraps the Elasticsearch search response.
type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

// Hit wraps the document returned in search response.
type Hit struct {
	URL        string        `json:"url"`
	Sort       []interface{} `json:"sort"`
	Highlights *struct {
		Title      []string `json:"title"`
		Alt        []string `json:"alt"`
		Transcript []string `json:"transcript"`
	} `json:"highlights,omitempty"`
}

// ElasticsearchRepository is the implementation of the Repository interface using Elasticsearch as the data store.
type ElasticsearchRepository struct {
	esClient *es8.Client
	index    string
}

// NewElasticsearchRepository creates a new instance of ElasticsearchRepository.
func NewElasticsearchRepository(esClient *es8.Client, index string) *ElasticsearchRepository {
	if err := createIndexIfNotExists(esClient, index); err != nil {
		log.Print(err)
	}

	return &ElasticsearchRepository{
		esClient: esClient,
		index:    index,
	}
}

func createIndexIfNotExists(esClient *es8.Client, index string) error {
	// First, check if the index exists.
	res, err := esClient.Indices.Exists([]string{index})
	if err != nil {
		return err
	}

	// If the index does not exist, create it.
	if res.StatusCode == http.StatusNotFound {
		req := esapi.IndicesCreateRequest{
			Index: index,
			Body: strings.NewReader(`
				{
					"mappings": {
						"properties": {
							"email": {
								"type": "keyword"
							}
							// Specify other fields as needed.
						}
					}
				}`),
		}

		res, err := req.Do(context.Background(), esClient)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return errors.New("Error creating index: " + res.String())
		}
	}

	return nil
}

func (er *ElasticsearchRepository) FindUser(ctx context.Context, email string) (*entitiy.User, error) {
	return nil, nil
}

type esSearchResponse struct {
	Hits struct {
		Hits []struct {
			Source entitiy.User `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// GetUser retrieves a user from Elasticsearch based on the given email.
func (er *ElasticsearchRepository) GetUser(ctx context.Context, email string) (*entitiy.User, error) {
	// TODO: implement me
	return nil, nil
}

// CreateUser creates a new user in Elasticsearch.
func (er *ElasticsearchRepository) CreateUser(ctx context.Context, user *entitiy.User) (*entitiy.User, error) {
	// TODO: implement me
	return user, nil
}

// UpdateUser updates an existing user in Elasticsearch.
func (er *ElasticsearchRepository) UpdateUser(ctx context.Context, user *entitiy.User) (*entitiy.User, error) {
	// TODO: implement me
	return user, nil
}

// DeleteUser deletes a user from Elasticsearch based on the given email.
func (er *ElasticsearchRepository) DeleteUser(ctx context.Context, email string) error {
	// TODO: implement me
	return nil
}
