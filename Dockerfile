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
RUN adduser -D goUser

COPY --from=base /go/goServer/main /
COPY --from=base /go/goServer/base/base.html /
COPY --from=base /go/goServer/tls /tls
RUN chmod +x /main && chown goUser:goUser /main && touch goServer.log && chown goUser:goUser goServer.log

USER goUser

WORKDIR /
ENTRYPOINT ["/main"]
