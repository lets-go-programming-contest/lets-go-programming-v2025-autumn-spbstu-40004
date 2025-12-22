//go:build dev

package config

import _ "embed"

//go:embed configurations/dev.yaml
var ConfigYaml []byte
