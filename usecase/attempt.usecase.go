package usecase

import (
	"database/sql"
	"fmt"
	"gapai-skor-api/domain"

	// "time"
	// "github.com/labstack/echo"
	"gapai-skor-api/repository/mysql/helper"
)

type AttemptUsecase struct {
	AttemptRepo        domain.AttemptRepository
	DB                 *sql.DB
	Transaction        helper.TransactionFunc
	AtttemptAnswerRepo domain.AttemptAnswerRepository
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

func (uc AttemptUsecase) GetAttemptHistory(id string) ([]domain.Attempt, error) {
	data, err := uc.AttemptRepo.GetAttemptHistory(id)
	return data, err
}

func (uc AttemptUsecase) GetAttemptTestUser(id string) ([]domain.Attempt, error) {
	data, err := uc.AttemptRepo.GetAttemptTestUser(id)
	return data, err
}

func (uc AttemptUsecase) GetAttemptWithAttemptAnswer(id string) (domain.Attempt, error) {

	data, err := uc.AttemptRepo.GetAttemptWithAttemptAnswer(id)

	var listening, structure, reading float64
	for _, answer := range data.AttemptAnswers {
		if answer.IsCorrect == 1 {
			switch answer.QuestionType {
			case "Listening":
				listening += 1.36
			case "Structure":
				structure += 1.7
			case "Reading":
				reading += 1.36
			}
		}
	}

	totalScore := (listening + structure + reading) * 10 / 3

	newData := domain.Attempt{
		ID:    id,
		Score: int16(totalScore),
	}

	err = uc.AttemptRepo.Update(&newData)

	data.Score = int16(totalScore)

	return data, err
}

func (uc AttemptUsecase) Create(input *domain.Attempt) (id string, err error) {

	lenAttempts, err := uc.AttemptRepo.VerifAttemptIsThere(input)
	if lenAttempts < 2 {
		id, err = uc.AttemptRepo.Create(nil, input)
	} else {
		// input.ID = id
		// err = uc.AttemptRepo.Update(input)
		err = fmt.Errorf("kode sudah kadaluarsa")
	}
	return id, err
}

func (uc AttemptUsecase) Delete(id string) error {
	err := uc.AttemptRepo.Delete(id)
	return err
}
