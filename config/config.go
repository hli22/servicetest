package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	*viper.Viper
}

var conf *Config

// Load configurations
func Load() {
	// load config from env variables
	conf = &Config{viper.New()}

	// Plain text
	conf.SetDefault("text_string", "Default text string")

	// vault setting
	conf.SetDefault("encrypt_string", "Default encrypt string")

	// fileee setting
	conf.SetDefault("file_string", "Default file string")
}

var once sync.Once

// GetConfig ...
func GetConfig() *Config {
	if conf == nil {
		once.Do(Load)
	}
	return conf
}

// IsDebugging ...
func IsDebugging() bool {
	return conf.GetString("DEBUG") != ""
}
