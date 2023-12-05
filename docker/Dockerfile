FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go mod download

CMD ["air"]

