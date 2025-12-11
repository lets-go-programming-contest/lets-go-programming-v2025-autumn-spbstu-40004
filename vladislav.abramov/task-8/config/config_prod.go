package config

func Load() (*Config, error) {
	return loadConfig(prodConfig)
}
