package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/wongpinter/movie-metadata/movie/delivery/grpc/proto/v1"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewMovieServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r1, err := c.GetSingleMovieByID(ctx, &pb.ByID{Id: "tt0112462"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Movie info: %s", r1.Title)

	r2, err := c.GetSingleMovieByTitle(ctx, &pb.ByTitle{Title: "The Dark Knight"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Movie info: %s", r2)

	resp, err := c.SearchMovies(ctx, &pb.SearchQuery{
		Query: "Batman", Page: 1,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Movie info: %s", resp)
}
