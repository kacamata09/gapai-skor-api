package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type TestUsecase struct {
	UserRepo domain.TestRepository
	DB       *sql.DB
}

func CreateTestUseCase(repo domain.TestRepository) domain.TestUsecase {
	usecase := TestUsecase{
		UserRepo: repo,
	}

	return &usecase
}

func (uc TestUsecase) GetAllData() ([]domain.Test, error) {
	data, err := uc.UserRepo.GetAll()
	return data, err
}

func (uc TestUsecase) GetByID(id string) (domain.Test, error) {
	data, err := uc.UserRepo.GetByID(id)
	return data, err
}

func (uc TestUsecase) Create(input *domain.Test) error {
	// usernameExisted, _ := uc.UserRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.UserRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.UserRepo.Create(input)
	return err
}
