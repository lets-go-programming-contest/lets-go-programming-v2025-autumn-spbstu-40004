package currenciesprocessing

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ParseCurrencies(path string) (*Currencies, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var (
		currencies Currencies
		decoder    = xml.NewDecoder(file)
	)

	err = decoder.Decode(&currencies)
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml file: %w", err)
	}

	return &currencies, nil
}
