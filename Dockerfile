FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

copy . .

RUN go build -o management-svc .

CMD ["./management-svc"]
