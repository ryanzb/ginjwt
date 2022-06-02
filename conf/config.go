package conf

import (
	"github.com/ryanzb/yaml"
)

type Config struct {
	Address    string `yaml:"address"`
	DBHost     string `yaml:"db_host"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

func Load(filename string) (*Config, error) {
	var config Config
	err := yaml.UnmarshalFile(filename, &config)
	return &config, err
}
