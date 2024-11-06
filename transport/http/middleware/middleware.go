package middleware

import "github.com/labstack/echo"

type Middleware struct {
}



func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Origin-Allow", "*")
		return next(c)
	}
}

func InitMiddleware() *Middleware {
	return &Middleware{}
}