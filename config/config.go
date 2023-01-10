package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	User User `json:"user"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, %s", err)
	}
	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	viper.Set("user", cfg.User)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Error writing config file, %s", err)
	}
	return err
}
