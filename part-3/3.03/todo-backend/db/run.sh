sops -d configmap.enc.yaml > configmap.yaml

echo "apply configmap and statefulset"
kubectl apply -f configmap.yaml

kubectl apply -f statefulset.yaml