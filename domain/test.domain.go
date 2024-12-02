package domain

// "database/sql"
// "time"
// "github.com/labstack/echo"

type Test struct {
	ID          string     `json:"id"`
	TestCode    string     `json:"test_code"`
	TestTitle   string     `json:"test_title"`
	Description string     `json:"description"`
	CreatedBy   string     `json:"created_by"`
	Duration    int16      `json:"duration"`
	Questions   []Question `json:"questions"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type TestWithQuestion struct {
	ID          string            `json:"id"`
	TestCode    string            `json:"test_code"`
	TestTitle   string            `json:"test_title"`
	Description string            `json:"description"`
	CreatedBy   string            `json:"created_by"`
	Duration    int16             `json:"duration"`
	Sessions    []QuestionSession `json:"sessions"`
}

type TestRepository interface {
	GetAll() ([]Test, error)
	GetByID(id string) (Test, error)
	GetByTestCode(testCode string) (Test, error)
	GetByTestCodeWithQuestions(testCode string) (Test, error)
	Create(test *Test) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type TestUsecase interface {
	GetAllData() ([]Test, error)
	GetByID(id string) (Test, error)
	GetByTestCodeWithQuestions(id string) (TestWithQuestion, error)
	Create(test *Test) error
}
