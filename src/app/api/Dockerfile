FROM golang:1.11.2

WORKDIR /go/src/app/api

# Copy over all files.
COPY . .

# Get dependencies.
RUN go get github.com/FiloSottile/gvt && \
    gvt restore

# Build the binary.
RUN go install -v ./...

CMD ["api"]