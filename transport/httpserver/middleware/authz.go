package middleware

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/pkg/claims"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/labstack/echo/v5"
)

func AccessCheck(service authzservice.Service, permissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			claims := claims.GetClaimsFromEchoContext(c)
			allowed, err := service.HasAccess(claims.UserID, claims.Role, permissions...)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": errmessage.SomethingWentWrong,
				})
			}

			if !allowed {
				return c.JSON(http.StatusForbidden, map[string]any{
					"message": errmessage.ErrorMsgForbidden,
				})
			}

			return next(c)
		}
	}
}
