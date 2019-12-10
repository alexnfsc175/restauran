# create image from the official Go image
FROM golang:alpine
RUN apk add --update tzdata \
  bash wget curl git;

# Create binary directory, install glide and fresh
RUN mkdir -p $GOPATH/bin && \
  go get github.com/pilu/fresh

# define work directory
ADD . /go/src/app
WORKDIR /go/src/app

EXPOSE 3001

# serve the app
CMD fresh -c runner.conf main.go