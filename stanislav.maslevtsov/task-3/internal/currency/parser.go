package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func Parse(path string) (*Currencies, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open xml currencies file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	var (
		curs    Currencies
		decoder = xml.NewDecoder(file)
	)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml currencies file: %w", err)
	}

	return &curs, nil
}
