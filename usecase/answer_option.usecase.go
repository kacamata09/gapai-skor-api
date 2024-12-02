package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type AnswerOptionUsecase struct {
	AnswerOptionRepo domain.AnswerOptionRepository
	DB               *sql.DB
}

func CreateAnswerOptionUseCase(repo domain.AnswerOptionRepository) domain.AnswerOptionUsecase {
	usecase := AnswerOptionUsecase{
		AnswerOptionRepo: repo,
	}

	return &usecase
}

func (uc AnswerOptionUsecase) GetAllData() ([]domain.AnswerOption, error) {
	data, err := uc.AnswerOptionRepo.GetAll()
	return data, err
}

func (uc AnswerOptionUsecase) GetByID(id string) (domain.AnswerOption, error) {
	data, err := uc.AnswerOptionRepo.GetByID(id)
	return data, err
}

func (uc AnswerOptionUsecase) Create(input *domain.AnswerOption) error {
	// usernameExisted, _ := uc.AnswerOptionRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.AnswerOptionRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.AnswerOptionRepo.Create(nil, input)
	return err
}
