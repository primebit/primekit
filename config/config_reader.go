package config

import (
	yaml "gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Reader struct{}

func (cr Reader) ReadYaml(config interface{}, filename string) error {
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(yamlAbsPath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
