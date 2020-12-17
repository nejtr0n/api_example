FROM golang:1.15.6-alpine
RUN apk add --no-cache git protoc && \
    go get github.com/smartystreets/goconvey && \
    go get github.com/google/wire/cmd/wire && \
    go get github.com/golang/protobuf/protoc-gen-go && \
    go get github.com/gogo/protobuf/protoc-gen-gofast
WORKDIR /app/src
CMD ["/go/bin/goconvey", "-host=0.0.0.0"]