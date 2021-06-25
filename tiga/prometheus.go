package tiga

import (
	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/prometheus"
	"github.com/liov/tiga/tiga/initialize"
)

var (
	metric *metrics.Metrics
)

func init() {
	sink, _ := prometheus.NewPrometheusSink()
	conf := metrics.DefaultConfig(initialize.InitConfig.Module)
	metric, _ = metrics.New(conf, sink)
	metric.EnableHostnameLabel = true
}
