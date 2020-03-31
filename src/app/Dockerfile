FROM node:12.16.1
# Create app directory
WORKDIR /usr/src/ui
# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY ./ui/package*.json ./
RUN echo nameserver 8.8.8.8 > /etc/resolv.conf
RUN npm install
COPY ./ui .
RUN npm run build

FROM golang:1.14
WORKDIR /go/src/app/api
# Copy over all files.
COPY ./api .
RUN echo nameserver 8.8.8.8 > /etc/resolv.conf
# Get dependencies.
RUN go get github.com/FiloSottile/gvt && \
    gvt restore
# Build the binary.
RUN go install -v ./...
# Update and generate the swagger spec.
RUN sed -i'' -e 's/\:8080//' ./api.go && \
    go get github.com/go-swagger/go-swagger/cmd/swagger && \
    swagger generate spec -o /go/src/app/api/static/swagger/swagger.json

FROM golang:1.14
WORKDIR /go/src/app/api
# Copy over all files.
COPY --from=0 /usr/src/ui/index.html /go/src/app/ui/
COPY --from=0 /usr/src/ui/static /go/src/app/ui/static/
COPY --from=1 /go/bin/api /go/bin
COPY --from=1 /go/src/app/api/static /go/src/app/ui/static/
EXPOSE 8080
CMD ["api"]