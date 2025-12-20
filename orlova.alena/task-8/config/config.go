package config

// Config структура для хранения конфигурации
type Config struct {
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

func GetConfig() Config {
	return getDefaultConfig()
}
