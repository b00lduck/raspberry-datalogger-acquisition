FROM hypriot/rpi-golang
WORKDIR /gopath1.5/src/github.com/b00lduck/raspberry-datalogger-acquisition
CMD ["raspberry-datalogger-acquisition"]

ADD .  /gopath1.5/src/github.com/b00lduck/raspberry-datalogger-acquisition
RUN go get
RUN go build
USER root
