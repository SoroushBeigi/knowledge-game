package main

import (
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqlac"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqluser"
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

func main() {
	cfg := config.Load("config.yml")

	authnSvc, userSvc, uv, admin, authzSvc := setupServices(*cfg)

	server := httpserver.New(*cfg, *authnSvc, *userSvc, *uv, *admin, *authzSvc)

	server.Serve()
}

func setupServices(cfg config.Config) (*authnservice.Service, *userservice.Service,
	*uservalidator.Validator, *adminservice.Service, *authzservice.Service,
) {
	authN := authnservice.New(cfg.Auth)

	mysqlRepo := mysql.New(cfg.MySQL)
	userMysql := mysqluser.New(mysqlRepo)
	acMysql := mysqlac.New(mysqlRepo)

	authZ := authzservice.New(acMysql)

	user := userservice.New(userMysql, authN)
	admin := adminservice.New()

	uv := uservalidator.New(userMysql)

	return authN, user, uv, admin, authZ
}
