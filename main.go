package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
)

func main() {
	http.HandleFunc("/users/register", handleRegister)
	server := http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, `"error":"invalid method"`)

		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"error": "could not read"}`))
		log.Println(err.Error())

		return
	}

	var uReq userservice.RegisterRequest

	err = json.Unmarshal(data, &uReq)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
		log.Println(err.Error())

		return
	}

	mysqlRepo := mysql.New()
	uService := userservice.New(mysqlRepo)
	user, err := uService.Register(uReq)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
		log.Println(err.Error())

		return
	}
	
	userJson, _ := json.Marshal(user)
	userJsonRaw := userJson[1 : len(userJson)-1]
	fmt.Fprintf(w, `{"message": "user created successfully", %v}`, string(userJsonRaw))
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
