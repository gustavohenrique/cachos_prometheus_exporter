package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Config struct {
	Namespace string
	Shell     string
	Timeout   time.Duration
}

type PrometheusExporter struct {
	sync.RWMutex
	up         prometheus.Gauge
	varnishTop VarnishTop
	config     Config
}

func NewPrometheusExporter(config Config) *PrometheusExporter {
	return &PrometheusExporter{
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: config.Namespace,
			Name:      "up",
			Help:      "Was the last scrape of varnish successful.",
		}),
		varnishTop: VarnishTop{
			Binary: "varnishtop",
			Args:   []string{"-1"},
		},
		config: config,
	}
}

func (pe *PrometheusExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- pe.up.Desc()
}

func (pe *PrometheusExporter) Collect(ch chan<- prometheus.Metric) {
	pe.Lock()
	defer pe.Unlock()

	pe.up.Set(1)
	output, err := pe.collectFromVarnishTop()
	if err != nil {
		log.Println(err)
		pe.up.Set(0)
	}

	metrics, err := pe.varnishTop.Parse(output)
	if err != nil {
		log.Println(err)
	}
	for _, gauge := range metrics.Gauges {
		metricType := prometheus.CountValue
		name := pe.config.Namespace + "_" + gauge.Name
		desc := prometheus.NewDesc(name, gauge.Description, nil, nil)
		ch <- prometheus.MustNewConstMetric(desc, metricType, gauge.Value)
	}
	ch <- pe.up
}

func (pe *PrometheusExporter) collectFromVarnishTop() (string, error) {
	varnishTop := pe.varnishTop
	shell := pe.config.Shell
	command := varnishTop.Binary + " " + strings.Join(varnishTop.Args, " ")
	buf := &bytes.Buffer{}
	ctx, _ := context.WithTimeout(context.Background(), pe.config.Timeout)
	cmd := exec.CommandContext(ctx, shell, "-c", command)
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s failed with %s", shell, command, err)
	}
	return buf.String(), nil

	// case "c", "a":
	// metricType = prometheus.CounterValue
	// case "g":
	// metricType = prometheus.GaugeValue
	// default:
	// metricType = prometheus.GaugeValue
	// }
}
