package config

import (
	"github.com/SoroushBeigi/knowledge-game/repository/mysql"
	"github.com/SoroushBeigi/knowledge-game/service/authservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Auth       authservice.Config `koanf:"auth"`
	HTTPServer HTTPServer         `koanf:"http_server"`
	MySQL      mysql.Config       `koanf:"mysql"`
}
