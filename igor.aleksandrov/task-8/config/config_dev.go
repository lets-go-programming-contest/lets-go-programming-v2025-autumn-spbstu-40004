//go:build dev

package config

import _ "embed"

//go:embed dev.yaml
var devConfigData []byte

func getConfigData() []byte {
	return devConfigData
}
