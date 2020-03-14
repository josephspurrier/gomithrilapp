FROM golang:1.14

WORKDIR /go/src/app/ingress

# Copy over all files.
COPY . .

# Build the binary.
RUN go install -v ./...

EXPOSE 80
CMD ["ingress"]