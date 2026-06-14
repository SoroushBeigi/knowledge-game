package middleware

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/claims"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/timestamp"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"github.com/labstack/echo/v5"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			claims := claims.GetClaimsFromEchoContext(c)
			_, err := service.UpsertPresence(
				c.Request().Context(),
				dto.UpsertPresenceRequest{UserID: claims.UserID, Timestamp: timestamp.Now()},
			)

			if err != nil {
				//TODO: consider logging only, without showing error
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": errmessage.SomethingWentWrong,
				})
			}

			return next(c)
		}
	}
}
