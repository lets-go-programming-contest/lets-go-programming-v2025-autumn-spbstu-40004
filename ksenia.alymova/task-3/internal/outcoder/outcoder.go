package outcoder

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Ksenia-rgb/task-3/internal/indecoder"
)

const filePerm int = 0666

func prepOutputPath(outputFile string) error {
	splitPath := strings.Split(outputFile, "/")

	pathLen := len(splitPath)
	if pathLen > 1 {
		err := os.Mkdir(splitPath[0], os.FileMode(filePerm))
		if err != nil && !os.IsExist(err) {
			return err
		}

		for i := 1; i < pathLen-1; i++ {
			splitPath[i] = splitPath[i-1] + "/" + splitPath[i]

			err := os.Mkdir(splitPath[i], os.FileMode(filePerm))
			if err != nil && !os.IsExist(err) {
				return err
			}
		}
	}

	return nil
}

func OutputProcess(outputFile string, inputData indecoder.BankData) error {
	err := prepOutputPath(outputFile)
	if err != nil {
		return err
	}

	outputByte, err := json.MarshalIndent(inputData.ValCurs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, outputByte, os.FileMode(filePerm))
}
