package processor

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/cherkasoov/task-3/internal/config"
	"github.com/cherkasoov/task-3/internal/models"
)

func Run(cfg *config.Config) error {

	xmlData, err := readXMLData(cfg.InputFile)
	if err != nil {

		return fmt.Errorf("error reading xml file: %w", err)
	}

	return nil
}

func readXMLData(path string) (*models.XMLValCurs, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var valCurs models.XMLValCurs

	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, err
	}

	return &valCurs, nil
}
