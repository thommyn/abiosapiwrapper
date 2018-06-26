package config

import (
	"encoding/json"
	"io/ioutil"
)

type Target struct {
	Host string
	JQuery string
}

type Config struct {
	IP string
	Port string
	Routes map[string]Target
	AllowedQueryParameters []string
	TimePerRequest int64
	BurstRequests int64
}

func LoadConfig(path string) (*Config, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
