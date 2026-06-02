package claims

import (
	"github.com/SoroushBeigi/knowledge-game/pkg/constants"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/labstack/echo/v5"
)

func GetClaimsFromEchoContext(c *echo.Context) *authnservice.Claims {
	claims := c.Get(constants.AuthMiddlewareContextKey)
	return claims.(*authnservice.Claims)
}
