package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type QuestionHandler struct {
	usecase domain.QuestionUsecase
}

func QuestionRoute(e *echo.Echo, uc domain.QuestionUsecase) {
	handler := QuestionHandler{
		usecase: uc,
	}
	e.GET("/question/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/question")
	})

	e.GET("/question", handler.GetAllHandler)
	e.POST("/question", handler.Create)
	e.GET("/question/:id", handler.GetByIDHandler)
}

func (h *QuestionHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all question")

	return resp
}

func (h *QuestionHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *QuestionHandler) Create(c echo.Context) error {
	var data domain.Question

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create question")
}
