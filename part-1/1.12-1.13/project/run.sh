#!/bin/zsh

set -e

DOCKER_IMAGE="devops_gin_0.1"    # Docker 镜像名称
DOCKER_TAG="latest"        # Docker 镜像标签
OUTPUT_BINARY="main"       # 输出二进制文件名称
PUSH_IMAGE="zhubao/devops_gin_0.1:latest"

echo "Compiling $SOURCE_FILE..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $OUTPUT_BINARY .
echo "Compilation successful. Output binary: $OUTPUT_BINARY"
echo "Building Docker image $DOCKER_IMAGE:$DOCKER_TAG..."
docker build -t "$DOCKER_IMAGE:$DOCKER_TAG" .
echo "Cleaning up..."
rm -f $OUTPUT_BINARY
echo "Docker image $DOCKER_IMAGE:$DOCKER_TAG built successfully."


echo "push image $PUSH_IMAGE"
docker tag $DOCKER_IMAGE:$DOCKER_TAG $PUSH_IMAGE
docker push $PUSH_IMAGE
echo "push over"
#echo "running image $DOCKER_IMAGE:$DOCKER_TAG"
#docker run $DOCKER_IMAGE:$DOCKER_TAG

echo "run k8s deployment,service,ingress"
kubectl apply -f manifests/
sleep 5

echo "print kubectl log"
kubectl logs deployment/gin-dep