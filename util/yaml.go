package util

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func ReadYamlConfig(filename string, config interface{}) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	file, err := os.ReadFile(abs)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}
	return nil
}
