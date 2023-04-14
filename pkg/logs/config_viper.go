package logs

import (
	"github.com/spf13/viper"
	"jocer/pkg/cfg"
	"jocer/pkg/cfg/viperx"
)

// CfgFromViper загружает конфиг с помощью viper.
func CfgFromViper(v *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	return &Config{
		Level:  viperx.Get(v, CfgKeyLevel.Map(keyMapping...), CfgDefaultLevel),
		Pretty: viperx.Get(v, CfgKeyPretty.Map(keyMapping...), false),
	}
}
