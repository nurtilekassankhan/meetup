package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type TransactionServer struct {
	Port string `json:"port" path:"port"`
	Host string `json:"host" path:"host"`
}

func (s *TransactionServer) Read() error {
	s.Port = viper.GetString("transaction_server.port")
	s.Host = viper.GetString("transaction_server.host")
	return nil
}

func (s *TransactionServer) Validate() error {
	if strings.Compare(s.Port, emptyString) == 0 {
		errors.New("empty transaction server port")
	}
	if strings.Compare(s.Host, emptyString) == 0 {
		errors.New("empty transaction server host")
	}
	return nil
}
