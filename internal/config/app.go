package config

import (
	"ya-iot/pkg/iot"

	"github.com/spf13/viper"
)

type AppConfig struct {
	IOT     iot.Config `mapstructure:"iot"`
	Devices []string   `mapstructure:"devices"`
	Listen  string     `mapstructure:"listen"`
}

func NewAppConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("iot.host", "https://api.iot.yandex.net")
	viper.BindEnv("iot.token", "IOT_TOKEN")

	viper.SetDefault("listen", "localhost:8080")

	config := new(AppConfig)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
