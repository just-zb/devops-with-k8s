#!/bin/sh

url=$(wget -S -O /dev/null https://en.wikipedia.org/wiki/Special:Random 2>&1 | grep -i 'Location:' | sed -n 's/.*location: *//p')

if [ -n "$url" ]; then
    # 使用 wget 发送 POST 请求
    echo "url got"
    echo $url
    # connect todo-backend by service: todo-backend-svc:80
    wget --header="Content-Type: application/json" \
         --post-data="{\"todo\": \"$url\", \"done\": false}" \
         http://todo-backend-svc/todos
else
    echo "Error: Could not retrieve the URL."
fi