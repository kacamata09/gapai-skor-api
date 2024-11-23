package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"

	"golang.org/x/crypto/bcrypt"
	// "time"
	// "github.com/labstack/echo"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
	DB       *sql.DB
}

func CreateUserUseCase(repo domain.UserRepository) domain.UserUsecase {
	usecase := UserUsecase{
		UserRepo: repo,
	}

	return &usecase
}

func (uc UserUsecase) GetAllData() ([]domain.User, error) {
	data, err := uc.UserRepo.GetAll()
	return data, err
}

func (uc UserUsecase) GetByID(id string) (domain.User, error) {
	data, err := uc.UserRepo.GetByID(id)
	return data, err
}

func (uc UserUsecase) Create(input *domain.User) error {
	// usernameExisted, _ := uc.UserRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.UserRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hashedPassword)

	err = uc.UserRepo.Create(input)
	return err
}
