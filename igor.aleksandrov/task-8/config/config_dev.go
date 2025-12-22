//go:build dev

package config

import _ "embed"

//go:embed dev.yaml
var rawDev []byte

func init() {
	currentSource = func() []byte {
		return rawDev
	}
}
