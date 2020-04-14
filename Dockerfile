FROM golang:1.14.2

ADD . /opt/mirei-tts

WORKDIR /opt/mirei-tts

RUN go build -trimpath -o main application.go

ENTRYPOINT [ "/opt/mirei-tts/main" ]
