package configfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ghodss/yaml"
)

type datacenterConfig struct {
	Hosts      []string `json:"hosts,omitempty" validate:"required"`
	DNSServer  string   `json:"dnsServer,omitempty" validate:"required"`
	DHCPServer string   `json:"dhcpServer,omitempty" validate:"required"`
}

func readConfigFile(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("Config filepath is empty")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error in reading config file: +%v", err)
	}
	return b, nil
}

func ParseDatacenterConfig(path string) (*datacenterConfig, error) {
	b, err := readConfigFile(path)
	if err != nil {
		return nil, err
	}

	config := datacenterConfig{}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, fmt.Errorf("Error decoding yaml file: +%v", err)
	}
	return &config, nil
}

func WriteDatacenterConfig(configDir string, config *datacenterConfig) error {
	b, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("Error writing configuration: +%v", err)
	}

	configPath := path.Join(configDir, "config.json")
	fmt.Printf("Config file written to: %s\n", configPath)
	err = ioutil.WriteFile(configPath, b, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to home directory: +%v", err)
	}
	return nil
}
