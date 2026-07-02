package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	presenceClient "github.com/SoroushBeigi/knowledge-game/adapter/presence"
	"github.com/SoroushBeigi/knowledge-game/adapter/redis"
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqlac"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql/mysqluser"
	"github.com/SoroushBeigi/knowledge-game/repository/redis/redismatching"
	"github.com/SoroushBeigi/knowledge-game/repository/redis/redispresence"
	"github.com/SoroushBeigi/knowledge-game/scheduler"
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver"
	"github.com/SoroushBeigi/knowledge-game/validator/matchingvalidator"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load("config.yml")

	var wg sync.WaitGroup
	done := make(chan bool)

	conn, err := grpc.Dial(":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	svc := setupServices(cfg, conn)
	server := httpserver.New(cfg, svc)

	go func() {
		sch, err := scheduler.New(*svc.Matching, cfg.Scheduler)
		if err != nil {
			log.Println("FATAL: Scheduler ERR: ", err)
		}

		wg.Add(1)
		sch.Start(done, &wg)
	}()

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

	wg.Wait()
	time.Sleep(2 * time.Second)
}

func setupServices(cfg *config.Config, presenceConn *grpc.ClientConn) *httpserver.Services {
	redisAdapter := redis.New(cfg.Redis)

	mysqlRepo := mysql.New(cfg.MySQL)
	userMysql := mysqluser.New(mysqlRepo)
	acMysql := mysqlac.New(mysqlRepo)
	matchingrepo := redismatching.New(redisAdapter)
	presenceRepo := redispresence.New(redisAdapter)

	presenceAdapter := presenceClient.New(presenceConn)
	presenceSvc := presenceservice.New(cfg.Presence, presenceRepo)

	authN := authnservice.New(cfg.Auth)
	authZ := authzservice.New(acMysql)
	user := userservice.New(userMysql, authN)
	admin := adminservice.New()
	matchingSvc := matchingservice.New(cfg.Matching, matchingrepo, presenceAdapter)

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
		Presence:          presenceSvc,
	}
}
