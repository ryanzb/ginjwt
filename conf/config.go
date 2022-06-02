package conf

import (
	"github.com/ryanzb/yaml"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Address string `yaml:"address"`
}

type DB struct {
	DSN string `yaml:"dsn"`
}

func Load(filename string) (*Config, error) {
	var config Config
	err := yaml.UnmarshalFile(filename, &config)
	return &config, err
}
