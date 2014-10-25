FROM golang:latest

RUN apt-get update && apt-get install -y libinotifytools-dev \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/lucas-clemente/treewatch
VOLUME /go/src/github.com/lucas-clemente/treewatch

RUN go get github.com/onsi/ginkgo/ginkgo github.com/onsi/gomega

CMD ginkgo watch
