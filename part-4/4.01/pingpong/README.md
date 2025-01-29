first, add readinessProbe in deployment.yaml,
then add logic code in main.go, to check the database status and response to readinessProbe.