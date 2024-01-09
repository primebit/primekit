package config

import "github.com/primebit/primekit/metrics"

type (
	Configuration struct {
		//Application
		HTTP                  map[string]ConfigurationHTTPServer `yaml:"http"`
		metrics.Configuration `yaml:",inline"`
	}

	ConfigurationHTTPServer struct {
		Concurrency int    `yaml:"concurrency"`
		KeepAlive   bool   `yaml:"keep_alive"`
		Port        string `yaml:"port"`
	}
)
