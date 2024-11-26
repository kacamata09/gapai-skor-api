package domain

import (
	"database/sql"
	// "time"
	// "github.com/labstack/echo"
)



type AnswerOption struct {
	ID string `json:"id"`
	QuestionID string `json:"question_id"`
	ContentAnswer string `json:"content_answer"`
	ImageURL string `json:"image_url"`
	AudioURL string `json:"audio_url"`
	IsCorrect int8 `json:"is_correct"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AnswerOptionRepository interface {
	GetAll() ([]AnswerOption, error)
	GetByID(id string) (AnswerOption, error)
	Create(tx *sql.Tx, answerOption *AnswerOption) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type AnswerOptionUsecase interface {
	GetAllData() ([]AnswerOption, error)
	GetByID(id string) (AnswerOption, error)
	Create(answerOption *AnswerOption) error
}