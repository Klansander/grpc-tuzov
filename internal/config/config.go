package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
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

const configPath = "./config/config.local.yaml"

var once sync.Once

func MustLoad() *Config {

	cfg := &Config{}

	path := fetchConfigPath()
	if path == "" {
		path = configPath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path + fmt.Sprintf("%v", err.Error()))
	}
	fmt.Println(path)
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	helpText := "Список переменных окружения"
	help, _ := cleanenv.GetDescription(cfg, &helpText)
	logrus.Info(help)

	return cfg
}

func fetchConfigPath() (res string) {

	once.Do(func() {
		flag.StringVar(&res, "config", configPath, "path to config file")
		flag.Parse()

	})
	return
}
