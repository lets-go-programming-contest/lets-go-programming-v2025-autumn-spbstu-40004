//go:build !dev

package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed prod.yaml
var prodConfigData []byte

func getDefaultConfig() Config {
	var cfg Config

	err := yaml.Unmarshal(prodConfigData, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse prod config: %v", err))
	}

	return cfg
}
