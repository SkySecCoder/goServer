# ----------------------------------------------------
# Building image
# ----------------------------------------------------

FROM golang:1.12.15-alpine3.11 AS base

COPY ./goServer/ /go/goServer
RUN cd goServer && go mod init 
RUN go build main.go

# ----------------------------------------------------
# Releasing image
# ----------------------------------------------------

FROM alpine:latest AS release
COPY --from=base /go/main /
RUN chmod +x /main

WORKDIR /
ENTRYPOINT ["/main"]
