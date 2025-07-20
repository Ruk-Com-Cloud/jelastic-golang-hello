package config

import "fmt"

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func (s *ServerConfig) GetPort() string {
	return s.Port
}

func (s *ServerConfig) GetHost() string {
	return s.Host
}

func (s *ServerConfig) IsProduction() bool {
	return s.Host != "localhost" && s.Host != "127.0.0.1"
}
