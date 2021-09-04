FROM golang:alpine

WORKDIR /app

COPY .  .

RUN go get -d  -v  \
    && go build -o demo .

ENTRYPOINT ["demo"]