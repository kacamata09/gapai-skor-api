package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type AttemptAnswerHandler struct {
	usecase domain.AttemptAnswerUsecase
}

func AttemptAnswerRoute(e *echo.Echo, uc domain.AttemptAnswerUsecase) {
	handler := AttemptAnswerHandler{
		usecase: uc,
	}
	e.GET("/attempt/answer/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/attempt/answer")
	})

	e.GET("/attempt/answer", handler.GetAllHandler)
	e.POST("/attempt/answer", handler.Create)
	e.GET("/attempt/answer/:id", handler.GetByIDHandler)
}

func (h *AttemptAnswerHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all attempt_answer")

	return resp
}

func (h *AttemptAnswerHandler) GetByIDHandler(c echo.Context) error {
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

func (h *AttemptAnswerHandler) Create(c echo.Context) error {
	var data domain.AttemptAnswer

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create attempt_answer")
}
