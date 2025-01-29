#!/bin/zsh

set -e

DOCKER_IMAGE="pingpong"    # Docker 镜像名称
DOCKER_TAG="latest"        # Docker 镜像标签
OUTPUT_BINARY="main"       # 输出二进制文件名称

echo "Compiling $SOURCE_FILE..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $OUTPUT_BINARY .
echo "Compilation successful. Output binary: $OUTPUT_BINARY"
echo "Building Docker image $DOCKER_IMAGE:$DOCKER_TAG..."
#docker build -t "$DOCKER_IMAGE:$DOCKER_TAG" .
docker buildx build --platform linux/amd64 -t "$DOCKER_IMAGE:$DOCKER_TAG" --no-cache .
echo "Cleaning up..."
rm -f $OUTPUT_BINARY
echo "Docker image $DOCKER_IMAGE:$DOCKER_TAG built successfully."

PUSH_IMAGE="zhubao/pingpong:latest"
echo "push image $PUSH_IMAGE"
docker tag $DOCKER_IMAGE:$DOCKER_TAG $PUSH_IMAGE
docker push $PUSH_IMAGE
echo "push over"
#echo "running image $DOCKER_IMAGE:$DOCKER_TAG"
#docker run $DOCKER_IMAGE:$DOCKER_TAG

#echo "delete related pod"
#kubectl delete -f manifests/ || true
echo "run k8s cluster"
kubectl apply -f manifests/

sleep 5

echo "print kubectl log"
kubectl describe deployment/pingpong-dep
