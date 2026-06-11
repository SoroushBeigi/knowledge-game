package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	Router          *echo.Echo
	httpServer      *http.Server
}

func New(cfg *config.Config, svc *Services) *Server {

	return &Server{
		Router:          echo.New(),
		config:          cfg,
		userHandler:     userhandler.New(*svc.Authn, *svc.User, *svc.UserValidator, cfg.Auth),
		adminHandler:    adminhandler.New(*svc.Authn, *svc.Admin, cfg.Auth, *svc.Authz),
		matchingHandler: matchinghandler.New(*svc.Authn, *svc.Matching, *svc.MatchingValidator, cfg.Auth),
	}
}

func (s *Server) Serve() {

	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	s.userHandler.SetRoutes(s.Router)
	s.adminHandler.SetRoutes(s.Router)
	s.matchingHandler.SetRoutes(s.Router)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.HTTPServer.Port),
		Handler: s.Router,
	}

	log.Fatal(s.httpServer.ListenAndServe())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
