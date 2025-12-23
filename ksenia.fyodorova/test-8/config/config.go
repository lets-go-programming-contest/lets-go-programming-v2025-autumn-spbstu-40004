package config

type AppConfig struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}
