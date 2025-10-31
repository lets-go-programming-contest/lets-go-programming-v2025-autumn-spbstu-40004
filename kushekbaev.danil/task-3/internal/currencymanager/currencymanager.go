package currencymanager

import (
	"encoding/json"
	"encoding/xml"
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
		return currenciesData, err
	}

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&currenciesData)
	if err != nil {
		return currenciesData, err
	}

	err = file.Close()
	if err != nil {
		return currenciesData, err
	}

	return currenciesData, nil
}

func Write(path string, currencies currencyparser.Currencies) error {
	data, err := json.MarshalIndent(currencies, "", "\t")
	if err != nil {
		return err
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, directoryPermission); err != nil {
		return err
	}

	err = os.WriteFile(path, data, filePermission)
	if err != nil {
		return err
	}

	return nil
}
