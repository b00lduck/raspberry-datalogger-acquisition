FROM rem/rpi-golang-1.7:latest

WORKDIR /gopath/src/github.com/b00lduck/raspberry-datalogger-acquisition
ENTRYPOINT ["raspberry-datalogger-acquisition"]

ADD . /gopath/src/github.com/b00lduck/raspberry-datalogger-acquisition
RUN go get
RUN go build
USER root
