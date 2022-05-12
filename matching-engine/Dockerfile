FROM rust:1.60

WORKDIR /usr/src/app

COPY ./api/ /usr/src/api
COPY ./matching-engine .


RUN apt update && apt-get -y install cmake protobuf-compiler build-essential libssl-dev 

RUN cargo install --path .
