package config

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

type Server struct {
	Port string `json:"port" path:"port"`
}

func (s *Server) Read() error {
	s.Port = viper.GetString("server.port")
	return nil
}

func (s *Server) Validate() error {
	if strings.Compare(s.Port, emptyString) == 0 {
		return errors.New("empty server port")
	}
	return nil
}
