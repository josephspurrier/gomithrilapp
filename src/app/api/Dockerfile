FROM golang:1.14

WORKDIR /go/src/app/api

# Copy over all files.
COPY . .

# Get dependencies.
RUN go get github.com/FiloSottile/gvt && \
    gvt restore

# Build the binary.
RUN go install -v ./...

# Update and generate the swagger spec.
RUN sed -i'' -e 's/\:8081//' ./api.go && \
    go get github.com/go-swagger/go-swagger/cmd/swagger && \
    swagger generate spec -o /go/src/app/api/static/swagger/swagger.json

EXPOSE 8081
CMD ["api"]