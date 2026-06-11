package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/SoroushBeigi/knowledge-game/adapter/redis"
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqlac"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqluser"
	"github.com/SoroushBeigi/knowledge-game/repository/redis/redismatching"
	"github.com/SoroushBeigi/knowledge-game/scheduler"
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

	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Add(1)
	go func() {
		sch := scheduler.New()
		sch.Start(done, &wg)
	}()

	server := httpserver.New(cfg, setupServices(cfg))
	go func() {

		server.Serve()

	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)
	<-quitChan
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	log.Println("Interrupt received. shutting down gracefully...")
	done <- true
	time.Sleep(cfg.Application.GracefulShutdownTimeout)

	<-ctxWithTimeout.Done()

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
