package userhandler

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/labstack/echo/v5"
)

func (h Handler) userRegister(c *echo.Context) error {
	var req dto.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrs, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return c.JSON(code, map[string]any{
			"message":     msg,
			"fieldErrors": fieldErrs,
		})
	}

	user, err := h.userSvc.Register(req)
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, user)
}
