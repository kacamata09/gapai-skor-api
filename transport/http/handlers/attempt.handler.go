package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type AttemptHandler struct {
	usecase domain.AttemptUsecase
}

func AttemptRoute(e *echo.Echo, uc domain.AttemptUsecase) {
	handler := AttemptHandler{
		usecase: uc,
	}
	e.GET("/attempt/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/attempt")
	})

	e.GET("/attempt", handler.GetAllHandler)
	e.POST("/attempt", handler.Create)
	e.GET("/attempt/:id", handler.GetByIDHandler)
	e.GET("/attempt/score/:id", handler.GetScore)
	e.GET("/attempt/history/:id", handler.GetHistory)
	e.GET("/attempt/user/:id", handler.GetAttemptTestUser)

}

func (h *AttemptHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all attempt")

	return resp
}

func (h *AttemptHandler) GetByIDHandler(c echo.Context) error {
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

func (h *AttemptHandler) GetHistory(c echo.Context) error {
	id := c.Param("id")

	data, err := h.usecase.GetAttemptHistory(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *AttemptHandler) GetAttemptTestUser(c echo.Context) error {
	id := c.Param("id")

	data, err := h.usecase.GetAttemptTestUser(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *AttemptHandler) GetScore(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetAttemptWithAttemptAnswer(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *AttemptHandler) Create(c echo.Context) error {
	var data domain.Attempt

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	id, err := h.usecase.Create(&data)

	data.ID = id

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create attempt")
}
