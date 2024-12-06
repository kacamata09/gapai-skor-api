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

func (uc QuestionUsecase) GetByTestID(id string) ([]domain.Question, error) {
	data, err := uc.QuestionRepo.GetByTestID(id)
	return data, err
}

func (uc QuestionUsecase) Create(input *domain.Question) error {

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

func (uc QuestionUsecase) UpdateWithAnswerOptions(id string, input *domain.Question) error {

	fmt.Println("ah shit")
	tx, err := uc.Transaction.BeginTransaction()
	if err != nil {
		return err
	}

	if err = uc.QuestionRepo.Update(id, tx, input); err != nil {
		uc.Transaction.RollbackTransaction(tx)
		fmt.Println("gagal update question")
		return err
	}

	for _, option := range input.AnswerOptions {
		option.QuestionID = id
		if err := uc.AnswerOptionRepo.Update(option.ID, tx, &option); err != nil {
			uc.Transaction.RollbackTransaction(tx)
			fmt.Println("gagal update option")
			return err
		}
	}

	uc.Transaction.CommitTransaction(tx)

	return err
}

func (uc QuestionUsecase) Delete(id string) (error) {
	err := uc.QuestionRepo.Delete(id)
	return err
}