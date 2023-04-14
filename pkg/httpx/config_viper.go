package httpx

import (
	"github.com/spf13/viper"
	"jocer/pkg/cfg"
	"jocer/pkg/cfg/viperx"
)

func CfgFromViper(loader *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	return &Config{
		Port:              viperx.Get(loader, CfgKeyPort.Map(keyMapping...), CfgDefaultPort),
		ReadHeaderTimeout: CfgDefaultReadHeaderTimeout,
	}
}
