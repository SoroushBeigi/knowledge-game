package config

import (
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/authnservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Auth       authnservice.Config `koanf:"auth"`
	HTTPServer HTTPServer         `koanf:"http_server"`
	MySQL      mysql.Config       `koanf:"mysql"`
}
