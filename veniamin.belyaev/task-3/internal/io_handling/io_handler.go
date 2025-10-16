package iohandler

import (
	"fmt"
	"os"
	"strings"
)

func ResolveFolders(filename string) error {
	stringsSlice := strings.Split(filename, "/")
	if len(stringsSlice) == 1 {
		return nil
	}

	folderPath := strings.Join(stringsSlice[:len(stringsSlice)-1], "/")

	err := os.MkdirAll(folderPath, os.ModePerm)

	if err != nil {
		return fmt.Errorf("i/o folders: %w", err)
	} else {
		return nil
	}
}

func WriteStringToFile(filename string, data []byte) error {
	err := ResolveFolders(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("i/o: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("i/o: %w", err)
	} else {
		return nil
	}
}
