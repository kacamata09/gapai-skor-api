package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"

	"golang.org/x/crypto/bcrypt"

	// "github.com/labstack/echo"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthUsecase struct {
	UserRepo domain.UserRepository
	DB       *sql.DB
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("your_secret_key")

func CreateAuthUseCase(repo domain.UserRepository) domain.AuthUsecase {
	usecase := AuthUsecase{
		UserRepo: repo,
	}

	return &usecase
}

func (uc AuthUsecase) Login(input *domain.UserLogin) (domain.LoginResponse, error) {

	// perbandingan hashed dan tidak
	dataUser, err := uc.UserRepo.GetByUsername(input.Username)

	if err != nil {
		return domain.LoginResponse{}, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(input.Password))
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// Create JWT token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: dataUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	// usernameExisted, _ := uc.UserRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.UserRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }
	dataUser.Password = ""

	data := domain.LoginResponse{
		Token:    tokenString,
		DataUser: dataUser,
	}

	return data, err
}
