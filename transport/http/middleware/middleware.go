package middleware

import "github.com/labstack/echo"

type Middleware struct {
}



func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Allow-Origin", "*") // Mengizinkan semua origin
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true") // Jika perlu kredensial

		// Tangani preflight OPTIONS request
		if c.Request().Method == "OPTIONS" {
			return c.NoContent(204) // Mengembalikan status 204 untuk preflight
		}

		return next(c)
	}
}

func InitMiddleware() *Middleware {
	return &Middleware{}
}