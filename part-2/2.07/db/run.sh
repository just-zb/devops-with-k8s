sops -d configmap.enc.yaml > configmap.yaml

echo "apply configmap and statefulset"
kubectl apply -f configmap.yaml

kubectl delete statefulset/postgres-stset

kubectl apply -f statefulset.yaml