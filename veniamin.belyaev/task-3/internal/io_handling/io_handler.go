package iohandler

import (
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

	return err
}

func WriteStringToFile(filename string, data []byte) error {
	err := ResolveFolders(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	_, err = file.Write(data)

	return err
}
