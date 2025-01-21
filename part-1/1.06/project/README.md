after run`k3d cluster delete` `k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2`

Run `run.sh`,visit http://localhost:8082, shows the html page.