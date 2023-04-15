FROM golang:1.20.2

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./ /app

RUN go mod tidy

CMD ["air"]
