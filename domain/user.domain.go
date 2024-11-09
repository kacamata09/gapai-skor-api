package domain

// "database/sql"
// "time"
// "github.com/labstack/echo"

type User struct {
	ID        string `json:"id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	Create(user *User) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type UserUsecase interface {
	GetAllData() ([]User, error)
	GetByID(id string) (User, error)
	Create(user *User) error
}
