package matchinghandler

import (
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"github.com/SoroushBeigi/knowledge-game/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authnservice.Config
	authSvc           authnservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	presenceService   presenceservice.Service
}

func New(authSvc authnservice.Service,
	matchingSvc matchingservice.Service,
	mv matchingvalidator.Validator,
	authConfig authnservice.Config,
	presenceSvc presenceservice.Service,
) *Handler {

	return &Handler{
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		authConfig:        authConfig,
		matchingValidator: mv,
		presenceService:   presenceSvc,
	}
}
