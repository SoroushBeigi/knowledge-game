package httpserver

import (
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/userhandler"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, uv uservalidator.Validator) *Server {

	return &Server{
		config: config,
		userHandler: *userhandler.New(authSvc,
			userSvc,
			uv,
			config.Auth,
		),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	s.userHandler.SetRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
