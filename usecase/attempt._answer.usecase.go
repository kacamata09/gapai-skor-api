package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type AttemptAnswerUsecase struct {
	AttemptAnswerRepo domain.AttemptAnswerRepository
	DB       *sql.DB
}

func CreateAttemptAnswerUseCase(repo domain.AttemptAnswerRepository) domain.AttemptAnswerUsecase {
	usecase := AttemptAnswerUsecase{
		AttemptAnswerRepo: repo,
	}

	return &usecase
}

func (uc AttemptAnswerUsecase) GetAllData() ([]domain.AttemptAnswer, error) {
	data, err := uc.AttemptAnswerRepo.GetAll()
	return data, err
}

func (uc AttemptAnswerUsecase) GetByID(id string) (domain.AttemptAnswer, error) {
	data, err := uc.AttemptAnswerRepo.GetByID(id)
	return data, err
}

func (uc AttemptAnswerUsecase) Create(input *domain.AttemptAnswer) error {
	// usernameExisted, _ := uc.AttemptAnswerRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.AttemptAnswerRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.AttemptAnswerRepo.Create(input)
	return err
}
