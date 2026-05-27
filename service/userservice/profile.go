package userservice

import (
	"log"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

func (s Service) GetProfile(req dto.GetProfileRequest) (dto.GetProfileResponse, error) {
	const op = "userservice.GetProfile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		log.Println("Service Profile:", err)

		return dto.GetProfileResponse{},
			richerror.New(op).WithErr(err).WithMetaData(map[string]any{"req": req})
	}

	return dto.GetProfileResponse{Name: user.Name}, nil
}
