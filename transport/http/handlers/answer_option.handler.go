package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type AnswerOptionHandler struct {
	usecase domain.AnswerOptionUsecase
}

func AnswerOptionRoute(e *echo.Echo, uc domain.AnswerOptionUsecase) {
	handler := AnswerOptionHandler{
		usecase: uc,
	}
	e.GET("/answer_option/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/answer_option")
	})

	e.GET("/answer_option", handler.GetAllHandler)
	e.POST("/answer_option", handler.Create)
	e.GET("/answer_option/:id", handler.GetByIDHandler)
}

func (h *AnswerOptionHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all answer_option")

	return resp
}

func (h *AnswerOptionHandler) GetByIDHandler(c echo.Context) error {
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

func (h *AnswerOptionHandler) Create(c echo.Context) error {
	var data domain.AnswerOption

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create answer_option")
}
