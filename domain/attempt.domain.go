package domain

import (
	"database/sql"
	// "time"
	// "github.com/labstack/echo"
)

type Attempt struct {
	ID             string          `json:"id"`
	UserID         string          `json:"user_id"`
	TestID         string          `json:"test_id"`
	Score          int16           `json:"score"`
	AttemptDate    string          `json:"attempt_date"`
	AttemptAnswers []AttemptAnswer `json:"attempt_answers"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
}

type AttemptRepository interface {
	GetAll() ([]Attempt, error)
	GetByID(id string) (Attempt, error)
	Create(tx *sql.Tx, attempt *Attempt) (string, error)
	Update(attempt *Attempt) error
	VerifAttemptIsThere(attempt *Attempt) (string, error)
	// Delete(id string) error
}

type AttemptUsecase interface {
	GetAllData() ([]Attempt, error)
	GetByID(id string) (Attempt, error)
	Create(attempt *Attempt) error
	// CreateWithAttemptAnswers(attempt *Attempt) error
}
