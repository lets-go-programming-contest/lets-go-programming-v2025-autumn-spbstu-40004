package xmlhandling

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumericalCode int     `json:"num_code"  xml:"NumCode"`
	CharacterCode string  `json:"char_code" xml:"CharCode"`
	Value         float64 `json:"value"     xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []Currency `xml:"Valute"`
}

func replaceCharInFileReader(file *os.File, oldChar, newChar string) (io.Reader, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("transforming file reader: %w", err)
	}

	transformed := strings.ReplaceAll(string(content), oldChar, newChar)

	return strings.NewReader(transformed), nil
}

func ParseXML(filePath string) ([]Currency, error) {
	var currencies CurrenciesXML

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("i/o: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	transformedReader, err := replaceCharInFileReader(file, ",", ".")
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(transformedReader)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&currencies); err != nil {
		return nil, fmt.Errorf("decoder: %w", err)
	}

	arrayLength := len(currencies.Currencies)
	currenciesArray := make([]Currency, arrayLength)

	for index := range arrayLength {
		currenciesArray[index] = Currency(currencies.Currencies[index])
	}

	sort.Slice(currenciesArray, func(i, j int) bool {
		return currenciesArray[i].Value > currenciesArray[j].Value
	})

	return currenciesArray, nil
}
