package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type AccountingServer struct {
	Port string `json:"port" path:"port"`
	Host string `json:"host" path:"host"`
}

func (s *AccountingServer) Read() error {
	s.Port = viper.GetString("accounting_server.port")
	s.Host = viper.GetString("accounting_server.host")
	return nil
}

func (s *AccountingServer) Validate() error {
	if strings.Compare(s.Port, emptyString) == 0 {
		errors.New("empty accounting server port")
	}
	if strings.Compare(s.Host, emptyString) == 0 {
		errors.New("empty accounting server host")
	}
	return nil
}
