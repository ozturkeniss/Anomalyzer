import requests
import pandas as pd
from datetime import datetime, timedelta
import numpy as np
import pandas as pd
from scipy.stats import zscore

def load_prometheus_history(prometheus_url, metric, hours=24):
    end = datetime.utcnow()
    start = end - timedelta(hours=hours)
    params = {
        "query": metric,
        "start": start.isoformat() + "Z",
        "end": end.isoformat() + "Z",
        "step": "60"
    }
    url = f"{prometheus_url}/api/v1/query_range"
    resp = requests.get(url, params=params)
    resp.raise_for_status()
    data = resp.json()["data"]["result"]
    if not data:
        return pd.DataFrame()
    values = data[0]["values"]
    df = pd.DataFrame(values, columns=["timestamp", "value"])
    df["timestamp"] = pd.to_datetime(df["timestamp"], unit="s")
    df["value"] = pd.to_numeric(df["value"])
    return df

def detect_spike(series, threshold=1.5):
    z = zscore(series)
    if len(z) == 0:
        return False
    return z[-1] > threshold

def detect_drop(series, threshold=-3):
    z = zscore(series)
    if len(z) == 0:
        return False
    return z[-1] < threshold

def detect_latency_anomaly(latencies, quantile=0.99, threshold=1.5):
    q = np.quantile(latencies, quantile)
    return latencies[-1] > q * threshold

def detect_error_rate_anomaly(errors, totals, threshold=0.05):
    if len(errors) == 0 or len(totals) == 0:
        return False
    rate = errors[-1] / max(totals[-1], 1)
    return rate > threshold

def detect_multivariate_anomaly(df, contamination=0.01):
    try:
        from sklearn.ensemble import IsolationForest
    except ImportError:
        return False
    if len(df) < 10:
        return False
    model = IsolationForest(contamination=contamination)
    preds = model.fit_predict(df)
    return preds[-1] == -1 