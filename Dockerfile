FROM golang:1.10.3

WORKDIR /go/src/app

ADD hello-world.go .
ADD Gopkg.lock .
ADD Gopkg.toml .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure

CMD ["go", "run", "hello-world.go"]
