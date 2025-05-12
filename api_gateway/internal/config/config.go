package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
	"time"
)

type Config struct {
	HTTPServer   HTTPServer   `yaml:"server"`
	UserService  UserService  `yaml:"userService"`
	AdminService AdminService `yaml:"adminService"`
}

type HTTPServer struct {
	Port string `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
}

type UserService struct {
	Host    string        `yaml:"host" env:"USER_SERVICE_HOST" env-default:"localhost"`
	Port    string        `yaml:"port" env:"USER_SERVICE_PORT" env-default:"5555"`
	Timeout time.Duration `yaml:"timeout" env:"USER_SERVICE_TIMEOUT" env-default:"15s"`
}

type AdminService struct {
	Host    string        `yaml:"host" env:"ADMIN_SERVICE_HOST" env-default:"localhost"`
	Port    string        `yaml:"port" env:"ADMIN_SERVICE_PORT" env-default:"5556"`
	Timeout time.Duration `yaml:"timeout" env:"ADMIN_SERVICE_TIMEOUT" env-default:"15s"`
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
