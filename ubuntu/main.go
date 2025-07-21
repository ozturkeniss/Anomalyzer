package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestErrors)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		RequestCount.Inc()

		// Simüle edilen hata: %10 olasılıkla hata
		if rand.Float64() < 0.1 {
			RequestErrors.Inc()
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			RequestDuration.Observe(time.Since(start).Seconds())
			return
		}

		w.Write([]byte("Hello, World!"))
		RequestDuration.Observe(time.Since(start).Seconds())
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
