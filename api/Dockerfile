FROM golang:1.11.2

WORKDIR /go/src/github.com/josephspurrier/govueapp/api
COPY . .

RUN go install -v ./...

CMD ["api"]