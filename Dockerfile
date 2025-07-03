FROM golang:1.22-alpine

RUN apk update && apk add --no-cache sqlite-dev gcc musl-dev

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

