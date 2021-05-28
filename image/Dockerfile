FROM golang:1.14-alpine3.12

LABEL description="p76081158/ue-requests-generator"

RUN apk update && \
    apk add --no-cache curl git && \
    git clone https://github.com/p76081158/ue-requests-generator.git && \
    cd ue-requests-generator && \
    go build && \
    rm -rf /var/cache/apk/*

WORKDIR /go/ue-requests-generator

# Default command
ENTRYPOINT [ "./ue-requests-generator" ]