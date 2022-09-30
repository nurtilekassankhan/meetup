package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type Transaction struct {
	Port string `json:"port" path:"port"`
	Host string `json:"host" path:"host"`
}

func (s *Transaction) Read() error {
	return viper.UnmarshalKey("transaction_server", s)
}

func (s *Transaction) Validate() error {
	if strings.Compare(s.Port, emptyString) == 0 {
		return errors.New("transaction server port is empty")
	}
	if strings.Compare(s.Port, emptyString) == 0 {
		return errors.New("transaction server host is empty")
	}
	return nil
}
