package main

import (
	"github.com/SoroushBeigi/knowledge-game/adapter/redis"
	"github.com/SoroushBeigi/knowledge-game/config"
	"github.com/SoroushBeigi/knowledge-game/repository/redis/redispresence"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
	"github.com/SoroushBeigi/knowledge-game/transport/grpcserver/presenceserver"
)

func main() {
	cfg := config.Load("config.yml")
	redisAdapter := redis.New(cfg.Redis)
	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.Presence, presenceRepo)

	server := presenceserver.New(*presenceSvc)

	server.Start()
}
