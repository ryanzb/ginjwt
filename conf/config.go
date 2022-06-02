package conf

import (
	"log"

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

func Load(filename string) *Config {
	var config Config
	if err := yaml.UnmarshalFile(filename, &config); err != nil {
		log.Fatalf("load config failed: %v", err)
		return nil
	}
	log.Printf("loaded config: %+v", config)
	return &config
}
