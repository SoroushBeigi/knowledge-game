package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("exiting")
			return

		default:
			now := time.Now()
			fmt.Println("Scheduled: ", now)
			time.Sleep(3 * time.Second)
		}
	}
}
