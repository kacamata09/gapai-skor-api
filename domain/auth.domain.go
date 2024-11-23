package domain

// "database/sql"
// "time"
// "github.com/labstack/echo"

type LoginResponse struct {
	Token    string `json:"id"`
	DataUser User   `json:"data_user"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUsecase interface {
	Login(login *UserLogin) (LoginResponse, error)
}
