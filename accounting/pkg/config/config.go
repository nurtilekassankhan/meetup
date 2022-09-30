package config

import "github.com/spf13/viper"

type Config struct {
	Server      Server      `path:"server" json:"server"`
	Transaction Transaction `path:"transaction" json:"transaction"`
}

func New(path string) (*Config, error) {

	conf := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := conf.Server.Read(); err != nil {
		return nil, err
	}
	if err := conf.Server.Validate(); err != nil {
		return nil, err
	}

	if err := conf.Transaction.Read(); err != nil {
		return nil, err
	}
	if err := conf.Transaction.Validate(); err != nil {
		return nil, err
	}

	return conf, nil
}
