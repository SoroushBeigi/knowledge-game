package httpserver

import (
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/adminhandler"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/userhandler"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config       config.Config
	userHandler  userhandler.Handler
	adminHandler adminhandler.Handler
}

func New(config config.Config, authSvc authnservice.Service, userSvc userservice.Service,
	uv uservalidator.Validator, adminSvc adminservice.Service, authzService authzservice.Service,
) *Server {

	return &Server{
		config: config,
		userHandler: *userhandler.New(
			authSvc,
			userSvc,
			uv,
			config.Auth,
		),
		adminHandler: *adminhandler.New(
			authSvc,
			adminSvc,
			config.Auth,
			authzService,
		),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	s.userHandler.SetRoutes(e)
	s.adminHandler.SetRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
