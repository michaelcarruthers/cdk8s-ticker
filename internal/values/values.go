package values

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ValuesConfig struct {
	ApiKey    string `yaml:"api_key"`
	Env       string `yaml:"environment"`
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	NumOfDays string `yaml:"number_of_days"`
	Image     string `yaml:"image"`
	Replicas  int64  `yaml:"replicas,omitempty"`
	Symbol    string `yaml:"symbol"`

	Service struct {
		Name string  `yaml:"name"`
		Port float64 `yaml:"port"`
		Type string  `yaml:"type"`
	} `yaml:"service"`
}

// Filename returns the values configuration file name
func Filename() (*string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("%s/values.yml", path)

	return &filename, nil
}

// Read reads and returns the content of the values configuration file
func Read() ValuesConfig {
	filename, err := Filename()
	if err != nil {
		log.Fatalf("error getting filename: %v", err)
	}

	data, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var content ValuesConfig
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		log.Fatalf("error unmarshalling data: %v", err)
	}

	return content
}
