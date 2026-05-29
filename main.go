package main

import (
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

func main() {
	cfg := config.Load("config.yml")

	authSvc, userSvc, uv := setupServices(*cfg)

	server := httpserver.New(*cfg, authSvc, userSvc, uv)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	auth := authservice.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.MySQL)
	user := userservice.New(mysqlRepo, auth)
	uv := uservalidator.New(mysqlRepo)

	return *auth, *user, *uv
}
