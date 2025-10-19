package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type (
	FloatComma float32

	Currency struct {
		NumCode  uint       `json:"num_code"  xml:"NumCode"`
		CharCode string     `json:"char_code" xml:"CharCode"`
		Value    FloatComma `json:"value"     xml:"Value"`
	}

	Currencies struct {
		Data []*Currency `xml:"Valute"`
	}
)

func (floatComma *FloatComma) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var rawToken string

	err := decoder.DecodeElement(&rawToken, &start)
	if err != nil {
		return fmt.Errorf("failed to decode element: %w", err)
	}

	rawToken = strings.Replace(rawToken, ",", ".", 1)

	floatToken, err := strconv.ParseFloat(rawToken, 32)
	if err != nil {
		return fmt.Errorf("failed to parse raw token to float: %w", err)
	}

	*floatComma = FloatComma(floatToken)

	return nil
}
