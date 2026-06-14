package httpserver

import (
	"github.com/SoroushBeigi/knowledge-game/service/adminservice"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/authzservice"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"github.com/SoroushBeigi/knowledge-game/service/userservice"
	"github.com/SoroushBeigi/knowledge-game/validator/matchingvalidator"
	"github.com/SoroushBeigi/knowledge-game/validator/uservalidator"
)

type Services struct {
	Authn    *authnservice.Service
	User     *userservice.Service
	Admin    *adminservice.Service
	Authz    *authzservice.Service
	Matching *matchingservice.Service
	Presence *presenceservice.Service

	UserValidator     *uservalidator.Validator
	MatchingValidator *matchingvalidator.Validator
}
