package httpserver

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/labstack/echo/v5"
)

func (s Server) userRegister(c *echo.Context) error {
	var uReq userservice.RegisterRequest
	err := c.Bind(&uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	user, err := s.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
