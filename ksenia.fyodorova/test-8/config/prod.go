//go:build !dev

package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed prod.yaml
var productionData []byte

func Load() (*AppConfig, error) {
	cfg := &AppConfig{}

	err := yaml.Unmarshal(productionData, cfg)
	if err != nil {
		return nil, &ConfigError{Msg: "не удалось разобрать YAML конфигурации", Err: err}
	}

	return cfg, nil
}

type ConfigError struct {
	Msg string
	Err error
}

func (e *ConfigError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}
