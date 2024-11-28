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

func (uc TestUsecase) GetByTestCodeWithQuestions(testCode string) (domain.Test, error) {
	data, err := uc.TestRepo.GetByTestCodeWithQuestions(testCode)
	return data, err
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
