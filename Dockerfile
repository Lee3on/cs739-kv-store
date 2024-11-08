# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Set environment variables to non-interactive
ENV DEBIAN_FRONTEND=noninteractive

# Install necessary packages and dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    autoconf \
    libtool \
    pkg-config \
    git \
    wget \
    unzip \
    cmake \
    libssl-dev \
    libprotobuf-dev \
    protobuf-compiler \
    telnet \
    && rm -rf /var/lib/apt/lists/*

# Install gRPC and Protocol Buffers C++ library
WORKDIR /tmp
RUN git clone -b v1.49.0 https://github.com/grpc/grpc
WORKDIR /tmp/grpc
RUN git submodule update --init --recursive
RUN mkdir -p cmake/build
WORKDIR /tmp/grpc/cmake/build
RUN cmake ../.. \
    -DgRPC_INSTALL=ON \
    -DgRPC_BUILD_TESTS=OFF \
    -DCMAKE_INSTALL_PREFIX=/usr/local
RUN make -j$(nproc)
RUN make install

# Set the working directory
WORKDIR /app

COPY ./server/server ./load_balancer/load_balancer run_servers.sh ./proto/kv739.proto ./client/client.cpp ./client/kv739_client.cpp ./client/kv739_client.h ./client/kv739_test*.cpp ./
COPY ./config ./config

RUN mkdir -p storage

# Generate C++ code from the proto file
RUN mkdir -p generated
RUN protoc --cpp_out=./generated --grpc_out=./generated \
    --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` kv739.proto

# Compile the client code
RUN g++ -std=c++11 client.cpp kv739_client.cpp \
    ./generated/kv739.pb.cc \
    ./generated/kv739.grpc.pb.cc \
    -I./generated \
    `pkg-config --cflags protobuf grpc grpc++` \
    `pkg-config --libs protobuf grpc grpc++` \
    -o kv739_client

# Compile the test code
RUN g++ -std=c++11 client.cpp kv739_test.cpp \
    ./generated/kv739.pb.cc \
    ./generated/kv739.grpc.pb.cc \
    -I./generated \
    `pkg-config --cflags protobuf grpc grpc++` \
    `pkg-config --libs protobuf grpc grpc++` \
    -o kv739_test

RUN g++ -std=c++11 client.cpp kv739_test1.cpp \
    ./generated/kv739.pb.cc \
    ./generated/kv739.grpc.pb.cc \
    -I./generated \
    `pkg-config --cflags protobuf grpc grpc++` \
    `pkg-config --libs protobuf grpc grpc++` \
    -o kv739_test1

# Command to run when starting the container
CMD ["sh", "-c", "/bin/bash run_servers.sh && sleep 10 && ./kv739_client"]