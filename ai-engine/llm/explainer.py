import requests

def explain_with_llama(anomaly_type, metrics, details=None, host="http://ollama:11434", model="llama2:7b-chat"):
    solution_types = "[scale_up, rate_limit, restart_service, check_logs, alarm, no_action, investigate]"
    prompt = (
        f"Metric: {metrics.get('name', 'unknown')}\n"
        f"Recent values: {metrics.get('recent', [])}\n"
        f"Current value: {metrics.get('value', 'unknown')} at {metrics.get('timestamp', '')}\n"
        f"Detected anomaly: {anomaly_type}\n"
        "Please explain the anomaly and suggest a solution in natural language.\n"
        f"At the end, output the solution type as a single line, using ONLY ONE of these exact options:\n"
        f"{solution_types}\n"
        "Format:\n"
        "Solution type: <one_of_the_above>\n"
        "Do not use any other word or variation. Only use one of the above."
    )
    response = requests.post(
        f"{host}/api/generate",
        json={"model": model, "prompt": prompt, "stream": False}
    )
    response.raise_for_status()
    return response.json()["response"].strip()
