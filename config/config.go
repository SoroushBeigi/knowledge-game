package config

import "github.com/SoroushBeigi/knowledge-game/service/authservice"

type HTTPServer struct {
	Port int
}

type Config struct {
	Auth       authservice.Config
	HTTPServer HTTPServer
}
