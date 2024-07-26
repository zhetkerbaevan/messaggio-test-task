FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o /messaggio-test-task cmd/messaggio-test-task/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /messaggio-test-task .

EXPOSE 8080

CMD ["./messaggio-test-task"]
