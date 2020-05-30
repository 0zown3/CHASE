FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd/chase
RUN go build -o main .
CMD ["./main"]
