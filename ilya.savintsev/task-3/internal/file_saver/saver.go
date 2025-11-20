package filesaver

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	ownerReadWrite = 0o600
	allReadWrite   = 0o755
)

var (
	errDirSave = errors.New("unable create directory")
	errWrtSave = errors.New("unable save file")
)

func SaveToFile(data []byte, outputFile string) error {
	dir := filepath.Dir(outputFile)

	if err := os.MkdirAll(dir, allReadWrite); err != nil {
		return errDirSave
	}

	err := os.WriteFile(outputFile, data, ownerReadWrite)
	if err != nil {
		return errWrtSave
	}

	return nil
}
