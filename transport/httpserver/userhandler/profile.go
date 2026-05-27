package userhandler

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/constants"
	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/labstack/echo/v5"
)

func getClaims(c *echo.Context) *authservice.Claims {
	claims := c.Get(constants.AuthMiddlewareContextKey)

	return claims.(*authservice.Claims)
}

func (h Handler) userProfile(c *echo.Context) error {
	claims := getClaims(c)

	resp, err := h.userSvc.GetProfile(dto.GetProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
