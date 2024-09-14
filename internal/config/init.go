package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	return viper.ReadInConfig()
}
