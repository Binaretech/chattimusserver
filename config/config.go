package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	data, _ := ioutil.ReadFile(envDir())

	viper.SetConfigType("env")

	if err := viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		logrus.Fatalf("Fatal error config file: %s \n", err)
	}

	viper.SetDefault("SERVER_PORT", 80)
}

func envDir() string {
	dir, _ := os.Getwd()
	return path.Join(dir, ".env")
}
