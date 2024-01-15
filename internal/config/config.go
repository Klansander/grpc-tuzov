package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path not found")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path + fmt.Sprintf("%v", err.Error()))
	}
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic(err.Error())
	}
	return cfg
}

func fetchConfigPath() (res string) {

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("./../../config/local.yaml")
	}
	return
}
