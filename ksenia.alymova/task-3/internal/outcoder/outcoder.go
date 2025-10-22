package outcoder

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Ksenia-rgb/task-3/internal/indecoder"
)

const filePerm = 0666

var ()

func prepOutputPath(outputFile string) error {
	splitPath := strings.Split(outputFile, "/")

	pathLen := len(splitPath)
	if pathLen > 1 {
		err := os.Mkdir(splitPath[0], os.FileMode(filePerm))
		if err != nil && !os.IsExist(err) {
			return fmt.Errorf("failed to create directory for output: %w", err)
		}

		for i := 1; i < pathLen-1; i++ {
			splitPath[i] = splitPath[i-1] + "/" + splitPath[i]

			err := os.Mkdir(splitPath[i], os.FileMode(filePerm))
			if err != nil && !os.IsExist(err) {
				return fmt.Errorf("failed to create directory for output: %w", err)
			}
		}
	}

	return nil
}

func OutputProcess(outputFile string, inputData indecoder.BankData) error {
	err := prepOutputPath(outputFile)
	if err != nil {
		return fmt.Errorf("failed to prepare path for output file: %w", err)
	}

	outputByte, err := json.MarshalIndent(inputData.ValCurs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to code output file: %w", err)
	}

	err = os.WriteFile(outputFile, outputByte, os.FileMode(filePerm))
	if err != nil {
		return fmt.Errorf("failed to wrile data in output file: %w", err)
	}

	return nil
}
