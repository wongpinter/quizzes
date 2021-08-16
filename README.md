# Stockbit Quiz Question 2

```
$ git clone https://github.com/wongpinter/stoctbit-quizz

$ cd stoctbit-quizz

$ go mod tidy

$ docker compose up -d // init mysql database

$ go run ./app

$ go run ./cmd/grpc-client // gRPC services

// From browser
// http://localhost:8090/v1/movie/search?query=super&page=1 # search service
// http://localhost:8090/v1/movie/title/batman # single record by title service
// http://localhost:8090/v1/movie/id/tt4116284 # single record by imdbid service
```