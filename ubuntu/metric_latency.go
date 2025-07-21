package main

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP isteklerinin yanıt süresi (saniye)",
			Buckets: prometheus.DefBuckets,
		},
	)
)
