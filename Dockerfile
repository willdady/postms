FROM golang:1.11.3 as builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install -v -a ./...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/ .
CMD ["./postms"]
