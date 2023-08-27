package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server   *Server             `yaml:"server"`
	MySQL    *MySQL              `yaml:"mysql"`
	PgSQL    *PgSQL              `yaml:"pgsql"`
	Zap      *Zap                `yaml:"zap"`
	Redis    *Redis              `yaml:"redis"`
	Cos      *Cos                `yaml:"cos"`
	Etcd     *Etcd               `yaml:"etcd"`
	Services map[string]*Service `yaml:"services"`
	Domain   map[string]*Domain  `yaml:"domain"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
