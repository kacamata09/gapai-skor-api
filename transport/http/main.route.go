package httpRoutes

import (
	"database/sql"
	repositoryMySql "gapai-skor-api/repository/mysql"
	handler "gapai-skor-api/transport/http/handlers"
	"gapai-skor-api/transport/http/middleware"
	"gapai-skor-api/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type Home struct {
	Message string `json:"message"`
}

func homeHandler(c echo.Context) error {
	data := Home {
		Message : "welcome to gapai-skor project",
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



}