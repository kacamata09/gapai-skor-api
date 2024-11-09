package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type UserTestDurationHandler struct {
	usecase domain.UserTestDurationUsecase
}

func UserTestDurationRoute(e *echo.Echo, uc domain.UserTestDurationUsecase) {
	handler := UserTestDurationHandler{
		usecase: uc,
	}
	e.GET("/user/test_duration", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/user/test_duration")
	})

	e.GET("/user/test_duration", handler.GetAllHandler)
	e.POST("/user/test_duration", handler.Create)
	e.GET("/user/test_duration/:id", handler.GetByIDHandler)
}

func (h *UserTestDurationHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all user_test_duration")

	return resp
}

func (h *UserTestDurationHandler) GetByIDHandler(c echo.Context) error {
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

func (h *UserTestDurationHandler) Create(c echo.Context) error {
	var data domain.UserTestDuration

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create user_test_duration")
}
