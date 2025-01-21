#!/bin/zsh

set -e

DOCKER_IMAGE="devops-0.1"    # Docker 镜像名称
DOCKER_TAG="latest"        # Docker 镜像标签
OUTPUT_BINARY="main"       # 输出二进制文件名称
SOURCE_FILE="main.go"

echo "Compiling $SOURCE_FILE..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $OUTPUT_BINARY $SOURCE_FILE
echo "Compilation successful. Output binary: $OUTPUT_BINARY"
echo "Building Docker image $DOCKER_IMAGE:$DOCKER_TAG..."
docker build -t "$DOCKER_IMAGE:$DOCKER_TAG" .
echo "Cleaning up..."
rm -f $OUTPUT_BINARY
echo "Docker image $DOCKER_IMAGE:$DOCKER_TAG built successfully."

PUSH_IMAGE="zhubao/devops-0.1:latest"
echo "push image $PUSH_IMAGE"
docker tag devops-0.1:latest $PUSH_IMAGE
docker push $PUSH_IMAGE
echo "push over"
#echo "running image $DOCKER_IMAGE:$DOCKER_TAG"
#docker run $DOCKER_IMAGE:$DOCKER_TAG

echo "run k8s cluster"
kubectl create deployment hashgenerator-dep --image=zhubao/devops-0.1:latest

sleep 2

echo "print kubectl log"
kubectl logs deployment/hashgenerator-dep