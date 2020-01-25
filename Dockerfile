FROM golang:1.13.6

COPY . /app
WORKDIR /app

RUN go build -o /bin/service github.com/knappjf/quickquestion

EXPOSE 8080

CMD ["/bin/service"]