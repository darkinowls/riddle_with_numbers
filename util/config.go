package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDriver          string        `mapstructure:"DB_DRIVER"`
	DBSource          string        `mapstructure:"DB_SOURCE"`
	ServerAddress     string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string        `mapstructure:"TOKEN_SECRET_KEY"`
	TokenDuration     time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("default")
	viper.SetConfigType("env")
	viper.AutomaticEnv() // to be able to rewrite
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
