package main

import (
	"fmt"

	"github.com/SoroushBeigi/knowledge-game/adapter/redis"
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqlac"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqluser"
	"github.com/SoroushBeigi/knowledge-game/repository/redis/redismatching"
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver"
	"github.com/SoroushBeigi/knowledge-game/validator/matchingvalidator"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

func main() {
	cfg := config.Load("config.yml")

	fmt.Println(cfg)

	server := httpserver.New(cfg, setupServices(cfg))

	server.Serve()
}

func setupServices(cfg *config.Config) *httpserver.Services {
	redisAdapter := redis.New(cfg.Redis)

	mysqlRepo := mysql.New(cfg.MySQL)
	userMysql := mysqluser.New(mysqlRepo)
	acMysql := mysqlac.New(mysqlRepo)
	matchingrepo := redismatching.New(redisAdapter)

	authN := authnservice.New(cfg.Auth)
	authZ := authzservice.New(acMysql)
	user := userservice.New(userMysql, authN)
	admin := adminservice.New()
	matchingSvc := matchingservice.New(cfg.Matching, matchingrepo)

	uv := uservalidator.New(userMysql)
	mv := matchingvalidator.New()

	return &httpserver.Services{
		Authn:             authN,
		User:              user,
		UserValidator:     uv,
		Admin:             admin,
		Authz:             authZ,
		Matching:          matchingSvc,
		MatchingValidator: mv,
	}
}
