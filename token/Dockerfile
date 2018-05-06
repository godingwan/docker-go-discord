FROM golang

COPY ./app /go/src/github.com/user/token/app
WORKDIR /go/src/github.com/user/token/app

RUN go get ./
RUN go build

CMD ["./app"]
