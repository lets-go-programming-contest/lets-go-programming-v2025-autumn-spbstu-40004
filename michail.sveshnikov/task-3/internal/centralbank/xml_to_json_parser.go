package centralbank

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type FloatComma float64

type (
	Currency struct {
		NumCode  int        `json:"num_code"  xml:"NumCode"`
		CharCode string     `json:"char_code" xml:"CharCode"`
		Value    FloatComma `json:"value"     xml:"Value"`
	}

	tempValCurs struct {
		Valutes []Currency `xml:"Valute"`
	}
)

func ParseXMLFile(filename string) ([]Currency, error) {
	xmlData, err := readXML(filename)
	if err != nil {
		return nil, err
	}

	tempValCursData, err := decodeXMLContent(xmlData)
	if err != nil {
		return nil, err
	}

	return tempValCursData.Valutes, nil
}

func readXML(filename string) ([]byte, error) {
	xmlData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read XML file: %w", err)
	}

	return xmlData, nil
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

func (fc *FloatComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var strData string

	err := d.DecodeElement(&strData, &start)
	if err != nil {
		return fmt.Errorf("cannot decode XML element: %w", err)
	}

	strData = strings.Replace(strData, ",", ".", 1)

	val, err := strconv.ParseFloat(strData, 64)
	if err != nil {
		return fmt.Errorf("cannot parse float value '%s': %w", strData, err)
	}

	*fc = FloatComma(val)

	return nil
}
