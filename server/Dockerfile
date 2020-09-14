FROM golang:1.14-alpine

WORKDIR /app

EXPOSE 80

COPY go.mod .
COPY go.sum .

COPY . .

CMD [ "go", "run", "main.go" ]