execute `kubectl get -n project deploy -o yaml \
    | linkerd inject - \
    | kubectl apply -f -` to apply the injected deployment.
execute `kubectl get -n project deploy -o yaml \
    | linkerd inject - > deploy.yaml` to generate the injected deploy.yaml file.