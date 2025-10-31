package indecoder

import (
	"encoding/xml"
	"os"
	"sort"

	"golang.org/x/text/encoding/charmap"
)

func ProcessCurrencyFile(filePath string) (CurrencyCollection, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return CurrencyCollection{}, err
	}
	defer file.Close()

	var data CurrencyCollection
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charmap.Windows1251.NewDecoder

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

	sort.Sort(sort.Reverse(data))
	return data, nil
}
