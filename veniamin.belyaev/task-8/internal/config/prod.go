//go:build prod || !dev

package config

import _ "embed"

//go:embed configurations/prod.yaml
var configYamlData []byte
