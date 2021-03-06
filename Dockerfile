FROM golang:1.17.3-buster

WORKDIR $GOPATH/src/github.com/Graphity/surge

COPY . .

RUN ["go", "install", "-v", "./..."]

CMD ["surge"]