package httpserver

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/httpmessage"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/labstack/echo/v5"
)

func (s Server) userRegister(c *echo.Context) error {
	var req dto.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if fieldErrs, err := s.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return c.JSON(code, map[string]any{
			"message":     msg,
			"fieldErrors": fieldErrs,
		})
	}

	user, err := s.userSvc.Register(req)
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, user)
}

func (s Server) userLogin(c *echo.Context) error {
	var req userservice.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.userSvc.Login(req)
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) userProfile(c *echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authSvc.ParseToken(authToken)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.GetProfile(userservice.GetProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmessage.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
