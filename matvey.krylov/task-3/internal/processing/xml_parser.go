package processing

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	types "github.com/15446-rus75/task-3/internal/currency_utils"
	"golang.org/x/net/html/charset"
)

func ParseCurrencyData(filePath string) (*types.CurrencyData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			_ = closeErr
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file content: %w", err)
	}

	contentWithDots := bytes.ReplaceAll(content, []byte(","), []byte("."))

	decoder := xml.NewDecoder(bytes.NewReader(contentWithDots))
	decoder.CharsetReader = charset.NewReaderLabel

	var data types.CurrencyData
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return &data, nil
}
