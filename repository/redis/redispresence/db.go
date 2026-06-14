package redispresence

import "github.com/SoroushBeigi/knowledge-game/adapter/redis"

type DB struct {
	adapter *redis.Adapter
}

func New(adapter *redis.Adapter) *DB {

	return &DB{adapter: adapter}
}
