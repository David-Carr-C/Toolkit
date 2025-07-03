package interfaces

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string, out interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(out); err != nil {
		return err
	}
	return nil
}
