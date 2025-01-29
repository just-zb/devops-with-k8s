FROM alpine:latest
LABEL authors="michael"

RUN apk add --no-cache bash curl

COPY script.sh .
RUN chmod +x script.sh

ENTRYPOINT ["bash","script.sh"]