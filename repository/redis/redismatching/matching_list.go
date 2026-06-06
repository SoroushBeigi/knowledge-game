package redismatching

import (
	"context"
	"fmt"
	"time"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"github.com/redis/go-redis/v9"
)

const waitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, cat entity.Category) error {
	const op = "redismatching.AddToWaitingList"

	client := d.adapter.Client()

	_, err := client.ZAdd(context.Background(), fmt.Sprintf("%s:%s", waitingListPrefix, cat), redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: fmt.Sprintf("%d", userID),
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithCode(richerror.UnexpectedCode)
	}

	return nil

}
