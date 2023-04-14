package httpx

import (
	"jocer/pkg/cfg"
	"net"
	"time"
)

const (
	CfgKeyPort cfg.Key = "HTTP_SERVER_PORT"

	CfgDefaultPort              = "8080"
	CfgDefaultReadHeaderTimeout = 30 * time.Second
)

type Config struct {
	Port              string
	ReadHeaderTimeout time.Duration
}

// Addr returns server address in format ":<port>".
func (c Config) Addr() string {
	return net.JoinHostPort("", c.Port)
}
