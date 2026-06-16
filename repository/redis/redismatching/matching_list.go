package redismatching

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"github.com/SoroushBeigi/knowledge-game/pkg/timestamp"
	"github.com/redis/go-redis/v9"
)

const waitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, cat entity.Category) error {
	const op = "redismatching.AddToWaitingList"

	client := d.adapter.Client()

	_, err := client.ZAdd(context.Background(), categoryKey(cat), redis.Z{
		Score:  float64(timestamp.Now()),
		Member: fmt.Sprintf("%d", userID),
	}).Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithCode(richerror.UnexpectedCode)
	}

	return nil

}

func (d DB) GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error) {
	const op = "redismatching.GetWaitingListByCategory"

	min := strconv.Itoa(int(timestamp.HourBeforeNow(2)))
	max := strconv.Itoa(int(timestamp.Now()))

	list, err := d.adapter.Client().ZRangeByScoreWithScores(
		ctx, categoryKey(category),
		&redis.ZRangeBy{Min: min, Max: max, Offset: 0, Count: 0},
	).Result()

	if err != nil {
		return nil,
			richerror.New(op).WithErr(err).WithCode(richerror.UnexpectedCode)
	}

	var result = make([]entity.WaitingMember, 0)

	for _, l := range list {
		member, ok := l.Member.(string)
		if !ok {
			log.Println(op, "l.Member TO string")
			continue
		}

		userID, convErr := strconv.Atoi(member)
		if convErr != nil {
			log.Println(op, "member to userID")
			continue
		}

		result = append(result, entity.WaitingMember{
			UserID:    uint(userID),
			Timestamp: int64(l.Score),
			Category:  category,
		})

	}
	return result, nil
}

func categoryKey(cat entity.Category) string {
	return fmt.Sprintf("%s:%s", waitingListPrefix, cat)
}
