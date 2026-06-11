package config

import "time"

var defaultConf = map[string]any{
	"auth.access_subject":          AccessTokenSubject,
	"auth.refresh_subject":         RefreshTokenSubject,
	"application.shutdown_timeout": time.Second * 5,
}
