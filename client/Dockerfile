FROM golang:1.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

# copy the source code into the container and build it
COPY . .
RUN make

# enable colored output
ENV TERM xterm-256color

