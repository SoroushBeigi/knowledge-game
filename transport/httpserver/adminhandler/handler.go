package adminhandler

import (
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
)

type Handler struct {
	authConfig authnservice.Config
	authnSvc   authnservice.Service
	authzSvc   authzservice.Service
	adminSvc   adminservice.Service
}

func New(authnSvc authnservice.Service,
	adminSvc adminservice.Service,
	authConfig authnservice.Config,
	authzSvc authzservice.Service,
) *Handler {

	return &Handler{
		authnSvc:   authnSvc,
		adminSvc:   adminSvc,
		authConfig: authConfig,
		authzSvc:   authzSvc,
	}
}
