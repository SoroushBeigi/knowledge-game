package httpserver

import (
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, auth authservice.Service, user userservice.Service) *Server {
	return &Server{
		config:  config,
		authSvc: auth,
		userSvc: user,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	uGroup:= e.Group("/users")

	uGroup.POST("/register", s.userRegister)
	uGroup.POST("/login", s.userLogin)
	uGroup.GET("/profile", s.userProfile)
	
	log.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
