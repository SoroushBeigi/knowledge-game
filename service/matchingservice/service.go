package matchingservice

import (
	"time"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

type Config struct {
	WaitTimeout time.Duration `koanf:"waiting_timeout"`
}

type Repo interface {
	AddToWaitingList(userID uint, cat entity.Category) error
}

type Service struct {
	config Config
	repo   Repo
}

func New(conf Config, repo Repo) *Service {
	return &Service{
		config: conf,
		repo:   repo,
	}
}

func (s Service) AddToWaitingList(req dto.AddToWaitingListRequest) (dto.AddToWaitingListResponse, error) {
	const op = "matchingservice.addToWaitList"

	err := s.repo.AddToWaitingList(req.UserID, req.Category)

	if err != nil {
		return dto.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithCode(richerror.UnexpectedCode)
	}

	return dto.AddToWaitingListResponse{Timeout: s.config.WaitTimeout}, nil

}

func (s Service) MatchWaitedUsers(req dto.MatchWaitedUsersRequest) (dto.MatchWaitedUsersResponse, error) {
	return dto.MatchWaitedUsersResponse{}, nil
}
