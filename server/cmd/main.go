package main

import (
	"anomalyzer-server/internal/eventstreamer"
	"anomalyzer-server/internal/metrics"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AISolution struct {
	AnomalyType  string `json:"anomaly_type"`
	SolutionType string `json:"solution_type"`
	Explanation  string `json:"explanation"`
}

func main() {
	fmt.Println("Anomalyzer Backend starting...")

	// AI-Engine'den gelen çözüm önerilerini dinle
	go listenToAISolutions()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/metrics-test", func(w http.ResponseWriter, r *http.Request) {
		data, err := metrics.QueryPrometheus("app_requests_total")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	http.HandleFunc("/send-test-metric", func(w http.ResponseWriter, r *http.Request) {
		publisher, err := eventstreamer.NewPublisher([]string{"kafka:9092"}, "metrics-topic")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to create publisher: " + err.Error()))
			return
		}
		defer publisher.Close()

		testMessage := map[string]interface{}{
			"name":      "app_requests_total",
			"values":    []int{100, 120, 300, 900, 10000},
			"timestamp": "2024-07-21T12:00:00Z",
		}

		messageBytes, _ := json.Marshal(testMessage)
		err = publisher.Publish(messageBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to publish message: " + err.Error()))
			return
		}

		w.Write([]byte("Test message sent to Kafka"))
	})

	// Aksiyon scriptleri için endpoint'ler
	http.HandleFunc("/action/scale-up", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing scale-up action...")
		// Kubernetes'te: kubectl scale deployment --replicas=3
		w.Write([]byte("Scale up action executed"))
	})

	http.HandleFunc("/action/scale-down", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing scale-down action...")
		// Kubernetes'te: kubectl scale deployment --replicas=1
		w.Write([]byte("Scale down action executed"))
	})

	http.HandleFunc("/action/restart-service", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing restart service action...")
		// Kubernetes'te: kubectl rollout restart deployment
		w.Write([]byte("Restart service action executed"))
	})

	http.HandleFunc("/action/alarm", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing alarm action...")
		// Slack/Discord webhook
		w.Write([]byte("Alarm action executed"))
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func listenToAISolutions() {
	consumer, err := eventstreamer.NewConsumer([]string{"kafka:9092"}, "ai-solutions-topic", "server-backend")
	if err != nil {
		log.Printf("Failed to create consumer: %v", err)
		return
	}
	defer consumer.Close()

	log.Println("Listening for AI solutions...")
	for {
		message, err := consumer.Consume()
		if err != nil {
			log.Printf("Error consuming message: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		var solution AISolution
		if err := json.Unmarshal(message, &solution); err != nil {
			log.Printf("Error parsing solution: %v", err)
			continue
		}

		log.Printf("Received AI solution: Anomaly=%s, Solution=%s", solution.AnomalyType, solution.SolutionType)
		log.Printf("Explanation: %s", solution.Explanation)

		// Çözüm önerisine göre aksiyon al
		executeAction(solution.SolutionType, solution.Explanation)
	}
}

func executeAction(solutionType string, explanation string) {
	switch solutionType {
	case "scale_up":
		log.Println("Executing scale_up action...")
		// Kubernetes'te: kubectl scale deployment --replicas=3
	case "scale_down":
		log.Println("Executing scale_down action...")
		// Kubernetes'te: kubectl scale deployment --replicas=1
	case "restart_service":
		log.Println("Executing restart_service action...")
		// Kubernetes'te: kubectl rollout restart deployment
	case "alarm":
		log.Println("Executing alarm action...")
		// Slack/Discord webhook
	case "rate_limit":
		log.Println("Executing rate_limit action...")
		// Service Mesh ile rate limiting
	case "check_logs":
		log.Println("Executing check_logs action...")
		// Kubernetes log aggregation
	case "investigate":
		log.Println("Executing investigate action...")
		// Detaylı analiz
	case "no_action":
		log.Println("No action needed...")
	default:
		log.Printf("Unknown solution type: %s", solutionType)
	}
}
