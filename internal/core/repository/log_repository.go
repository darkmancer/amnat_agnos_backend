package repository

import (
	"database/sql"
	"log"
	"strong_password_recommendation/internal/core/port"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) port.LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) LogRequestResponse(request string, response int) error {
	_, err := r.db.Exec("INSERT INTO logs (request, response) VALUES ($1, $2)", request, response)
	if err != nil {
		log.Printf("Error logging request/response: %v", err)
	}

	return err
}
