package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Port
	Port int
	// DataDir
	DataDir string
	// ListenIp
	ListenIp string
}

func NewConfig(port int, dataDir string, listenIp string) *Config {
	return &Config{
		Port:     port,
		DataDir:  dataDir,
		ListenIp: listenIp,
	}
}

func (c *Config) Sync() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	file, err := os.Create("config.yml")
	if err != nil {
		return err
	}
	defer func(File *os.File) {
		_ = File.Close()
	}(file)
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func TryLoadConfig() (*Config, bool) {
	file, err := os.Open("config.yml")
	if err != nil {
		return nil, false
	}
	defer func(File *os.File) {
		_ = File.Close()
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, false
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, false
	}
	return &config, true
}
