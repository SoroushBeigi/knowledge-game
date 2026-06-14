package userhandler

import (
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/middleware"
	"github.com/labstack/echo/v5"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	uGroup := e.Group("/users")

	uGroup.POST("/register", h.userRegister)
	uGroup.POST("/login", h.userLogin)
	uGroup.GET("/profile", h.userProfile,
		middleware.Auth(h.authSvc, h.authConfig),
		middleware.UpsertPresence(h.presenceSvc),
	)
}
