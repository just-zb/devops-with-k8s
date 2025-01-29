First create gcloud secret and configmap
`kubectl create secret generic gcloud-service-account \
  --from-file=service-account-key.json=./service-account-key.json`
`kubectl apply -f postgres-secret.yaml`
`kubectl apply -f job-env.yaml`
Then, apply backup-db job, `kubectl apply -f backupDB.yaml`