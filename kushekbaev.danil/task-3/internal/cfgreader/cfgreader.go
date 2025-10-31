package cfgreader

import (
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

var errEmptyFlag = errors.New("config flag is empty")

type Cfg struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Parse() (Cfg, error) {
	var cfg Cfg

	cfgPath := flag.String("config", "", "input path config file")
	flag.Parse()

	if *cfgPath == "" {
		return cfg, errEmptyFlag
	}

	data, err := os.ReadFile(*cfgPath)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
