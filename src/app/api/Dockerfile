FROM golang:1.11.2

WORKDIR /go/src/app/api/cmd/api
COPY . .

RUN go install -v ./...

CMD ["api"]