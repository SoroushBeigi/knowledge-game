package userhandler

import "github.com/labstack/echo/v5"

func (h Handler) SetRoutes(e *echo.Echo) {
	uGroup := e.Group("/users")

	uGroup.POST("/register", h.userRegister)
	uGroup.POST("/login", h.userLogin)
	uGroup.GET("/profile", h.userProfile)
}
