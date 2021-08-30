package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		ListenAddr string `yaml:"listenAddr", envconfig"LISTENADDR"`
		Port int `yaml:"port", envconfig:"PORT"`
	} `yaml:"server", envconfig:"SERVER"`
	Directories[] struct {
		FilePath string `yaml:"filepath"`
		UrlPath string `yaml:"urlpath"`
	} `yaml:"directories"`
}

func loadYamlConfig(config *Config) {
	yamlConfigFile, err := os.Open("config/files/config.yaml")
	if err != nil {
		panic("Unable to load config file: " + err.Error())
	}
	defer yamlConfigFile.Close()

	yamlDecoder := yaml.NewDecoder(yamlConfigFile)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		panic("Unable to decode config file: " + err.Error())
	}
}

func loadEnvConfig(config *Config) {
	err := envconfig.Process("TESTAPP", config)
	if err != nil {
		panic("Unable to parse env config !?!?!" + err.Error())
	}
}
func LoadConfig() Config {
	var config Config
	loadYamlConfig(&config)
	loadEnvConfig(&config)
	return config
}