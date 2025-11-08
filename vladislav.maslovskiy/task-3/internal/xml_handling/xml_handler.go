package xmlhandling

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumericalCode int     `json:"num_code"  xml:"NumCode"`
	CharacterCode string  `json:"char_code" xml:"CharCode"`
	Value         float64 `json:"value"     xml:"Value"`
}

type CurrencyList struct {
	Currencies []Currency `xml:"Valute"`
}

func transformFileReader(file *os.File, searchChar, replaceChar string) (io.Reader, error) {
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("transforming file reader: %w", err)
	}

	modifiedContent := strings.ReplaceAll(string(fileContent), searchChar, replaceChar)

	return strings.NewReader(modifiedContent), nil
}

func ParseXML(filePath string) ([]Currency, error) {
	var currencyData CurrencyList

	fileHandle, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("i/o xml: %w", err)
	}

	defer func() {
		if err := fileHandle.Close(); err != nil {
			panic(err.Error())
		}
	}()

	modifiedReader, err := transformFileReader(fileHandle, ",", ".")
	if err != nil {
		return nil, err
	}

	xmlDecoder := xml.NewDecoder(modifiedReader)
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	if err := xmlDecoder.Decode(&currencyData); err != nil {
		return nil, fmt.Errorf("decoder: %w", err)
	}

	currencyCount := len(currencyData.Currencies)
	currencySlice := make([]Currency, currencyCount)

	for position := range currencyCount {
		currencySlice[position] = currencyData.Currencies[position]
	}

	return currencySlice, nil
}
