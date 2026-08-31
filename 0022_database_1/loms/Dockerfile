FROM golang:1.25 AS builder

WORKDIR /app

COPY go.work ./
COPY loms/go.mod loms/go.sum ./loms/
COPY pkg/go.mod pkg/go.sum ./pkg/

COPY . .

WORKDIR /app/loms
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/loms ./cmd/app

FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /bin/loms /app/loms

CMD ["./loms"]
