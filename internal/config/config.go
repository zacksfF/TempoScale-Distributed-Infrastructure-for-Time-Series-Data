package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server  serverConfig
	DB      dbconf
	Setting SettingConf
}

type serverConfig struct {
	Port      string `env:TEMPOSCALE_APP_Port, required`
	IP        string `env:TEMPOSCALE_APP_IP, required`
	SecretKey byte   `env:TEMPOSCALE_APP_SECRET_KEY, required`
}

type dbconf struct {
	Host               string `env:TEMPOSCALE_DB_HOST, required`
	Port               string `env:TEMPOSCALE_DB_PORT, required`
	User               string `env:TEMPOSCALE_DB_USER, required`
	Password           string `env:TEMPOSCALE_DB_PASSWORD, required`
	DBName             string `env:TEMPOSCALE_DB_NAME, required`
	HasAutoMigarations bool   `env:TEMPOSCALE_DB_HAS_AUTO_MIGRATIONS, default=true`
}

type SettingConf struct {
	HasDebugging bool `env:TEMPOSCALE_HAS_DEBBUGING, default=true`
	HasAnlyzer   bool `env:TEMPOSCALE_HAS_ANALYZER, default=true`
}

// reading environment variables and decoding them into a Config
func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
