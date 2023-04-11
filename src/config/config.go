package config

import "os"

type DatabaseConfig struct {
	Url string `json:"url"`
}

// Config struct for structuring the config data, can be extended accordingly
type Config struct {
	FunctionName string         `json:"function_name"`
	Database     DatabaseConfig `json:"database"`
}

// NewConfig  a function to read the configuration file and return config struct
func NewConfig() Config {
	return Config{
		FunctionName: os.Getenv("FUNCTION_NAME"),
		Database: DatabaseConfig{
			Url: os.Getenv("DATABASE_URL"),
		},
	}
}
