#!/bin/bash
# Kullanım: ./restart_service.sh <deployment>
DEPLOYMENT=${1:-myapp}
kubectl rollout restart deployment "$DEPLOYMENT" 