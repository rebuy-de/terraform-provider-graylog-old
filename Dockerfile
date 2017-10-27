# Source: https://github.com/rebuy-de/golang-template
# Version: 1.3.2-snapshot

FROM golang:1.9-alpine

RUN apk add --no-cache git make

# Configure Go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Install Go Tools
RUN go get -u golang.org/x/lint/golint

# Install Glide
RUN go get -u github.com/Masterminds/glide/...

WORKDIR /go/src/github.com/Masterminds/glide

RUN git checkout v0.12.3
RUN go install

COPY . /go/src/github.com/rebuy-de/terraform-provider-graylog
WORKDIR /go/src/github.com/rebuy-de/terraform-provider-graylog
RUN CGO_ENABLED=0 make install

ENTRYPOINT ["/go/bin/terraform-provider-graylog"]
