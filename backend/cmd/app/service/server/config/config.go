package config

import (
	"net/http"
	"time"
)

type ServerConfig struct {
	Router       http.Handler
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
