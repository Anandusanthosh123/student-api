package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" ` // struct tags in go - mentioning env from /config/local.yaml,from ENV named variable
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// this is an important function it must sucessfully execute when called if error app will not start
func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path  is not set")
		}
	}
	// check file available on that path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist:%s", configPath)
	}
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file:%s", err.Error())
	}
	return &cfg
}
