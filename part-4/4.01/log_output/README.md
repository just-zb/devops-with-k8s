run `run.sh`
Like pingpong, add readinessProbe to deployment.yaml, handle request to "/healthz" with logic : 
pingpong server ok : 200
pingpong server not ready: 500