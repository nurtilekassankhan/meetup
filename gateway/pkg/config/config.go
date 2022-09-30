package config

import (
	"github.com/spf13/viper"
)

const emptyString = ""

type Config struct {
	Server      Server            `path:"server" json:"server"`
	Accounting  AccountingServer  `path:"accounting_server" json:"accounting_server"`
	Transaction TransactionServer `path:"transaction_server" json:"transaction_server"`
}

func New(path string) (*Config, error) {

	conf := &Config{
		Server:      Server{},
		Accounting:  AccountingServer{},
		Transaction: TransactionServer{},
	}

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

	if err := conf.Accounting.Read(); err != nil {
		return nil, err
	}
	if err := conf.Accounting.Validate(); err != nil {
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
