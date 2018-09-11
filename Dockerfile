FROM golang:1.9-alpine

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN apk add librdkafka-dev build-base
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/malike/go-kafka-alert
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN GOOS=linux go build -a -o /go-kafka-alert .

COPY /go-kafka-alert ./
COPY /configuration.json ./
ENTRYPOINT ["./go-kafka-alert"]
# CMD [ "profile", "config" ]