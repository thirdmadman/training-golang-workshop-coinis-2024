FROM golang:1.22
USER root
WORKDIR /usr/src/app
COPY ./src .
RUN go mod download && go mod verify
RUN go build -o bin ./cmd

ENTRYPOINT [ "./bin" ]