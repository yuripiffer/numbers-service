FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o numbers-service

# ---- Runtime image ----
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/numbers-service .
CMD ["./numbers-service"]