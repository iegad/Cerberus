package cfg

import (
	"encoding/json"
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var Instance *Config

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

func Init(fname string) error {
	data, err := os.ReadFile(fname)
	if err != nil {
		return err
	}

	Instance = &Config{}
	err = yaml.Unmarshal(data, Instance)
	if err != nil {
		return err
	}

	if len(Instance.Server.Host) == 0 {
		return errors.New("config.server.host is invalid")
	}

	if Instance.Server.MaxConn <= 0 {
		return errors.New("config.server.maxConn is invalid")
	}

	if Instance.Server.Timeout <= 0 {
		return errors.New("config.server.timeout is invalid")
	}

	return nil
}
