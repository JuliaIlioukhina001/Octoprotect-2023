FROM golang:1.21-alpine as builder
RUN apk update && apk add --no-cache git gcc musl-dev upx ca-certificates
WORKDIR /src/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o octoprotect .
RUN upx --lzma octoprotect

FROM alpine
WORKDIR /app/
COPY --from=builder /src/octoprotect /app/octoprotect
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app/octoprotect"]
ENTRYPOINT "/app/octoprotect"
