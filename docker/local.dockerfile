# create image from the official Go image
FROM golang:alpine

RUN apk add --update tzdata \
  bash g++ wget curl git;

# Create binary directory, install glide and fresh
RUN mkdir -p $GOPATH/bin && \
  go get github.com/go-delve/delve/cmd/dlv

# define work directory
ADD . /go/src/app
WORKDIR /go/src/app

EXPOSE 3001 2345

# serve the app
CMD dlv debug --output=./tmp/restaurant --listen=:2345 --headless --api-version=2 --log --continue --accept-multiclient