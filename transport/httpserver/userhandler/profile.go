package userhandler

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/claims"
	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/labstack/echo/v5"
)

func (h Handler) userProfile(c *echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c)

	resp, err := h.userSvc.GetProfile(c.Request().Context(), dto.GetProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
