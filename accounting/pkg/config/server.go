package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

const emptyString = ""

type Server struct {
	Port    string `json:"port" path:"port"`
	Network string `json:"network" path:"network"`
}

func (s *Server) Read() error {
	return viper.UnmarshalKey("server", s)
}

func (s *Server) Validate() error {
	if strings.Compare(s.Port, emptyString) == 0 {
		return errors.New("empty server port")
	}
	if strings.Compare(s.Network, emptyString) == 0 {
		return errors.New("empty network port")
	}
	return nil
}
