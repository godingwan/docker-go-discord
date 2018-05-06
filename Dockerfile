FROM golang

COPY ./app /go/src/github.com/user/disgord/app
WORKDIR /go/src/github.com/user/disgord/app

RUN go get ./
RUN go build

CMD ["./app"]
