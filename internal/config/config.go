package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	HTTPServer
	Storage
}

type HTTPServer struct {
	Address  string        `yaml:"address" env-default:":8080"`
	Timeout  time.Duration `yaml:"timeout" env-default:"5s"`
	TimeIdle time.Duration `yaml:"time_idle" env-default:"60s"`
}

type Storage struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
}

func MustLoad() Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("CONFIG_PATH does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return cfg
}
