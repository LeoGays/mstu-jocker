package db

import (
	"github.com/spf13/viper"
	"jocer/pkg/cfg"
	"jocer/pkg/cfg/viperx"
)

// CfgFromViper загружает и возвращает конфигурацию подключения к СУБД с помощью viper.
func CfgFromViper(loader *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	return &Config{
		Driver:      viperx.Get(loader, CfgKeyDriver.Map(keyMapping...), CfgDefaultDriver),
		Host:        viperx.Get(loader, CfgKeyHost.Map(keyMapping...), CfgDefaultHost),
		Port:        viperx.Get(loader, CfgKeyPort.Map(keyMapping...), CfgDefaultPort),
		Database:    viperx.Get(loader, CfgKeyDatabase.Map(keyMapping...), "test"),
		User:        viperx.Get(loader, CfgKeyDBUser.Map(keyMapping...), "user"),
		Password:    viperx.Get(loader, CfgKeyDBPassword.Map(keyMapping...), "test"),
		SSLMode:     viperx.Get(loader, CfgKeySSLMode.Map(keyMapping...), CfgDefaultSSLMode),
		SSLCertPath: viperx.Get(loader, CfgKeySSLCertPath.Map(keyMapping...), ""),
		Debug:       viperx.Get(loader, CfgKeyDebug.Map(keyMapping...), false),
		Migration:   viperx.Get(loader, CfgKeyMigrationSrc.Map(keyMapping...), ""),
	}
}

// CfgFromViperTest загружает и возвращает конфигурацию подключения к СУБД c маркером test с помощью viper.
func CfgFromViperTest(v *viper.Viper) *Config {
	return CfgFromViper(v, cfg.KeyWithSuffix(cfg.MarkerTest))
}

// CfgFSFromViper загружает и возвращает конфигурацию подключения к файловой БД с помощью viper.
func CfgFSFromViper(loader *viper.Viper, keyMapping ...cfg.KeyMap) *ConfigFileSource {
	return &ConfigFileSource{
		Driver:    viperx.Get(loader, CfgKeyDriver.Map(keyMapping...), CfgDefaultDriver),
		Database:  viperx.Get(loader, CfgKeyDatabase.Map(keyMapping...), ""),
		Debug:     viperx.Get(loader, CfgKeyDebug.Map(keyMapping...), false),
		Migration: viperx.Get(loader, CfgKeyMigrationSrc.Map(keyMapping...), ""),
		Options:   loader.GetStringMapString(CfgKeyOptions.Map(keyMapping...).String()),
	}
}
