FROM golang:1.10.1-alpine3.8
RUN apk add --no-cache  --repository http://dl-3.alpinelinux.org/alpine/edge/community/ \
      bash              \
      gcc				\
      git 				\
      librdkafka-dev    \
      libressl-dev      \
      musl-dev          \
      zlib-dev			\
      wget 			&&  \
      #
      # Install dep
      wget -nv -O /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \
      chmod a+rx /usr/local/bin/dep
WORKDIR $GOPATH/src/github.com/malike/go-kafka-alert
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -o ./go-kafka-alert .

FROM scratch
COPY --from=builder /go-kafka-alert ./
COPY --from=builder /configuration.json ./
ENTRYPOINT ["./go-kafka-alert"]
# CMD [ "profile", "config" ]