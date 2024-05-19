package main

import (
	"log"
	"net/http"
	"slices"
	"ya-iot/internal/collector"
	"ya-iot/internal/config"
	"ya-iot/pkg/iot"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	if len(config.Devices) == 0 {
		log.Fatal("No devices specified")
	}

	client := iot.NewClient(&config.IOT)

	iotInfo, err := client.GetIOTInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("IOT Info: %+v", iotInfo)

	var devices = make([]string, 0, len(iotInfo.Devices))
	for _, device := range iotInfo.Devices {
		if !slices.Contains(config.Devices, device.Id) {
			continue
		}
		log.Printf("Device: %+v", device)
		devices = append(devices, device.Id)
	}

	reg := prometheus.NewRegistry()

	provider := collector.NewCollector(reg, client, devices)
	reg.MustRegister(provider)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(config.Listen, nil))
}
