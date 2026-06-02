package middleware

import (
	"github.com/SoroushBeigi/knowledge-game/pkg/constants"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	ejwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Auth(service authnservice.Service, config authnservice.Config) echo.MiddlewareFunc {
	return ejwt.WithConfig(ejwt.Config{
		ContextKey:    constants.AuthMiddlewareContextKey,
		SigningKey:    config.SignKey,
		SigningMethod: "HS256",
		ParseTokenFunc: func(c *echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {

				return nil, err
			}

			return claims, nil
		},
	})
}
