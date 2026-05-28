package authservice

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
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

type Config struct {
	SignKey           string
	AccessExpireTime  time.Duration
	RefreshExpireTime time.Duration
	AccessSubject     string
	RefreshSubject    string
}
type Service struct {
	config Config
}

func New(config Config) *Service {

	return &Service{
		config: config,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessExpireTime, s.config.AccessSubject)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshExpireTime, s.config.RefreshSubject)
}

func (s Service) ParseToken(tokenStr string) (*Claims, error) {
	const op = "authservice.ParseToken"
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {

		return []byte(s.config.SignKey), nil
	})

	if err != nil {
		log.Println(op, "error while parsing JWT:", err)
		return nil, richerror.New(op).WithCode(richerror.ForbiddenCode).WithMessage(err.Error())
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Printf("userID: %v, expires at: %v\n", claims.UserID, claims.ExpiresAt)

		return claims, nil

	} else {
		log.Println(op, "error while parsing JWT")

		return nil, richerror.New(op).WithCode(richerror.ForbiddenCode).WithMessage(err.Error())
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
	tokenStr, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {

		return "", err
	}

	return tokenStr, nil
}
