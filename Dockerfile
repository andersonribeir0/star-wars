FROM golang:1.18.3-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o ms

FROM alpine:3.15.4
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/ms .
CMD ["./ms"]