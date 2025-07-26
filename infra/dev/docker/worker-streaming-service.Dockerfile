FROM alpine
WORKDIR /app

ADD shared shared
ADD build build

ENTRYPOINT build/worker-streaming-service