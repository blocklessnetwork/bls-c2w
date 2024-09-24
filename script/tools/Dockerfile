ARG GOPROXY=

FROM scratch AS assets
COPY . .

FROM rust:1.74.1-buster AS rust-common
RUN apt-get update -y
WORKDIR /work
COPY --link --from=assets wasi-vfs  wasi-vfs
RUN cd wasi-vfs &&\
    cargo build --package wasi-vfs-cli --release && \
    mkdir -p /tools/wasi-vfs/  && \
    mv target/release/wasi-vfs /tools/wasi-vfs/ && \
    cargo clean

WORKDIR /work
COPY --link --from=assets wizer wizer
WORKDIR /work/wizer/
RUN cargo build --bin wizer --all-features --release && \
    RUSTFLAGS="-C link-arg=-s" cargo build --bin wizer --all-features && \
    mkdir -p /tools/wizer/ && \
    mv include target/release/wizer /tools/wizer/ && \
    mv target/debug/wizer /tools/wizer/wizer-debug && \
    cargo clean


FROM golang:1.22-bullseye AS golang-base
ARG GOPROXY
COPY --link --from=assets /assets/ /work/
RUN if test "${GOPROXY}" = "true" ; then go env -w GOPROXY=${GOPROXY},direct ; fi
WORKDIR /work
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o /bin/create-spec ./cmd/create-spec

FROM scratch as output
COPY --link --from=rust-common /tools/wasi-vfs /tools/wasi-vfs
COPY --link --from=rust-common /tools/wizer /tools/wizer
COPY --link --from=golang-base /bin/create-spec /tools/create-spec

FROM output

