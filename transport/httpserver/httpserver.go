package httpserver

import (
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/adminhandler"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/matchinghandler"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver/userhandler"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config          *config.Config
	userHandler     *userhandler.Handler
	adminHandler    *adminhandler.Handler
	matchingHandler *matchinghandler.Handler
}

func New(cfg *config.Config, svc *Services) *Server {

	return &Server{
		config:          cfg,
		userHandler:     userhandler.New(*svc.Authn, *svc.User, *svc.UserValidator, cfg.Auth),
		adminHandler:    adminhandler.New(*svc.Authn, *svc.Admin, cfg.Auth, *svc.Authz),
		matchingHandler: matchinghandler.New(*svc.Authn, *svc.Matching, *svc.MatchingValidator, cfg.Auth),
	}
}

func (s *Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	s.userHandler.SetRoutes(e)
	s.adminHandler.SetRoutes(e)
	s.matchingHandler.SetRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
