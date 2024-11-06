package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func UserRoute(e *echo.Echo, uc domain.UserUsecase) {
	handler := UserHandler {
		usecase: uc,
	}
	e.GET("/user/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/user")
	})
	
	e.GET("/user", handler.GetAllHandler)
	e.POST("/user", handler.Create)
	e.GET("/user/:id", handler.GetByIDHandler)
}     

func (h *UserHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all user")

	return resp
}

func (h *UserHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil{
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *UserHandler) Create(c echo.Context) error {
	var data domain.User

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data) 
	
	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create user")
}
