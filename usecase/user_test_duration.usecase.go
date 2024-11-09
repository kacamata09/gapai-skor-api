package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type UserTestDurationUsecase struct {
	UserTestDurationRepo domain.UserTestDurationRepository
	DB                   *sql.DB
}

func CreateUserTestDurationUseCase(repo domain.UserTestDurationRepository) domain.UserTestDurationUsecase {
	usecase := UserTestDurationUsecase{
		UserTestDurationRepo: repo,
	}

	return &usecase
}

func (uc UserTestDurationUsecase) GetAllData() ([]domain.UserTestDuration, error) {
	data, err := uc.UserTestDurationRepo.GetAll()
	return data, err
}

func (uc UserTestDurationUsecase) GetByID(id string) (domain.UserTestDuration, error) {
	data, err := uc.UserTestDurationRepo.GetByID(id)
	return data, err
}

func (uc UserTestDurationUsecase) Create(input *domain.UserTestDuration) error {
	// usernameExisted, _ := uc.UserTestDurationRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.UserTestDurationRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.UserTestDurationRepo.Create(input)
	return err
}
