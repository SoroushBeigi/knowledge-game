package config

import (
	"time"

	"github.com/SoroushBeigi/knowledge-game/adapter/redis"
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
	"github.com/SoroushBeigi/knowledge-game/service/matchingservice"
	"github.com/SoroushBeigi/knowledge-game/service/presenceservice"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"shutdown_timeout"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application Application            `koanf:"application"`
	Auth        authnservice.Config    `koanf:"auth"`
	HTTPServer  HTTPServer             `koanf:"http_server"`
	MySQL       mysql.Config           `koanf:"mysql"`
	Matching    matchingservice.Config `koanf:"matching"`
	Redis       redis.Config           `koanf:"redis"`
	Presence    presenceservice.Config `koanf:"presence"`
}
