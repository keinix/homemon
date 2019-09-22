package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LocalConfig struct {
	Scan   ScanConfig
	Server ServerConfig
	temp   string
}

type ScanConfig struct {
	Frequency int
	Scans     []string
}

type ServerConfig struct {
	Port string
	Ip   string
}

func (c *LocalConfig) SetFromFile(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, &c); err != nil {
		return err
	}
	return nil
}
