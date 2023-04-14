package presenters

import (
	"github.com/spf13/viper"
	"jocer/pkg/cfg"
	"jocer/pkg/cfg/viperx"
)

func CfgFromViper(loader *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	loader.SetDefault(CfgKeySecuredKeywords.String(), CfgDefaultKeySecuredKeyword)

	return &Config{
		SecuredKeywords: loader.GetStringSlice(CfgKeySecuredKeywords.String()),
		MaxStringLength: viperx.Get(loader, CfgKeyMaxStringLength.Map(keyMapping...), CfgDefaultMaxStringLength),
	}
}
