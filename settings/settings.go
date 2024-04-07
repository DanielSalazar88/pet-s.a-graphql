package settings

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	errReadYaml = "error leyendo el archivo yaml"
)

type MySettings struct {
	MySQL mySQL `yaml:"mysql"`
}

func New() (*MySettings, error) {
	var s MySettings

	y, err := readYaml()
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(y, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func readYaml() ([]byte, error) {
	content, err := os.ReadFile("settings.yaml")
	if err != nil {
		return nil, fmt.Errorf("%s:%s", errReadYaml, err)
	}

	return content, nil
}
