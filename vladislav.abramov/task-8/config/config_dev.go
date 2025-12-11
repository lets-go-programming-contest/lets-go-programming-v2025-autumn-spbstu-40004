//go:build dev

package config

func Load() (*Config, error) {
	return loadConfig(devConfig)
}
