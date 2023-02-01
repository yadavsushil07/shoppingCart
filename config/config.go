package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_HOST    string `mapstructure:"DB_HOST"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	DB_PORT    string `mapstructure:"DB_PORT"`
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASS"`
	DB_DRIVER  string `mapstructure:"DB_DRIVER"`
	DB_CHARSET string `mapstructure:"DB_CHARSET"`
	PORT       int    `mapstructure:"PORT"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	// fmt.Println(env)
	return &env
}
