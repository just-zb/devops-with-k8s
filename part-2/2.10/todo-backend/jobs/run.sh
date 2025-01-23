echo "build image"
docker build -t zhubao/read-hourly:latest .
echo "push image"
docker push zhubao/read-hourly:latest
echo "start cronjob"
kubectl delete -f cronJob.yaml || true
kubectl apply -f cronJob.yaml