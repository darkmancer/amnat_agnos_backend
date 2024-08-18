FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/server .
COPY --from=build /app/.env .
EXPOSE 8081
CMD ["./server"]