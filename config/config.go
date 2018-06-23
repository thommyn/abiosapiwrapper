package config

import (
	"encoding/json"
	"io/ioutil"
)

type Target struct {
	Host string
	Converter string
}

type Config struct {
	IP string
	Port string
	Routes map[string]Target
	AllowedQueryParameters []string
	TimePerToken int64
	BurstTokens int64
}

func LoadDefaultConfig() *Config {
	routes := map[string]Target {
		"/series/live": Target{Host: "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", Converter: "none"},
		"/players/live": Target{Host: "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", Converter: "liveplayers"},
		"/teams/live": Target{Host: "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", Converter: "liveteams"},
	}

	conf := Config {
		IP: "localhost",
		Port: "8080",
		Routes: routes,
		AllowedQueryParameters: []string{"access_token", "page"},
		TimePerToken: 2000,
		BurstTokens: 5,
	}

	return &conf
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
