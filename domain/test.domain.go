package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)



type Test struct {
	ID string `json:"id"`
	TestCode string `json:"test_code"`
	TestTitle string `json:"test_title"`
	Description string `json:"description"`
	CreatedBy string `json:"created_by"`
	Duration int16 `json:"duration"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TestRepository interface {
	GetAll() ([]Test, error)
	GetByID(id string) (Test, error)
	GetByTestCode(testCode string) (Test, error)
	Create(test *Test) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type TestUsecase interface {
	GetAllData() ([]Test, error)
	GetByID(id string) (Test, error)
	Create(test *Test) error
}