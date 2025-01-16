# build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /backend ./cmd/backend

# copy
FROM alpine:latest

WORKDIR /app

COPY --from=builder /backend .

EXPOSE 8080

CMD [ "./backend" ]
