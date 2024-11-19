package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type QuestionUsecase struct {
	QuestionRepo domain.QuestionRepository
	DB       *sql.DB
}

func CreateQuestionUseCase(repo domain.QuestionRepository) domain.QuestionUsecase {
	usecase := QuestionUsecase{
		QuestionRepo: repo,
	}

	return &usecase
}

func (uc QuestionUsecase) GetAllData() ([]domain.Question, error) {
	data, err := uc.QuestionRepo.GetAll()
	return data, err
}

func (uc QuestionUsecase) GetByID(id string) (domain.Question, error) {
	data, err := uc.QuestionRepo.GetByID(id)
	return data, err
}

func (uc QuestionUsecase) Create(input *domain.Question) error {
	// usernameExisted, _ := uc.QuestionRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.QuestionRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.QuestionRepo.Create(input)
	return err
}
