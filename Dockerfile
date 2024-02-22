FROM golang:1.20.5 AS builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o tts main.go

FROM debian:stable-slim
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=builder /workspace/tts /app

CMD ["/app/tts"]
