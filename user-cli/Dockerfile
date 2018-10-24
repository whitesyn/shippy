FROM golang:1.11.1 as builder
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/whitesyn/shippy/user-cli
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/whitesyn/shippy/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]