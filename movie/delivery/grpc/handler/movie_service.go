package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/wongpinter/movie-metadata/movie/delivery/grpc/proto/v1"
	ma "github.com/wongpinter/movie-metadata/movie/repository"
)

type MoviesServer struct {
	pb.UnimplementedMovieServiceServer
}

func NewMovieServer() pb.MovieServiceServer {
	return new(MoviesServer)
}

var (
	api = ma.MovieApi(&http.Client{}, 30*time.Second)
)

func (s *MoviesServer) Ping(ctx context.Context, req *pb.Empty) (*pb.Pong, error) {
	log.Printf("Ping bro")

	return &pb.Pong{Pong: "PONG"}, nil
}

func (s *MoviesServer) GetSingleMovieByID(ctx context.Context, id *pb.ByID) (*pb.MovieResponse, error) {
	log.Printf("Requesting movie with id %s", id.Id)

	response, err := api.GetByID(context.Background(), id.Id)

	if err != nil {
		log.Printf("Cannot connect to API server: %v", err)
	}

	return &pb.MovieResponse{
		ImdbID:     response.ImdbID,
		Title:      response.Title,
		Year:       response.Year,
		Rated:      response.Rated,
		Released:   response.Released,
		Runtime:    response.Runtime,
		Genre:      response.Genre,
		Writer:     response.Writer,
		Actors:     response.Actors,
		Plot:       response.Plot,
		Language:   response.Language,
		Country:    response.Country,
		Awards:     response.Awards,
		Poster:     response.Poster,
		Metascore:  response.Metascore,
		ImdbRating: response.ImdbRating,
		ImdbVotes:  response.ImdbVotes,
		Type:       response.Type,
		Dvd:        response.DVD,
		BoxOffice:  response.BoxOffice,
		Production: response.Production,
		Website:    response.Website,
		Response:   response.Response,
	}, nil
}

func (s *MoviesServer) GetSingleMovieByTitle(ctx context.Context, title *pb.ByTitle) (*pb.MovieResponse, error) {
	log.Printf("Requesting movie with name %s", title.Title)

	response, err := api.GetByTitle(context.Background(), title.Title)

	if err != nil {
		log.Printf("Cannot connect to API server: %v", err)
	}

	return &pb.MovieResponse{
		ImdbID:     response.ImdbID,
		Title:      response.Title,
		Year:       response.Year,
		Rated:      response.Rated,
		Released:   response.Released,
		Runtime:    response.Runtime,
		Genre:      response.Genre,
		Writer:     response.Writer,
		Actors:     response.Actors,
		Plot:       response.Plot,
		Language:   response.Language,
		Country:    response.Country,
		Awards:     response.Awards,
		Poster:     response.Poster,
		Metascore:  response.Metascore,
		ImdbRating: response.ImdbRating,
		ImdbVotes:  response.ImdbVotes,
		Type:       response.Type,
		Dvd:        response.DVD,
		BoxOffice:  response.BoxOffice,
		Production: response.Production,
		Website:    response.Website,
		Response:   response.Response,
	}, nil
}

func (s *MoviesServer) SearchMovies(ctx context.Context, query *pb.SearchQuery) (*pb.MovieSearchResults, error) {
	if query.Page <= 0 {
		query.Page = 1
	}

	log.Printf("Requesting movie with query %s, page %d", query.Query, query.Page)

	responses, err := api.Search(context.Background(), query.Query, int(query.Page))

	if err != nil {
		log.Printf("Cannot connect to API server: %v", err)
	}

	var movies []*pb.Result

	for _, response := range responses.Search {
		movie := &pb.Result{
			Imdbid: response.ImdbID,
			Title:  response.Title,
			Year:   response.Year,
			Poster: response.Poster,
		}

		movies = append(movies, movie)
	}

	return &pb.MovieSearchResults{
		Result:       movies,
		TotalResults: responses.TotalResults,
		Response:     responses.Response,
	}, nil
}
