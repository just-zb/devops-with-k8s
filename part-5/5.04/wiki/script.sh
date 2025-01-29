#!/bin/bash
output_dir="/usr/share/nginx/html"
num=1
while true; do
    sleep_time=$((RANDOM % 301 + 600))
    output_file="$output_dir/random-$num.html"
    ((num+=1))
    curl -sL "https://en.wikipedia.org/wiki/Special:Random" -o $output_file
    sleep "$sleep_time"
done