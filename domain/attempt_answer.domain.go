package domain

import (
	"database/sql"
	// "time"
	// "github.com/labstack/echo"
)

type AttemptAnswer struct {
	ID                     string `json:"id"`
	AttemptID              string `json:"attempt_id"`
	QuestionID             string `json:"question_id"`
	SelectedAnswerOptionID string `json:"selected_answer_option_id"`
	IsCorrect              int8   `json:"is_correct"`
	QuestionType           string `json:"question_type"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type AttemptAnswerRepository interface {
	GetAll() ([]AttemptAnswer, error)
	GetByID(id string) (AttemptAnswer, error)
	Create(tx *sql.Tx, attempt_answer *AttemptAnswer) error
	Update(tx *sql.Tx, attempt_answer *AttemptAnswer) error
	VerifAttemptAnswerIsThere(attempt_answer *AttemptAnswer) (string, error)
	// Update(ar *Article) error
	// Delete(id string) error
}

type AttemptAnswerUsecase interface {
	GetAllData() ([]AttemptAnswer, error)
	GetByID(id string) (AttemptAnswer, error)
	Create(attempt_answer *AttemptAnswer) error
}
