package centralbank

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"

	"golang.org/x/net/html/charset"
)

type (
	Currency struct {
		NumCode  int     `json:"num_code"`
		CharCode string  `json:"char_code"`
		Value    float64 `json:"value"`
	}

	tempValCurs struct {
		XMLName xml.Name     `xml:"ValCurs"`
		Valutes []tempValute `xml:"Valute"`
	}

	tempValute struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}
)

func ParseXMLFile(filename string) ([]Currency, error) {
	content, err := readAndPreprocessXML(filename)
	if err != nil {
		return nil, err
	}

	tempValCursData, err := decodeXMLContent(content)
	if err != nil {
		return nil, err
	}

	return convertToCurrencies(tempValCursData)
}

func readAndPreprocessXML(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open XML file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			_ = closeErr
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read XML file: %w", err)
	}

	content = bytes.ReplaceAll(content, []byte(","), []byte("."))

	return content, nil
}

func decodeXMLContent(content []byte) (*tempValCurs, error) {
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = charset.NewReaderLabel

	var tempValCursData tempValCurs

	err := decoder.Decode(&tempValCursData)
	if err != nil {
		return nil, fmt.Errorf("cannot parse XML: %w", err)
	}

	return &tempValCursData, nil
}

func convertToCurrencies(tempValCursData *tempValCurs) ([]Currency, error) {
	currencies := make([]Currency, 0, len(tempValCursData.Valutes))

	for _, tempValute := range tempValCursData.Valutes {
		currency, err := convertValuteToCurrency(tempValute)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func convertValuteToCurrency(tempValute tempValute) (Currency, error) {
	numCode, err := parseNumCode(tempValute.NumCode)
	if err != nil {
		return Currency{}, fmt.Errorf("cannot parse NumCode for currency %s: %w", tempValute.CharCode, err)
	}

	value, err := parseValue(tempValute.Value)
	if err != nil {
		return Currency{}, fmt.Errorf("cannot parse Value for currency %s: %w", tempValute.CharCode, err)
	}

	return Currency{
		NumCode:  numCode,
		CharCode: tempValute.CharCode,
		Value:    value,
	}, nil
}

func parseNumCode(numCodeStr string) (int, error) {
	if numCodeStr == "" {
		return 0, nil
	}

	numCode, err := strconv.Atoi(numCodeStr)
	if err != nil {
		return 0, fmt.Errorf("strconv.Atoi: %w", err)
	}

	return numCode, nil
}

func parseValue(valueStr string) (float64, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	return value, nil
}
