package ioutils

import (
	"fmt"
	"os"
	"strings"
)

func resolveFolders(filename string) error {
	stringsSlice := strings.Split(filename, "/")
	if len(stringsSlice) == 1 {
		return nil
	}

	folderPath := strings.Join(stringsSlice[:len(stringsSlice)-1], "/")

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("i/o folders: %w", err)
	}

	return nil
}

func WriteBytesToFile(filename string, data []byte) error {
	err := resolveFolders(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("i/o: %w", err)
	}

	return nil
}
