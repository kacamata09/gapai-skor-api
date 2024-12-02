package usecase

import (
	"database/sql"
	"gapai-skor-api/domain"
	// "time"
	// "github.com/labstack/echo"
)

type TestUsecase struct {
	TestRepo domain.TestRepository
	DB       *sql.DB
}

func CreateTestUseCase(repo domain.TestRepository) domain.TestUsecase {
	usecase := TestUsecase{
		TestRepo: repo,
	}

	return &usecase
}

func (uc TestUsecase) GetAllData() ([]domain.Test, error) {
	data, err := uc.TestRepo.GetAll()
	return data, err
}

func (uc TestUsecase) GetByID(id string) (domain.Test, error) {
	data, err := uc.TestRepo.GetByID(id)
	return data, err
}

func (uc TestUsecase) GetByTestCodeWithQuestions(testCode string) (domain.TestWithQuestion, error) {
	data, err := uc.TestRepo.GetByTestCodeWithQuestions(testCode)
	newData := domain.TestWithQuestion {
		ID : data.ID,
		TestCode : data.TestCode,
		TestTitle : data.TestTitle,
		Description : data.Description,
		CreatedBy : data.CreatedBy,
		Duration : data.Duration,
	} 

	
	sessions := [3]domain.QuestionSession{
		{ID : 1, SessionType : "Listening"},
		{ID : 2, SessionType : "Structure"},
		{ID : 3, SessionType : "Reading"},
	}
	for _, question := range(data.Questions) {
		newFormatQuestion := domain.QuestionWithOptions{
			ID
			ContentQuestion
			ImageURL
			AudioURL
			QuestionType
			QuestionNumber
			SelectedAnswer
			AnswerOptions
			PlayCount
		}
		if question.QuestionType == "Listening" {
			sessions[0].Questions = append(sessions[0].Questions, question)
		}
		if question.QuestionType == "Structure" {
			sessions[1].Questions = append(sessions[1].Questions, question)
		}
		if question.QuestionType == "Reading" {
			sessions[2].Questions = append(sessions[2].Questions, question)
		}
	}

	newData.Sessions = sessions
	return newData, err
}

func (uc TestUsecase) Create(input *domain.Test) error {
	// usernameExisted, _ := uc.TestRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.TestRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }

	err := uc.TestRepo.Create(input)
	return err
}
