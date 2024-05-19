package collector

import (
	"fmt"
	"log"
	"strings"
	"ya-iot/pkg/iot"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	client  *iot.Client
	reg     *prometheus.Registry
	devices []string
	metrics map[string]*prometheus.Desc
}

func NewCollector(reg *prometheus.Registry, client *iot.Client, devices []string) *Collector {
	return &Collector{
		client:  client,
		reg:     reg,
		devices: devices,
		metrics: make(map[string]*prometheus.Desc),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	for _, device := range c.devices {
		info, err := c.client.GetDeviceInfo(device)
		if err != nil {
			log.Printf("Error occurred while getting device info: %s", err)
			continue
		}
		for _, metric := range info.Properties {
			if metric.State == nil {
				continue
			}
			var gauge *prometheus.Desc
			key := fmt.Sprintf("%s:%s", device, metric.State.Instance)
			name := fmt.Sprintf("iot:%s:%s", strings.ReplaceAll(info.Type, ".", "_"), strings.ReplaceAll(metric.State.Instance, ".", "_"))
			if gauge, ok := c.metrics[key]; !ok {
				gauge = prometheus.NewDesc(name, "", nil, prometheus.Labels{"device": info.Name, "type": info.Type})
				c.metrics[key] = gauge
			}
			gauge = c.metrics[key]
			log.Printf("Setting gauge %s to %f (%s)", name, metric.State.Value, key)
			ch <- prometheus.MustNewConstMetric(
				gauge,
				prometheus.GaugeValue,
				metric.State.Value,
			)
		}
	}
}
