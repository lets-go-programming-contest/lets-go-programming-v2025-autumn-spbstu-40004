//go:build !dev

package config

import _ "embed"

//go:embed prod.yaml
var rawProd []byte

func init() {
	currentSource = func() []byte {
		return rawProd
	}
}
