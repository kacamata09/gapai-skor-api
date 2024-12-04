package domain

import (
	"database/sql"
	// "time"
	// "github.com/labstack/echo"
)

type Question struct {
	ID              string         `json:"id"`
	TestID          string         `json:"test_id"`
	ContentQuestion string         `json:"content_question"`
	ImageURL        string         `json:"image_url"`
	AudioURL        string         `json:"audio_url"`
	QuestionType    string         `json:"question_type"`
	QuestionNumber  int16          `json:"question_number"`
	Points          int16          `json:"points"`
	AnswerOptions   []AnswerOption `json:"answer_options"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}

type QuestionWithOptions struct {
	ID              string   `json:"id"`
	ContentQuestion string   `json:"text"`
	ImageURL        string   `json:"image"`
	AudioURL        string   `json:"audio"`
	QuestionNumber  int16    `json:"question_number"`
	SelectedAnswer  string   `json:"selected_answer"`
	AnswerOptions   []string `json:"options"`
	AnswerOptionsID []string `json:"answer_id"`
	PlayCount       int8     `json:"play_count"`
}

type QuestionSession struct {
	ID          int16                 `json:"id"`
	SessionType string                `json:"session_type"`
	Questions   []QuestionWithOptions `json:"questions"`
}

type QuestionRepository interface {
	GetAll() ([]Question, error)
	GetByTestID(id string) ([]Question, error)
	GetByID(id string) (Question, error)
	Create(tx *sql.Tx, question *Question) (string, error)
	// Update(ar *Article) error
	// Delete(id string) error
}

type QuestionUsecase interface {
	GetAllData() ([]Question, error)
	GetByID(id string) (Question, error)
	GetByTestID(id string) ([]Question, error)
	Create(question *Question) error
	CreateWithAnswerOptions(question *Question) error
}
