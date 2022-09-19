package conf

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Policies      []Policy `yaml:"Policies"`
	PrometheusKey string   `yaml:"PrometheusKey"`
	Authorization string   `yaml:"Authorization"`
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
		return "DeclinedLabel=\"You have no access to This, This returned empty thing\"", nil
	}

	groupsMap := make(map[string]string)
	for _, v := range groups {
		if _, ok := groupsMap[v.Label]; ok {
			groupsMap[v.Label] += "|"
			groupsMap[v.Label] += v.Value
		} else {
			groupsMap[v.Label] = v.Value
		}

	}
	log.Print(groupsMap)

	for label, value := range groupsMap {
		filter_label += label
		if strings.Contains(value, "*") || strings.Contains(value, "|") {
			filter_label += "=~\""
		} else {
			filter_label += "=\""
		}

		filter_label += value
		filter_label += "\","
	}

	return strings.Trim(filter_label, ","), nil
}
