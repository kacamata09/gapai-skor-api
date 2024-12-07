package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type TestHandler struct {
	usecase domain.TestUsecase
}

func TestRoute(e *echo.Echo, uc domain.TestUsecase) {
	handler := TestHandler{
		usecase: uc,
	}
	e.GET("/test/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/test")
	})

	e.GET("/test", handler.GetAllHandler)
	e.POST("/test", handler.Create)
	e.PUT("/test/:id", handler.UpdateHandler)
	e.DELETE("/test/:id", handler.DeleteHandler)
	e.GET("/test/:id", handler.GetByIDHandler)
	e.GET("/test/code/:test_code", handler.GetByTestCodeWithQuestionsHandler)
}

func (h *TestHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all test")

	return resp
}

func (h *TestHandler) GetByIDHandler(c echo.Context) error {
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

func (h *TestHandler) GetByTestCodeWithQuestionsHandler(c echo.Context) error {
	testCode := c.Param("test_code")
	// testCode = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(testCode)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByTestCodeWithQuestions(testCode)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by test_code")
	return resp
}

func (h *TestHandler) Create(c echo.Context) error {
	var data domain.Test

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create test")
}

func (h *TestHandler) UpdateHandler(c echo.Context) error {
	id := c.Param("id")

	var data domain.Test

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Update(id, &data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success update")
	return resp
}

func (h *TestHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success update")
	return resp
}
