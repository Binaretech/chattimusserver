package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.AutomaticEnv()

	Config.SetDefault("SERVER_PORT", 80)
}
