package adminhandler

import (
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/middleware"
	"github.com/labstack/echo/v5"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	uGroup := e.Group("/admin/users")

	uGroup.GET("/", h.listUsers,
		middleware.Auth(h.authnSvc, h.authConfig),
		middleware.AccessCheck(h.authzSvc, entity.UserListPermission),
	)

}
