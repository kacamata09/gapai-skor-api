package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)



type Question struct {
	ID string `json:"id"`
	TestID string `json:"test_id"`
	ContentQuestion string `json:"content_question"`
	ImageURL string `json:"image_url"`
	AudioURL string `json:"audio_url"`
	QuestionType string `json:"question_type"`
	Points int16 `json:"points"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type QuestionRepository interface {
	GetAll() ([]Question, error)
	GetByID(id string) (Question, error)
	Create(question *Question) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type QuestionUsecase interface {
	GetAllData() ([]Question, error)
	GetByID(id string) (Question, error)
	Create(question *Question) error
}