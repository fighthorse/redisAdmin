FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
COPY .. .
RUN go mod download
COPY config/ /app/config
COPY assets/ /app/assets
RUN go build -ldflags="-s -w" -o /app/redisAdminExe main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/config/ /app/config/
COPY --from=builder /app/assets/ /app/assets/
COPY --from=builder /app/redisAdminExe /app/redisAdminExe

ENTRYPOINT ["./redisAdminExe", "-env"]
CMD ["/app/config/local"]