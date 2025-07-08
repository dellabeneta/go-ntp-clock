FROM golang:alpine3.21 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 appgroup && \
    adduser -D -u 1001 -G appgroup appuser

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/static ./static

RUN chown -R appuser:appgroup /root/

USER appuser

EXPOSE 8080

CMD ["./main"]