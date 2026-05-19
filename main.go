package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/transport/httpserver"
	"github.com/joho/godotenv"
)

const (
	AccessTokenSubject     = "at"
	RefreshTokenSubject    = "rt"
	AccessTokenExpireTime  = time.Hour * 24
	RefreshTokenExpireTime = time.Hour * 24 * 7
)

func main() {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal(envErr)
	}
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:           os.Getenv("SIGN_SECRET"),
			AccessExpireTime:  AccessTokenExpireTime,
			RefreshExpireTime: RefreshTokenExpireTime,
			AccessSubject:     AccessTokenSubject,
			RefreshSubject:    RefreshTokenSubject,
		},
	}
	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()

}

func mainTestDB() {
	mysqlRepo := mysql.New()

	mysqlRepo.Register(entity.User{
		PhoneNumber: "0910101",
		Name:        "Ssss",
	})

	isUnique, err := mysqlRepo.IsPhoneNumberUnique("09010101")
	fmt.Println(isUnique, err)
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	auth := authservice.New(cfg.Auth)
	mysqlRepo := mysql.New()
	user := userservice.New(mysqlRepo, auth)

	return *auth, *user

}
