FROM golang:1.17.3-alpine3.14

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN apk --no-cache add curl

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
