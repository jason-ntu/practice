FROM golang:1.18-alpine

WORKDIR /Desktop/practice

COPY . .

RUN go build main.go

CMD [ "./main" ]