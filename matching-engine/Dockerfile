#FROM rust:1.60
#
#WORKDIR /usr/src/app
#
#RUN apt update && apt-get -y install cmake protobuf-compiler build-essential libssl-dev 
#
#COPY ./api/ /usr/src/api
#COPY ./matching-engine .
#
#RUN cargo build --bin grpc-server

FROM rust:1.60 as builder

WORKDIR /usr/src
RUN USER=root cargo new matching-engine

COPY ./matching-engine/Cargo.toml ./matching-engine/Cargo.lock ./matching-engine/rust-toolchain /usr/src/matching-engine/
COPY ./api .
WORKDIR /usr/src/matching-engine
RUN apt update && apt-get -y install cmake protobuf-compiler build-essential libssl-dev 
RUN mkdir -p src \
  && echo "fn main() {}" > ./src/server.rs \
  && echo "fn main() {}" > ./src/client.rs \
  && cargo build --release
RUN rm src/*.rs

COPY ./api /usr/src/api
ADD ./matching-engine ./
RUN cargo build --release \
  && mv target/release/* /bin/ \
  && rm -rf /usr/src/matching-engine

WORKDIR /
