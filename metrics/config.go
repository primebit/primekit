package metrics

type Configuration struct {
	Metrics ConfigurationMetrics `yaml:"metrics"`
}

type ConfigurationMetrics struct {
	Gauges     []string `yaml:"gauges"`
	Counters   []string `yaml:"counters"`
	Histograms []string `yaml:"histograms"`
}
