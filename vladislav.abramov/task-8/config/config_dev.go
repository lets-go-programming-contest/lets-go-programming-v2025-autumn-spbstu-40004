package config

func Load() *Config {
	return loadConfig("dev.yaml")
}
