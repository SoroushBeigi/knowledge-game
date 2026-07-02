package presenceservice

import (
	"context"
	"fmt"
	"time"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

type Config struct {
	Prefix string        `koanf:"prefix"`
	Expire time.Duration `koans:"expire"`
}

type Repository interface {
	Upsert(ctx context.Context, key string, timestamp int64, expireDuration time.Duration) error
}

type Service struct {
	config Config
	repo   Repository
}

func New(config Config, repo Repository) *Service {
	return &Service{config: config, repo: repo}
}

func (s Service) UpsertPresence(ctx context.Context, req dto.UpsertPresenceRequest) (dto.UpsertPresenceResponse, error) {
	const op = "presenceservice.UpsertPresence"
	err := s.repo.Upsert(ctx,
		fmt.Sprintf("%s:%d", s.config.Prefix, req.UserID),
		req.Timestamp,
		s.config.Expire,
	)

	if err != nil {
		return dto.UpsertPresenceResponse{}, richerror.New(op).WithErr(err)
	}

	return dto.UpsertPresenceResponse{}, nil

}

func (s Service) GetPresence(ctx context.Context, req dto.GetPresenceRequest) (dto.GetPresenceResponse, error) {
	//TODO: implement!!!
	//TODO: implement!!!
	return dto.GetPresenceResponse{Items: []dto.GetPresenceItem{
		{UserID: 1, Timestamp: 12315616},
		{UserID: 2, Timestamp: 45345645354},
		{UserID: 3, Timestamp: 456485654},
	}}, nil
}
