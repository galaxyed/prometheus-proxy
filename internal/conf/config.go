package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Policies []Label `yaml:"Policies"`
}
type Policy struct {
	Name   string  `yaml:"Name"`
	APIKey string  `yaml:"APIKey"`
	Labels []Label `yaml:"Labels"`
}
type Label struct {
	Label string `yaml:"Label"`
	Value string `yaml:"Value"`
}


// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
    // Create config structure
    config := &Config{}

    // Open config file
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Init new YAML decode
    d := yaml.NewDecoder(file)

    // Start YAML decoding from file
    if err := d.Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}