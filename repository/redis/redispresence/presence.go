package redispresence

import (
	"context"
	"time"

	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

func (d DB) Upsert(ctx context.Context, key string, timestamp int64, expireDuration time.Duration) error {
	const op = "redispresence.Upsert"
	_, err := d.adapter.Client().Set(ctx, key, timestamp, expireDuration).Result()
	
	if err != nil {
		return richerror.New(op).WithErr(err).WithCode(richerror.UnexpectedCode)
	}

	return nil
}
