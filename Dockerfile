FROM golang:alpine AS builder
RUN apk update && \
    apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 go build -o tharkey-api
    
FROM scratch
COPY --from=builder /app/tharkey-api /
ENTRYPOINT [ "/tharkey-api" ]