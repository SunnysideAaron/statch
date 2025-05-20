FROM golang:1.23.6-bookworm

WORKDIR /

# RUN go clean -modcache

COPY . ./statch

WORKDIR /statch
