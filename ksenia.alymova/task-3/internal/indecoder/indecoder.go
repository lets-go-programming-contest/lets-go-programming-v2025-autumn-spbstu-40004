package indecoder

import (
	"encoding/xml"
	"os"
	"sort"

	"golang.org/x/net/html/charset"
)

func InputProcess(inputFile string) (BankData, error) {
	var inputData BankData

	inputReader, err := os.Open(inputFile)
	if err != nil {
		return inputData, err
	}

	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&inputData)
	if err != nil {
		return inputData, err
	}

	for index := range inputData.ValCurs {
		err := inputData.ValCurs[index].convertFloatValue()
		if err != nil {
			return inputData, err
		}
	}

	sort.Sort(inputData)

	return inputData, err
}
