package userhandler

import (
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

type Handler struct {
	authConfig    authnservice.Config
	authSvc       authnservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authnservice.Service,
	userSvc userservice.Service,
	uv uservalidator.Validator,
	authConfig authnservice.Config,
) *Handler {

	return &Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: uv,
		authConfig:   authConfig,
	}
}
