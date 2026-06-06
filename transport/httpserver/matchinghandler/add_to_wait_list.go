package matchinghandler

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/claims"
	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/labstack/echo/v5"
)

func (h Handler) addToWaitingList(c *echo.Context) error {
	var req dto.AddToWaitingListRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claims := claims.GetClaimsFromEchoContext(c)
	req.UserID = claims.UserID

	if fieldErrs, err := h.matchingValidator.ValidateAddToMatchingRequest(req); err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return c.JSON(code, map[string]any{
			"message":     msg,
			"fieldErrors": fieldErrs,
		})
	}

	resp, err := h.matchingSvc.AddToWaitingList(req)
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
