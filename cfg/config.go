package cfg

import (
	"encoding/json"
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type consul struct {
	Host    string `json:"host"    yaml:"host"`
	Service string `json:"service" yaml:"service"`
	ID      string `json:"id"      yaml:"id"`
}

type server struct {
	Host     string `json:"host"     yaml:"host"`
	Protocol string `json:"protocol" yaml:"protocol"`
	MaxConn  int    `json:"maxConn"  yaml:"maxConn"`
	Timeout  int    `json:"timeout"  yaml:"timeout"`
}

type Config struct {
	Server *server `json:"server" yaml:"server"`
	Consul *consul `json:"consul" yaml:"consul"`
}

func (this_ *Config) String() string {
	data, _ := json.Marshal(this_)
	return string(data)
}

func New(fname string) (*Config, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	this_ := &Config{}
	err = yaml.Unmarshal(data, this_)
	if err != nil {
		return nil, err
	}

	if len(this_.Server.Host) == 0 {
		return nil, errors.New("config.server.host is invalid")
	}

	if this_.Server.MaxConn <= 0 {
		return nil, errors.New("config.server.maxConn is invalid")
	}

	if this_.Server.Timeout <= 0 {
		return nil, errors.New("config.server.timeout is invalid")
	}

	return this_, nil
}
