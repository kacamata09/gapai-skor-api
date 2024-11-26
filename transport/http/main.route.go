package httpRoutes

import (
	"database/sql"
	repositoryMySql "gapai-skor-api/repository/mysql"
	handler "gapai-skor-api/transport/http/handlers"
	"gapai-skor-api/transport/http/middleware"
	"gapai-skor-api/usecase"
	"net/http"
	"gapai-skor-api/repository/mysql/helper"
	"github.com/labstack/echo"
)

type Home struct {
	Message string `json:"message"`
}

func homeHandler(c echo.Context) error {
	data := Home{
		Message: "welcome to gapai-skor project",
	}
	return c.JSON(http.StatusOK, data)
}

func StartHttp(e *echo.Echo, db *sql.DB) {
	// init middleware
	middleware := middleware.InitMiddleware()
	e.Use(middleware.CORS)

	// assign home
	e.GET("/", homeHandler)

	// user
	userRepo := repositoryMySql.CreateRepoUser(db)
	userUseCase := usecase.CreateUserUseCase(userRepo)
	handler.UserRoute(e, userUseCase)

	// test
	testRepo := repositoryMySql.CreateRepoTest(db)
	testUseCase := usecase.CreateTestUseCase(testRepo)
	handler.TestRoute(e, testUseCase)

	// user duration test
	userTestDurationRepo := repositoryMySql.CreateRepoUserTestDuration(db)
	userTestDurationUseCase := usecase.CreateUserTestDurationUseCase(userTestDurationRepo)
	handler.UserTestDurationRoute(e, userTestDurationUseCase)

	// answer option
	answerOptionRepo := repositoryMySql.CreateRepoAnswerOption(db)
	answerOptionUseCase := usecase.CreateAnswerOptionUseCase(answerOptionRepo)
	handler.AnswerOptionRoute(e, answerOptionUseCase)


	// create transaction
	transaction := helper.CreateTransaction(db)

	// question
	questionRepo := repositoryMySql.CreateRepoQuestion(db)
	// questionUseCase := usecase.CreateQuestionUseCase(questionRepo)
	questionUsecase := usecase.CreateQuestionUseCase(usecase.QuestionUsecase{
		QuestionRepo: questionRepo,
		AnswerOptionRepo: answerOptionRepo,
		Transaction: transaction,
	})
	handler.QuestionRoute(e, questionUsecase)


	// attempt
	attemptRepo := repositoryMySql.CreateRepoAttempt(db)
	attemptUseCase := usecase.CreateAttemptUseCase(attemptRepo)
	handler.AttemptRoute(e, attemptUseCase)

	// attempt answer
	attemptAnswerRepo := repositoryMySql.CreateRepoAttemptAnswer(db)
	attemptAnswerUseCase := usecase.CreateAttemptAnswerUseCase(attemptAnswerRepo)
	handler.AttemptAnswerRoute(e, attemptAnswerUseCase)

	// auth
	authUseCase := usecase.CreateAuthUseCase(userRepo)
	handler.AuthRoute(e, authUseCase)

}
