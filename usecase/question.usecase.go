package usecase

import (
	// "database/sql"
	"gapai-skor-api/domain"
	"gapai-skor-api/repository/mysql/helper"

	// "time"
	"fmt"
	// "github.com/labstack/echo"
)

type QuestionUsecase struct {
	QuestionRepo domain.QuestionRepository
	// DB           *sql.DB
	AnswerOptionRepo domain.AnswerOptionRepository
	Transaction      helper.TransactionFunc
}

// func CreateQuestionUseCase(repo domain.QuestionRepository) domain.QuestionUsecase {
// 	usecase := QuestionUsecase{
// 		QuestionRepo: repo,
// 	}

//		return &usecase
//	}
func CreateQuestionUseCase(repo QuestionUsecase) domain.QuestionUsecase {

	return &repo
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

	_, err := uc.QuestionRepo.Create(nil, input)
	return err
}

func (uc QuestionUsecase) CreateWithAnswerOptions(input *domain.Question) error {

	fmt.Println("ah shit")
	tx, err := uc.Transaction.BeginTransaction()
	if err != nil {
		return err
	}

	var questionID string

	if questionID, err = uc.QuestionRepo.Create(tx, input); err != nil {
		uc.Transaction.RollbackTransaction(tx)
		fmt.Println("gagal create question")
		return err
	}

	for _, option := range input.AnswerOptions {
		option.QuestionID = questionID
		if err := uc.AnswerOptionRepo.Create(tx, &option); err != nil {
			uc.Transaction.RollbackTransaction(tx)
			fmt.Println("gagal create option")
			return err
		}
	}

	uc.Transaction.CommitTransaction(tx)

	return err
}
