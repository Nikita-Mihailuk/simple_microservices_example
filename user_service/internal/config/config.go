package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	GRPC GRPCConfig `yaml:"grpc"`
	DB   DataBase   `yaml:"db"`
}

type GRPCConfig struct {
	Port int `yaml:"port" env:"GRPC_PORT" env-default:"5555"`
}

type DataBase struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
	Username string `yaml:"userName" env:"DATABASE_USERNAME" env-default:"postgres"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-default:"00000000"`
	Name     string `yaml:"name" env:"DATABASE_NAME" env-default:"simple_microservices_example"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		path := fetchConfigPath()
		instance = LoadConfigByPath(path)
	})

	return instance
}

func LoadConfigByPath(path string) *Config {
	var cfg Config

	if path != "" {
		err := cleanenv.ReadConfig(path, &cfg)
		if err != nil {
			panic(err)
		}
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
