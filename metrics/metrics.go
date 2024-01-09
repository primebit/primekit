package metrics

import (
	"github.com/primebit/primekit/config"
	"github.com/rcrowley/go-metrics"
	"io"
	"runtime"
	"time"
)

var (
	gauges     = make(map[string]metrics.Gauge)
	counters   = make(map[string]metrics.Counter)
	histograms = make(map[string]metrics.Histogram)
	m          runtime.MemStats
)

func Init() {
	// Register system metrics
	RegisterGauge("goroutins_count")
	RegisterGauge("mem_alloc")
	go exportCommonData()

	// Register from config
	for _, gauge := range config.Config.Metrics.Gauges {
		RegisterGauge(gauge)
	}

	for _, counter := range config.Config.Metrics.Counters {
		RegisterCounter(counter)
	}

	for _, histogram := range config.Config.Metrics.Histograms {
		RegisterHistogram(histogram)
	}
}

func RegisterGauge(name string) {
	if _, found := gauges[name]; found {
		return
	}
	gauges[name] = metrics.NewGauge()
	metrics.Register(name, gauges[name])
}

func RegisterCounter(name string) {
	if _, found := counters[name]; found {
		return
	}
	counters[name] = metrics.NewCounter()
	metrics.Register(name, counters[name])
}

func RegisterHistogram(name string) {
	if _, found := histograms[name]; found {
		return
	}
	histograms[name] = metrics.NewHistogram(metrics.NewUniformSample(1028))
	metrics.Register(name, histograms[name])
}

func UpdateGauge(name string, value int64) {
	if gauge, found := gauges[name]; found {
		gauge.Update(value)
	}
}

func UpdateHistogram(name string, value int64) {
	if histogram, found := histograms[name]; found {
		histogram.Update(value)
	}
}

func IncCounter(name string, value int64) {
	if counter, found := counters[name]; found {
		counter.Inc(value)
	}
}

func Export(writer io.Writer) {
	metrics.WriteJSONOnce(metrics.DefaultRegistry, writer)
	Clear()
}

func Clear() {
	for _, counter := range counters {
		counter.Clear()
	}
	for _, histogram := range histograms {
		histogram.Clear()
	}
}

func exportCommonData() {
	for range time.Tick(time.Duration(1) * time.Second) {
		UpdateGauge("routins_count", int64(runtime.NumGoroutine()))

		runtime.ReadMemStats(&m)
		UpdateGauge("mem_alloc", int64(bToMb(m.Alloc)))
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
