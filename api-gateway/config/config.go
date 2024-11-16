package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSrvUrl string `mapstructure:"AUTH_SRV_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/envs")

	viper.SetConfigName("dev")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&c)
	return
}
