FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .


RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ADD . .

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
