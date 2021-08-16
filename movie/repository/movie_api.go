package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/wongpinter/movie-metadata/domain"
)

var (
	baseURL = "https://omdbapi.com"
	apikey  = "faf7e5bb"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MovieApiRepository struct {
	client  HTTPClient
	timeout time.Duration
}

func MovieApi(client HTTPClient, timeout time.Duration) domain.MovieRepository {
	return &MovieApiRepository{
		client:  client,
		timeout: timeout,
	}
}

func (api *MovieApiRepository) Search(ctx context.Context, query string, page int) (*domain.SearchResult, error) {
	uri := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", baseURL, apikey, query, page)

	ctx, cancel := context.WithTimeout(ctx, api.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		log.Fatalf("Cannot create request: %v\n", err)
	}

	response, err := api.client.Do(request)

	if err != nil {
		log.Fatalf("Cannot connect to API Server: %v\n", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(http.StatusText(response.StatusCode))
	}

	var result *domain.SearchResult

	return result, json.NewDecoder(response.Body).Decode(&result)
}

func (api *MovieApiRepository) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	uri := fmt.Sprintf("%s/?apikey=%s&i=%s", baseURL, apikey, id)

	ctx, cancel := context.WithTimeout(ctx, api.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		log.Fatalf("Cannot create request: %v\n", err)
	}

	response, err := api.client.Do(request)

	if err != nil {
		log.Fatalf("Cannot connect to API Server: %v\n", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(http.StatusText(response.StatusCode))
	}

	var result *domain.Movie

	return result, json.NewDecoder(response.Body).Decode(&result)
}

func (api *MovieApiRepository) GetByTitle(ctx context.Context, name string) (*domain.Movie, error) {

	cleanName := url.QueryEscape(name)

	uri := fmt.Sprintf("%s/?apikey=%s&t=%s", baseURL, apikey, cleanName)

	log.Printf("Url: %v", uri)

	ctx, cancel := context.WithTimeout(ctx, api.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		log.Fatalf("Cannot create request: %v\n", err)
	}

	response, err := api.client.Do(request)

	if err != nil {
		log.Fatalf("Cannot connect to API Server: %v\n", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(http.StatusText(response.StatusCode))
	}

	var result *domain.Movie

	return result, json.NewDecoder(response.Body).Decode(&result)
}
