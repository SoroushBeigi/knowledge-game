package userhandler

import (
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, uv uservalidator.Validator) *Handler {
	return &Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: uv,
	}
}
