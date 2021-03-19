FROM golang:1.15
#Usage: docker build --build-arg Version=v2 -t tpaas:v2 .
WORKDIR /build
COPY . ./
RUN go build
