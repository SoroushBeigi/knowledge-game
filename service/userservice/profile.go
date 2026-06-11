package userservice

import (
	"context"
	"log"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

func (s Service) GetProfile(ctx context.Context, req dto.GetProfileRequest) (dto.GetProfileResponse, error) {
	const op = "userservice.GetProfile"

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		log.Println("Service Profile:", err)

		return dto.GetProfileResponse{},
			richerror.New(op).WithErr(err).WithMetaData(map[string]any{"req": req})
	}

	return dto.GetProfileResponse{Name: user.Name}, nil
}
