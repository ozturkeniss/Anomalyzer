#!/bin/bash
# Kullanım: ./rate_limit.sh <configmap_yaml>
CONFIG=${1:-rate-limit-config.yaml}
kubectl apply -f "$CONFIG" 