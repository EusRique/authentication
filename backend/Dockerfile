FROM golang:1.18

WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    go get github.com/spf13/cobra/cobra@latest

CMD ["tail", "-f", "/dev/null"]