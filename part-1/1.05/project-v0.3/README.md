After running run.sh,
Run `kubectl get pods`
->
```
kubectl get po
NAME                       READY   STATUS    RESTARTS   AGE
gin-dep-7644558c4f-mhqh7   1/1     Running   0          57s
```
Run `kubectl port-forward gin-dep-7644558c4f-mhqh7 9999:9999`
Visit http://localhost:9999 , we can see the html page.