1. Run `run.sh` in /db, to create postgres pod, it will decode configmap.enc.yaml using my own local sops.
2. Run `run.sh` to start the todo backend container, and visit http://todo-app.localtest.me:8081/ , the todos submitted by brower will be saved into database, also can be retreived from the database.

data inside configmap.yaml:
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
    name: postgres-cfg
    namespace: project
data:
    # kv format
    POSTGRES_PASSWORD: password
    POSTGRES_USER: myuser
    POSTGRES_DB: mydb
    POSTGRES_HOST: postgres-svc
```