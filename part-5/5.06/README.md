run `k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2 --k3s-arg "--disable=traefik@server:0"`

run `kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-crds.yaml`

run `kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-core.yaml`
run `kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.17.0/kourier.yaml`
run `kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'`

run `kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.17.0/serving-default-domain.yaml`
## Deploying a Knative Service
after install kn cli.
run `kn service create hello \
--image ghcr.io/knative/helloworld-go:latest \
--port 8080 \
--env TARGET=World`

run `kn service update hello \
--env TARGET=Knative`

![alt text](<CleanShot 2025-01-29 at 5 .04.48@2x.png>)
![alt text](<CleanShot 2025-01-29 at 5 .04.15@2x.png>)

## test Traffic splitting
run `kn service update hello \
--traffic hello-00001=50 \
--traffic @latest=50`
![alt text](<CleanShot 2025-01-29 at 5 .06.53@2x.png>)

## Autoscaling
![alt text](<CleanShot 2025-01-29 at 5 .09.24@2x.png>)
![alt text](<CleanShot 2025-01-29 at 5 .10.09@2x.png>)