FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/api/main.go

FROM alpine:latest

COPY --from=builder /app/main /main
COPY --from=builder /app/public /public

EXPOSE 8080

CMD [ "/main" ]

