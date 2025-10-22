package indecoder

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html/charset"
)

func InputProcess(inputFile string) (BankData, error) {
	var inputData BankData

	inputReader, err := os.Open(inputFile)
	if err != nil {
		return inputData, fmt.Errorf("failed to open input file: %w", err)
	}

	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&inputData)
	if err != nil {
		return inputData, fmt.Errorf("failed to decode input file: %w", err)
	}

	for index := range inputData.ValCurs {
		err := inputData.ValCurs[index].convertFloatValue()
		if err != nil {
			return inputData, fmt.Errorf("failed to convert value from input file: %w", err)
		}
	}

	sort.Sort(inputData)

	return inputData, nil
}
