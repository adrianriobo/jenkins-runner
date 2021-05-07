package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/adrianriobo/jkrunner/pkg/jkrunner"
	"github.com/adrianriobo/jkrunner/pkg/util"
	"gopkg.in/yaml.v2"
)

type Config struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func CreateConfig(url, username, password string) error {
	return saveConfig(&Config{
		URL:      url,
		Username: username,
		Password: password})
}

func LoadConfig() (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(getConfigFilePath())
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func saveConfig(config *Config) error {
	if data, err := yaml.Marshal(config); err == nil {
		if err = ioutil.WriteFile(getConfigFilePath(), data, 0644); err != nil {
			return err
		}
	}
	return nil
}

func getConfigFilePath() string {
	return filepath.Join(util.GetHomeDir(), jkrunner.Home, jkrunner.ConfigFileName)
}
