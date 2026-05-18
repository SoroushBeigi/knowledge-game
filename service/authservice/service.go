package authservice

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/golang-jwt/jwt/v5"
)



type AuthParser interface {
	ParseToken(tokenStr string) (*Claims, error)
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func (c Claims) Validate() error {
	if c.UserID < 1 {

		return errors.New("missing or invalid user id")
	}

	return nil
}

type Service struct {
	signKey           string
	accessExpireTime  time.Duration
	refreshExpireTime time.Duration
	accessSubject     string
	refreshSubject    string
}

func New(accessSubject, refreshSubject string,
	accessExpireTime, refreshExpireTime time.Duration) *Service {

	signKey := os.Getenv("SIGN_SECRET")

	return &Service{
		signKey:           signKey,
		accessSubject:     accessSubject,
		refreshSubject:    refreshSubject,
		accessExpireTime:  accessExpireTime,
		refreshExpireTime: refreshExpireTime,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.accessExpireTime, s.accessSubject)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.refreshExpireTime, s.refreshSubject)
}

func (s Service) ParseToken(tokenStr string) (*Claims, error) {
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {

		return []byte(s.signKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Printf("userID: %v, expires at: %v\n", claims.UserID, claims.ExpiresAt)

		return claims, nil

	} else {
		log.Println("error while parsing JWT")

		return nil, err
	}

}

func (s Service) createToken(userID uint, expireDuration time.Duration, subject string) (string, error) {

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			Subject:   subject, //access or refresh
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.signKey))
	if err != nil {

		return "", err
	}

	return tokenStr, nil
}
