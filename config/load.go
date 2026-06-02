package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(ymlConfigPath string) *Config {
	godotenv.Load(".env")
	var k = koanf.New(".")

	err := k.Load(confmap.Provider(map[string]any{
		"auth.access_subject":  AccessTokenSubject,
		"auth.refresh_subject": RefreshTokenSubject,
	}, "."), nil)

	if err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	if err := k.Load(file.Provider(ymlConfigPath), yaml.Parser()); err != nil {
		log.Printf("warning: error loading config.yml: %v", err)
	}

	err = k.Load(env.Provider("GAMEAPP_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "GAMEAPP_")
		s = strings.ToLower(s)
		return strings.Replace(s, "__", ".", -1)
	}), nil)
	if err != nil {
		log.Fatalf("error loading env config: %v", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	return &cfg
}
