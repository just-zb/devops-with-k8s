send a http request to https://en.wikipedia.org/wiki/Special:Random, 
and the redirect url will be in the response's header, in tag `location`.
run `run.sh` to build the image and create the job.
to test, run kubectl create job --from=cronjob/read-hourly read-hourly-manual