//go:build dev

package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed dev.yaml
var devConfigData []byte

func getDefaultConfig() Config {
	var cfg Config

	err := yaml.Unmarshal(devConfigData, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse dev config: %v", err))
	}

	return cfg
}
