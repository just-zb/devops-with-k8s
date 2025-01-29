#!/bin/zsh

set -e

OUTPUT_BINARY="main"       # 输出二进制文件名称

echo "Compiling $SOURCE_FILE..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $OUTPUT_BINARY .
echo "Compilation successful. Output binary: $OUTPUT_BINARY"
PUSH_IMAGE="zhubao/log-output:latest"
docker buildx build --platform linux/amd64,linux/arm64 -t $PUSH_IMAGE --push .
echo "Cleaning up..."
rm -f $OUTPUT_BINARY

echo "run k8s cluster"
kubectl delete -f manifests/ || true
kubectl apply -f manifests/
sleep 5

echo "print kubectl log"
kubectl describe deployment/log-output-dep
