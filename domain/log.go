package domain

import (
	"time"
)

type Logger struct {
	Uri        string    `json:"uri"`
	Created_at time.Time `json:"created_at"`
}

type LoggerRepository interface {
	Store(url string) (bool, error)
}
