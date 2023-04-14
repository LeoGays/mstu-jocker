package db

import (
	"fmt"
	"jocer/pkg/cfg"
	"strings"
)

const (
	DriverPostgres = "postgres" // Имя драйвера PostgresSQL
	DriverSQLite3  = "sqlite3"  // Имя драйвера SQLite

	SSLModeDisable = "disable" // Значение опции SSL Mode

	CfgKeyHost         cfg.Key = "DB_HOST"             // Хост СУБД.
	CfgKeyPort         cfg.Key = "DB_PORT"             // Порт СУБД.
	CfgKeySSLMode      cfg.Key = "DB_SSL_MODE"         // Режим SSL.
	CfgKeySSLCertPath  cfg.Key = "DB_SSL_CERT_PATH"    // Путь к сертификату для SSL.
	CfgKeyDriver       cfg.Key = "DB_DRIVER"           // Имя драйвера БД.
	CfgKeyDatabase     cfg.Key = "DB_DATABASE"         // Имя БД.
	CfgKeyDBUser       cfg.Key = "DB_USER"             // Пользователь БД.
	CfgKeyDBPassword   cfg.Key = "DB_PASSWORD"         // Пароль пользователя БД.
	CfgKeyDebug        cfg.Key = "DB_DEBUG"            // Флаг отладки (bool).
	CfgKeyMigrationSrc cfg.Key = "DB_MIGRATION_SOURCE" // Источник скриптов миграции (каталог ФС).
	CfgKeyOptions      cfg.Key = "DB_OPTIONS"          // Опции подключения к БД.

	CfgDefaultHost    = "127.0.0.1"    // Хост СУБД по-умолчанию
	CfgDefaultPort    = "5432"         // Порт хоста СУБД по-умолчанию
	CfgDefaultDriver  = DriverPostgres // Драйвер БД по-умолчанию
	CfgDefaultSSLMode = SSLModeDisable // Режим SSL по-умолчанию
)

var (
	_ ConnectionDescriber = (*Config)(nil)
	_ ConnectionDescriber = (*ConfigFileSource)(nil)
)

type (
	// ConnectionDescriber описание конфига подключения к БД, является универсальным представлением параметров
	// подключения к различным СУБД.
	ConnectionDescriber interface {
		fmt.Stringer

		IsDebug() bool                 // Возвращает флаг отладки.
		MigrationSource() string       // Возвращает источник миграционных скриптов.
		DatabaseName() string          // Возвращает имя БД.
		DataSourceName() string        // Возвращает DSN для создания подключения к БД, пригодный для использования в sql.Open().
		DefaultDataSourceName() string // Возвращает дефолтный DSN для создания подключения к БД, пригодный для использования в sql.Open().
		DriverName() string            // Имя драйвера БД.
	}

	// ConfigFileSource описывает конфигурацию файлового источника данных (подходит для SQLite).
	ConfigFileSource struct {
		Driver    string            // Имя драйвера БД.
		Database  string            // Имя БД.
		Migration string            // Имя источника скриптов миграции БД.
		Debug     bool              // Флаг отладки.
		Options   map[string]string // Дополнительные опции подключения к БД.
	}

	// Config описывает конфигурацию подключения к удаленной СУБД (подходит для PostgresSQL, MySQL и т.д.).
	Config struct {
		Driver      string // Имя драйвера БД.
		Database    string // Имя БД.
		Migration   string // Имя источника скриптов миграции БД.
		Debug       bool   // Флаг отладки.
		Host        string // Хост СУБД.
		Port        string // Порт СУБД.
		User        string // Пользователь БД.
		Password    string // Пароль пользователя БД.
		SSLMode     string // Режим SSL.
		SSLCertPath string // Путь к сертификату для SSL.
	}
)

// IsDebug см. ConnectionDescriber.IsDebug().
func (d *ConfigFileSource) IsDebug() bool {
	return d.Debug
}

// MigrationSource см. ConnectionDescriber.MigrationSource().
func (d *ConfigFileSource) MigrationSource() string {
	return d.Migration
}

// DatabaseName см. ConnectionDescriber.DatabaseName().
func (d *ConfigFileSource) DatabaseName() string {
	return d.Database
}

func (d *ConfigFileSource) String() string {
	return fmt.Sprintf("%s:file:%s", d.Driver, d.Database)
}

// DataSourceName см. ConnectionDescriber.DataSourceName().
func (d *ConfigFileSource) DataSourceName() string {
	var builder strings.Builder

	builder.WriteString("file:")
	builder.WriteString(d.Database)

	delim := '?'
	for key, value := range d.Options {
		builder.WriteRune(delim)
		builder.WriteString(key)
		builder.WriteRune('=')
		builder.WriteString(value)

		delim = '&'
	}

	return builder.String()
}

func (d *ConfigFileSource) DefaultDataSourceName() string {
	return ""
}

// DriverName см. ConnectionDescriber.DriverName().
func (d *ConfigFileSource) DriverName() string {
	return d.Driver
}

// DataSourceName см. ConnectionDescriber.DataSourceName().
func (c *Config) DataSourceName() string {
	if c.SSLMode == SSLModeDisable {
		return fmt.Sprintf(
			"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
			c.Host, c.Port, c.Database, c.SSLMode, c.User, c.Password,
		)
	}

	return fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s user=%s password=%s",
		c.Host, c.Port, c.Database, c.SSLMode, c.SSLCertPath, c.User, c.Password,
	)
}

// DefaultDataSourceName см. ConnectionDescriber.DefaultDataSourceName().
// Имплементация подходит для PostgreSQL.
func (c *Config) DefaultDataSourceName() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		c.Host, c.Port, "postgres", c.SSLMode, c.User, c.Password,
	)
}

// DriverName см. ConnectionDescriber.DriverName().
func (c *Config) DriverName() string {
	return c.Driver
}

func (c *Config) String() string {
	return fmt.Sprintf("%s@%s:%s/%s ssl=%s", c.User, c.Host, c.Port, c.Database, c.SSLMode)
}

// IsDebug см. ConnectionDescriber.IsDebug().
func (c *Config) IsDebug() bool {
	return c.Debug
}

// MigrationSource см. ConnectionDescriber.MigrationSource().
func (c *Config) MigrationSource() string {
	return c.Migration
}

// DatabaseName см. ConnectionDescriber.DatabaseName().
func (c *Config) DatabaseName() string {
	return c.Database
}

// NewCfgSQLiteMem создает параметры подключения к in-memory БД SQLite.
func NewCfgSQLiteMem(database string) *ConfigFileSource {
	return &ConfigFileSource{
		Driver:   DriverSQLite3,
		Database: database,
		Options: map[string]string{
			"mode":  "memory",
			"cache": "shared",
			"_fk":   "1",
		},
	}
}
