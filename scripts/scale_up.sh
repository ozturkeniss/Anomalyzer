#!/bin/bash
# Kullanım: ./scale_up.sh <deployment> <replica_sayisi>
DEPLOYMENT=${1:-myapp}
REPLICAS=${2:-3}
kubectl scale deployment "$DEPLOYMENT" --replicas="$REPLICAS" 