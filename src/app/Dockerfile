FROM node:12.16.1
# Create app directory.
WORKDIR /usr/src/ui
# Copy both package.json and package-lock.json.
COPY ./ui/package*.json ./
# Reliable DNS.
RUN echo nameserver 8.8.8.8 > /etc/resolv.conf
# Copy over all files.
COPY ./ui .
# Get dependencies.
RUN npm install
# Build the dist.
RUN npm run build

FROM golang:1.14
# Create app directory.
WORKDIR /go/src/app/api
# Copy over all files.
COPY ./api .
# Reliable DNS.
RUN echo nameserver 8.8.8.8 > /etc/resolv.conf
# Get dependencies.
RUN go get github.com/FiloSottile/gvt && \
    gvt restore
# Update and generate the swagger spec.
RUN sed -i'' -e 's/\:8080//' ./api.go && \
    go get github.com/go-swagger/go-swagger/cmd/swagger && \
    swagger generate spec -o /go/src/app/swagger.json
# Build the binary.
RUN go install -v ./...

FROM golang:1.14
WORKDIR /go/src/app/webserver
# Copy over all files.
COPY --from=0 /usr/src/ui/dist /go/src/app/webserver/
COPY --from=1 /go/bin/api /go/bin
COPY --from=1 /go/src/app/swagger.json /go/src/app/webserver/static/
EXPOSE 80
CMD ["api"]