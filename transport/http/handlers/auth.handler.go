package handler

import (
	// "fmt"

	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func AuthRoute(e *echo.Echo, uc domain.AuthUsecase) {
	handler := AuthHandler{
		usecase: uc,
	}
	e.GET("/auth/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/auth")
	})

	e.POST("/auth/login", handler.Login)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var data domain.UserLogin

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	token, err := h.usecase.Login(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, token, "success login")
}
