package adminhandler

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/labstack/echo/v5"
)

func (h Handler) listUsers(c *echo.Context) error {
	list, err := h.adminSvc.ListAllUsers()
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"users": list,
	})
}
