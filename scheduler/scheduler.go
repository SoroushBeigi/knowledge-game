package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/go-co-op/gocron/v2"
)

type Config struct {
	IntervalSecond int `koanf:"interval_second"`
}

type Scheduler struct {
	sch      gocron.Scheduler
	matchSvc matchingservice.Service
	config   Config
}

func New(matchSvc matchingservice.Service, conf Config) (*Scheduler, error) {
	sch, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		sch:      sch,
		matchSvc: matchSvc,
		config:   conf,
	}, nil
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := s.sch.NewJob(
		gocron.DurationJob(
			time.Second*time.Duration(s.config.IntervalSecond),
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	_, err := s.matchSvc.MatchWaitedUsers(ctx, dto.MatchWaitedUsersRequest{})
	if err != nil {
		log.Println("MatchWaitedUsers scheduler error")
		//TODO: update metrics
	}
}
