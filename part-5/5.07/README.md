deploy pingpong to Knative.
## 1. deploy postgres database
create configmap `kubectl apply -f db/statefulset.yaml`
create postgres statefulset and service `kubectl apply -f db/statefulset.yaml`
## 2. convert pingpong project to stateless Knative service
change main.go.
build image and push: run `run.sh`
## deploy the project
`kubectl apply -f manifests/Knative.yaml`