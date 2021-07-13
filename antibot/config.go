package antibot

import (
	"time"

	"github.com/ardanlabs/conf"
)

type Config struct {
	conf.Version

	LogLevel string `conf:"default:info,env:LOG_LEVEL"`

	Web struct {
		Port            string        `conf:"default:8081,env:WEB_PORT"`
		Host            string        `conf:"default:0.0.0.0,env:WEB_HOST"`
		ReadTimeout     time.Duration `conf:"default:5s"`
		WriteTimeout    time.Duration `conf:"default:5s"`
		ShutdownTimeout time.Duration `conf:"default:5s"`
	}
}
