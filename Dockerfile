FROM golang:1.6

USER nobody

RUN mkdir -p /go/src/github.com/dtylman/tiritis
WORKDIR /go/src/github.com/dtylman/tiritis

COPY . /go/src/github.com/dtylman/tiritis
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]
