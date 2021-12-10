FROM golang:1.17.3 AS Builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.14
COPY --from=Builder /app .
CMD ["./main"]
