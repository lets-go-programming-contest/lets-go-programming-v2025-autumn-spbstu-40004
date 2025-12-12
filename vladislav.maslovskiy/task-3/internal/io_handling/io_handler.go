package iohandling

import (
	"fmt"
	"os"
	"strings"
)

func CreateDirectories(filePath string) error {
	pathComponents := strings.Split(filePath, "/")
	if len(pathComponents) == 1 {
		return nil
	}

	directoryPath := strings.Join(pathComponents[:len(pathComponents)-1], "/")

	err := os.MkdirAll(directoryPath, os.ModePerm)

	if err != nil {
		return fmt.Errorf("i/o folders: %w", err)
	} else {
		return nil
	}
}

func WriteDataToFile(filePath string, content []byte) error {
	err := CreateDirectories(filePath)
	if err != nil {
		return err
	}

	fileHandle, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("i/o: %w", err)
	}

	defer func() {
		if err := fileHandle.Close(); err != nil {
			panic(err.Error())
		}
	}()

	_, err = fileHandle.Write(content)
	if err != nil {
		return fmt.Errorf("i/o: %w", err)
	} else {
		return nil
	}
}
