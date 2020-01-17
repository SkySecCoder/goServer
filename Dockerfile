# ----------------------------------------------------
# Building image
# ----------------------------------------------------

FROM golang:1.12.15-alpine3.11 AS base

COPY ./goServer/ /go/goServer
RUN apk add git && cd goServer && go mod init goServer 
RUN cd goServer && go build main.go

# ----------------------------------------------------
# Releasing image
# ----------------------------------------------------

FROM alpine:latest AS release
COPY --from=base /go/goServer/main /
COPY --from=base /go/goServer/base/base.html /
RUN chmod +x /main

WORKDIR /
ENTRYPOINT ["/main"]
