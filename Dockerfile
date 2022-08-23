FROM golang:1.16-alpine

WORKDIR /cake-store

COPY . .

RUN go build -o cake-store

EXPOSE 8080

CMD ./cake-store