package indecoder

import (
	"encoding/xml"
	"io"
	"os"
	"sort"

	"golang.org/x/text/encoding/charmap"
)

func ProcessCurrencyFile(filePath string) (CurrencyCollection, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return CurrencyCollection{}, err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic(closeErr.Error())
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var data CurrencyCollection
	err = decoder.Decode(&data)
	if err != nil {
		return CurrencyCollection{}, err
	}

	for i := range data.Items {
		err = data.Items[i].TransformValue()
		if err != nil {
			return CurrencyCollection{}, err
		}
	}

	sort.Sort(data)

	return data, nil
}
