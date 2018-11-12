#FROM golang:1.11.2-alpine3.8 AS build
##RUN apk --no-cache add gcc g++ make ca-certificates
##WORKDIR /go/src/github.com/lfordyce/hero_cqrs
##RUN go get -u golang.org/x/vgo
#WORKDIR $GOPATH/src/github.com/lfordyce/hero_cqrs
#
## Populate the module cache based on the go.{mod,sum} files.
##COPY go.mod .
##COPY go.sum .
##RUN vgo list -e $(vgo list -f '{{.Path}}' -m all)
#
#COPY util util
#COPY event event
#COPY db db
#COPY search search
#COPY schema schema
#COPY hero-services hero-services
#COPY pusher-service pusher-service
#COPY query-service query-service
#
##RUN go install ./...
##RUN go version && go get -u -v golang.org/x/vgo
##RUN vgo install ./...
#
#RUN vgo install -getmode=local ./...
#
#FROM alpine:3.8
#WORKDIR /user/bin
#COPY --from=build /go/bin .
#####


FROM golang:latest as build
#FROM golang:1.11.2-alpine3.8 AS build
RUN go version && go get -u -v golang.org/x/vgo
#WORKDIR $GOPATH/src/github.com/lfordyce/hero_cqrs
WORKDIR go/src/github.com/lfordyce/hero_cqrs

# Populate the module cache based on the go.{mod,sum} files.
#COPY go.mod go.mod
#COPY go.sum go.sum
#RUN vgo list -e $(vgo list -f '{{.Path}}' -m all)s

#COPY go.mod go.mod
#COPY go.sum go.sum

#COPY go.mod go.sum ./

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod ./
COPY go.sum ./
RUN vgo list -e $(vgo list -f '{{.Path}}' -m all)

COPY util util
COPY event event
COPY db db
COPY search search
COPY schema schema
COPY hero-service hero-service
COPY query-service query-service
COPY pusher-service pusher-service

#RUN vgo list -e $(vgo list -f '{{.Path}}' -m all)s

#RUN go version && go get -u -v golang.org/x/vgo
#RUN vgo install ./...
RUN vgo install ./...
#RUN vgo install -getmode=local ./...

FROM gcr.io/distroless/base
#FROM alpine:3.8
#COPY --from=build /go/bin/hero_cqrs /
WORKDIR /usr/bin
COPY --from=build /go/bin .