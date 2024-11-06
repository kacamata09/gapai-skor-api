package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)



type UserTestDuration struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	TestID string `json:"test_id"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Duration int16 `json:"duration"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserTestDurationRepository interface {
	GetAll() ([]UserTestDuration, error)
	GetByID(id string) (UserTestDuration, error)
	Create(userTestDuration *UserTestDuration) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type UserTestDurationUsecase interface {
	GetAllData() ([]UserTestDuration, error)
	GetByID(id string) (UserTestDuration, error)
	Create(userTestDuration *UserTestDuration) error
}