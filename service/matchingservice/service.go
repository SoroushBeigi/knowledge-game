package matchingservice

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"github.com/SoroushBeigi/knowledge-game/pkg/timestamp"
	funk "github.com/thoas/go-funk"
)

type Config struct {
	WaitTimeout time.Duration `koanf:"waiting_timeout"`
}

// TODO: add context
type Repo interface {
	AddToWaitingList(userID uint, cat entity.Category) error
	GetWaitingListByCategory(ctx context.Context, cat entity.Category) ([]entity.WaitingMember, error)
}

type PresenceClient interface {
	GetPresence(ctx context.Context, req dto.GetPresenceRequest) (dto.GetPresenceResponse, error)
}

type Service struct {
	config         Config
	repo           Repo
	presenceClient PresenceClient
}

func New(conf Config, repo Repo, pc PresenceClient) *Service {
	return &Service{
		config:         conf,
		repo:           repo,
		presenceClient: pc,
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

func (s Service) MatchWaitedUsers(ctx context.Context, _ dto.MatchWaitedUsersRequest) (dto.MatchWaitedUsersResponse, error) {

	var wg sync.WaitGroup
	for _, cat := range entity.AllCats() {
		wg.Add(1)
		go s.match(ctx, cat, &wg)
	}

	wg.Wait()

	return dto.MatchWaitedUsersResponse{}, nil
}

func (s Service) match(ctx context.Context, cat entity.Category, wg *sync.WaitGroup) {
	const op = "matchingservice.match"

	defer wg.Done()

	waitingMembers, err := s.repo.GetWaitingListByCategory(ctx, entity.TechCat)

	if err != nil {
		log.Println("match service error for cat: ", cat)
		//TODO: update metrics
		return
	}

	userIDs := make([]uint, 0)
	for _, wm := range waitingMembers {
		userIDs = append(userIDs, wm.UserID)
	}

	if len(userIDs) < 2 {
		return
	}

	presList, err := s.presenceClient.GetPresence(ctx, dto.GetPresenceRequest{UserIDs: userIDs})

	//TODO: merge presList with waitingMembers based on userID.
	//consider presence timestamp of each user. remove from waitinglist if timestamp is old
	//if ts < timestamp.SecondBeforeNow(10) {
	//remove from list
	//			}

	var list = make([]entity.WaitingMember, 0)

	presenceUserIDs := make([]uint, len(waitingMembers))
	for _, item := range presList.Items {
		presenceUserIDs = append(presenceUserIDs, item.UserID)
	}

	for _, wm := range waitingMembers {
		if funk.ContainsUInt(presenceUserIDs, wm.UserID) &&
			wm.Timestamp < timestamp.SecondBeforeNow(10) {
			list = append(list, wm)
		} else {
			//TODO: remove from wait list
		}
	}

	for i := 0; i < len(list)-1; i = i + 2 {

		mu := entity.MatchedUsers{
			Category: cat,
			UserIDs:  []uint{list[i].UserID, list[i+1].UserID},
		}

		fmt.Println(mu)
	}
}
