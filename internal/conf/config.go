package conf

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Policies []Policy `yaml:"Policies"`
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

func GetFilter(policies []Policy, ApiKey string) (string, error) {
	var groups []Label
	var filter_label string
	for _, v := range policies {
		if v.APIKey == ApiKey {
			log.Print("Found")
			groups = append(groups, v.Labels...)
		}
	}

	if len(groups) == 0 {
		return "DeadEndpoint=\"You have no access to This, This returned empty thing\"", nil
	}

	for _, v := range groups {
		filter_label += v.Label
		if strings.Contains(v.Value, "*") {
			filter_label += "=~\""
		} else {
			filter_label += "=\""
		}

		filter_label += v.Value
		filter_label += "\","
	}
	return strings.Trim(filter_label, ","), nil
}
