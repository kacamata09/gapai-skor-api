package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type AttemptUsecase struct {
	AttemptRepo domain.AttemptRepository
	DB       *sql.DB
}

func CreateAttemptUseCase(repo domain.AttemptRepository) domain.AttemptUsecase {
	usecase := AttemptUsecase{
		AttemptRepo: repo,
	}

	return &usecase
}

func (uc AttemptUsecase) GetAllData() ([]domain.Attempt, error) {
	data, err := uc.AttemptRepo.GetAll()
	return data, err
}

func (uc AttemptUsecase) GetByID(id string) (domain.Attempt, error) {
	data, err := uc.AttemptRepo.GetByID(id)
	return data, err
}

func (uc AttemptUsecase) Create(input *domain.Attempt) error {
	// usernameExisted, _ := uc.AttemptRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.AttemptRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.AttemptRepo.Create(input)
	return err
}
