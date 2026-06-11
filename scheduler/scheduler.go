package scheduler

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/go-co-op/gocron/v2"
)

type Scheduler struct {
	sch      gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(matchSvc matchingservice.Service) (*Scheduler, error) {
	sch, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		sch:      sch,
		matchSvc: matchSvc,
	}, nil
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := s.sch.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(s.MatchWaitedUsers),
	)
	if err != nil {
		log.Println("Scheduler Error: ", err)
	}

	s.sch.Start()

	<-done

	fmt.Println("stop scheduler..")
	err = s.sch.StopJobs()
	if err != nil {
		log.Println("Scheduler StopJobs Error: ", err)
	}

}

func (s Scheduler) MatchWaitedUsers() {
	log.Println("MatchWaitedUsers")
	s.matchSvc.MatchWaitedUsers(dto.MatchWaitedUsersRequest{})
}
