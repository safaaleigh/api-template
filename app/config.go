package app

import "github.com/BurntSushi/toml"

// DatabaseConfig values for database configuration
type DatabaseConfig struct {
	Name string
	User string
}

// ServerConfig values for server configuration
type ServerConfig struct {
	Port int
}

// Config consolidated configuration values
type Config struct {
	DB     DatabaseConfig `toml:"database"`
	Server ServerConfig   `toml:"server"`
}

// LoadConfig loads configuration values
func LoadConfig() Config {

	var config Config
	if _, err := toml.DecodeFile("./app/config.toml", &config); err != nil {
		panic(err)
	}

	return config
}
