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
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type AttemptAnswerRepository interface {
	GetAll() ([]AttemptAnswer, error)
	GetByID(id string) (AttemptAnswer, error)
	Create(tx *sql.Tx, attempt_answer *AttemptAnswer) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type AttemptAnswerUsecase interface {
	GetAllData() ([]AttemptAnswer, error)
	GetByID(id string) (AttemptAnswer, error)
	Create(attempt_answer *AttemptAnswer) error
}
