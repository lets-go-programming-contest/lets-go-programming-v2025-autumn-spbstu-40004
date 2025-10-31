package currencymanager

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Z-1337/task-3/internal/currencyparser"
	"golang.org/x/net/html/charset"
)

const (
	filePermission      = 0o644
	directoryPermission = 0o755
)

func Read(path string) (currencyparser.Currencies, error) {
	var currenciesData currencyparser.Currencies

	file, err := os.Open(path)
	if err != nil {
		return currenciesData, fmt.Errorf("opening file: %w", err)
	}

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&currenciesData)
	if err != nil {
		return currenciesData, fmt.Errorf("decoding xml: %w", err)
	}

	err = file.Close()
	if err != nil {
		return currenciesData, fmt.Errorf("closing file: %w", err)
	}

	return currenciesData, nil
}

func Write(path string, currencies currencyparser.Currencies) error {
	data, err := json.MarshalIndent(currencies.AllCurrencies, "", "\t")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, directoryPermission); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	err = os.WriteFile(path, data, filePermission)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
