package parser

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"

	"github.com/shycoshy/task-3/internal/domain"
)

func Parse(path string) (*domain.ValCurs, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data domain.ValCurs
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
