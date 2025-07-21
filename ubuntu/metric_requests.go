package main

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Toplam istek sayısı",
		},
	)
)
