package ioutils

import (
	"fmt"
	"os"
	"strings"
)

const (
	filePerm = 0o0600
)

func resolveFolders(filename string) error {
	stringsSlice := strings.Split(filename, "/")
	if len(stringsSlice) == 1 {
		return nil
	}

	folderPath := strings.Join(stringsSlice[:len(stringsSlice)-1], "/")

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create dir %w", err)
	}

	return nil
}

func WriteBytesToFile(filename string, data []byte) error {
	err := resolveFolders(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, filePerm)
	if err != nil {
		return fmt.Errorf("failed to write file %w", err)
	}

	return nil
}
