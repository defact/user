FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/defact/user
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/defact/user
ADD . /go/src/github.com/defact/user
RUN rm -rf /go/src/github.com/defact/user/vendor
RUN dep ensure
RUN GOOS=linux CGO_ENABLED=0 go build

FROM alpine:latest
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /go/src/github.com/defact/user/user /app
COPY --from=builder /go/src/github.com/defact/user/config/*.json /app/config/
COPY --from=builder /go/src/github.com/defact/user/migrations/*.sql /app/migrations/
ENTRYPOINT ./user