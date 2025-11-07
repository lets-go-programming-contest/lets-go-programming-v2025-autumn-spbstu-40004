package xmlhandling

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	regexForReplacingCommasToDots = `(<[A-Za-z0-9]+>[0-9]+),([0-9]+<\/[A-Za-z]+>)`
)

type Currency struct {
	NumericalCode int     `json:"num_code"  xml:"NumCode"`
	CharacterCode string  `json:"char_code" xml:"CharCode"`
	Value         float64 `json:"value"     xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []Currency `xml:"Valute"`
}

func replaceCommaWithDotInFloat(file *os.File) (io.Reader, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("transforming file reader: %w", err)
	}

	regex := regexp.MustCompile(regexForReplacingCommasToDots)
	transformed := regex.ReplaceAllString(string(content), "$1.$2")

	return strings.NewReader(transformed), nil
}

func ParseXML(filePath string) ([]Currency, error) {
	var currencies CurrenciesXML

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("i/o xml: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	transformedReader, err := replaceCommaWithDotInFloat(file)
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
		currenciesArray[index] = currencies.Currencies[index]
	}

	return currenciesArray, nil
}
