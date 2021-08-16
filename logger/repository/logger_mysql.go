package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	domain "github.com/wongpinter/movie-metadata/domain"
)

type mysqlLoggerRepository struct {
	DB *sql.DB
}

func Logger(db *sql.DB) domain.LoggerRepository {
	return &mysqlLoggerRepository{
		DB: db,
	}
}

func (lo *mysqlLoggerRepository) Store(url string) (bool, error) {
	query := "insert into logs(url, created_at) values (?, ?)"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	_, err := lo.DB.Exec(query, url, time.Now().In(loc))

	if err != nil {
		return false, err
	}

	return true, nil
}
