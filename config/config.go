package config

import (
	"jocer/pkg/cfg/viperx"
	"jocer/pkg/db"
	"jocer/pkg/httpx"
)

const appPrefix = "JOCKER"

type Config struct {
	HTTP *httpx.Config
	DB   *db.Config
}

func LoadFromViper() *Config {
	loader := viperx.EnvLoader(appPrefix)

	return &Config{
		HTTP: httpx.CfgFromViper(loader),
		DB:   db.CfgFromViperTest(loader),
	}
}
