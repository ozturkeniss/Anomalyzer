import json
from anomaly.detector import detect_spike
from llm.explainer import explain_with_llama
from kafka.consumer import kafka_consume
from kafka.producer import kafka_produce

def process_metric_message(msg):
    # Örnek: msg = '{"name": "app_requests_total", "values": [100, 120, 300, 900, 1500], "timestamp": "2024-07-21T12:00:00Z"}'
    data = json.loads(msg)
    recent = data.get("values", [])
    anomaly_type = "spike" if detect_spike(recent) else "normal"
    print(f"Detected anomaly type: {anomaly_type}")
    metrics = {
        "name": data.get("name", "unknown"),
        "recent": recent,
        "value": recent[-1] if recent else None,
        "timestamp": data.get("timestamp", "")
    }
    explanation = explain_with_llama(anomaly_type, metrics)
    # Solution type'ı parse et
    import re
    match = re.search(r"Solution type:\s*(\w+)", explanation)
    solution_type = match.group(1) if match else "unknown"
    result = {
        "anomaly_type": anomaly_type,
        "solution_type": solution_type,
        "explanation": explanation
    }
    return json.dumps(result)

def main():
    print("[AI-Engine] Kafka consumer starting...")
    for msg in kafka_consume("metrics-topic"):
        print("Received metric:", msg)
        result = process_metric_message(msg)
        print("Publishing result:", result)
        kafka_produce("ai-solutions-topic", result)

if __name__ == "__main__":
    main() 