package config

import "time"

type Config struct {
	Period     time.Duration  `config:"period"`
	Timeout    time.Duration  `config:"timeout"`
	Privileged bool           `config:"privileged"`
	UseIPv4    bool           `config:"useipv4"`
	UseIPv6    bool           `config:"useipv6"`
	Targets    []TargetConfig `config:"targets" validate:"required"`
}

var DefaultConfig = Config{
	Period:     1 * time.Second,
	Timeout:    10 * time.Second,
	Privileged: true,
	UseIPv4:    true,
	UseIPv6:    true,
}

type TargetConfig struct {
	Name string   `config:"name" validate:"required"`
	Tags []string `config:"tags"`
	Desc string   `config:"desc"`
}
