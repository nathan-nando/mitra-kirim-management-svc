FROM golang:1.22-alpine

WORKDIR /app

COPY .env .

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o management-svc .

CMD ["./management-svc"]
