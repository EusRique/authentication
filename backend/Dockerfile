FROM golang:1.18

WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    go install github.com/spf13/cobra-cli@latest

CMD ["tail", "-f", "/dev/null"]