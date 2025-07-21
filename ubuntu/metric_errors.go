package main

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_requests_errors_total",
			Help: "Toplam hata sayısı",
		},
	)
)
