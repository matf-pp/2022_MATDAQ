FROM rust:1.60

WORKDIR /usr/src/app

RUN apt update && apt-get -y install cmake protobuf-compiler build-essential libssl-dev 

COPY ./api/ /usr/src/api
COPY ./matching-engine .

RUN cargo build --bin grpc-server

