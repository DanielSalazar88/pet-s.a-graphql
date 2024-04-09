package settings

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v2"
)

//go:embed settings.yaml
var settingsFile []byte

type MySettings struct {
	Mysql    mysqlFields    `yaml:"mysql"`
	Handlers handlersFields `yaml:"handlers"`
}

func ReadYaml() *MySettings {
	var s MySettings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		log.Fatal(err)
	}

	return &s
}
