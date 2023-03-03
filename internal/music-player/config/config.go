package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	ListenHTTP struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"http_listen"`
	ListenGRPC struct {
		Network string `yaml:"network" env-default:"tcp"`
		Port    string `yaml:"port" env-default:"8080"`
	} `yaml:"grpc_listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
			panic(err)
		}
	})
	return instance
}
