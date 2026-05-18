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

	// http.HandleFunc("/users/register", handleRegister)
	// http.HandleFunc("/users/login", handleLogin)
	// http.HandleFunc("/users/profile", handleProfile)
	// server := http.Server{Addr: ":8080"}
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}

// func handleRegister(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		fmt.Fprintf(w, `"error":"invalid method"`)

// 		return
// 	}

// }

// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		fmt.Fprintf(w, `"error":"invalid method"`)

// 		return
// 	}

// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		w.Write([]byte(`{"error": "could not read"}`))
// 		log.Println(err.Error())

// 		return
// 	}

// 	var lReq userservice.LoginRequest

// 	err = json.Unmarshal(data, &lReq)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
// 		log.Println(err.Error())

// 		return
// 	}
// 	authService := authservice.New(AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireTime, RefreshTokenExpireTime)

// 	mysqlRepo := mysql.New()
// 	uService := userservice.New(mysqlRepo, authService)
// 	resp, err := uService.Login(lReq)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
// 		log.Println(err.Error())

// 		return
// 	}

// 	data, err = json.Marshal(resp)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
// 		log.Println(err.Error())

// 		return
// 	}

// 	fmt.Fprintf(w, "%s", string(data))
// }

// func handleProfile(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		fmt.Fprintf(w, `"error":"invalid method"`)

// 		return
// 	}

// 	authService := authservice.New(AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireTime, RefreshTokenExpireTime)

// 	auth := r.Header.Get("Authorization")
// 	claims, err := authService.ParseToken(auth)

// 	if err != nil {
// 		fmt.Fprintf(w, `"error":"invalid token "`)
// 	}

// 	mysqlRepo := mysql.New()
// 	userSvc := userservice.New(mysqlRepo, authService)

// 	profile, err := userSvc.GetProfile(userservice.GetProfileRequest{UserID: claims.UserID})
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
// 		log.Println(err.Error())

// 		return
// 	}

// 	data, err := json.Marshal(profile)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
// 		log.Println(err.Error())

// 		return
// 	}

// 	fmt.Fprintf(w, "%s", string(data))
// }

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
