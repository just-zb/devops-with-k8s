add readinessProbe and livenessProbe in deployment, 
add endpoint "/healthz" in project, check db connection.
when give wrong db url, readinessProbe will fail, and livenessProbe will restart the pod.