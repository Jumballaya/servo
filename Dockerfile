# Builder
FROM golang as builder
WORKDIR /go/src/github.com/jumballaya/servo
COPY . .
RUN make build-linux

# Runner
FROM debian:stretch
WORKDIR /bin/
COPY --from=builder /go/src/github.com/jumballaya/servo/dist/servo_unix ./servo
CMD ["servo"]
