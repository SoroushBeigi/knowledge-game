package matchinghandler

import (
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/middleware"
	"github.com/labstack/echo/v5"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	mGroup := e.Group("/matchmaking")

	mGroup.POST("/add-to-waiting-list",
		h.addToWaitingList,
		middleware.Auth(h.authSvc, h.authConfig),
		middleware.UpsertPresence(h.presenceService),
	)

}
