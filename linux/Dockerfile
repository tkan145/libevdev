FROM ubuntu:20.04

# Disable interactive prompts on package installation
ENV DEBIAN_FRONTEND noninteractive

# Dependencies to get the git sources and go binaries
RUN apt-get update && apt-get install -y  --no-install-recommends \
        ca-certificates \
        curl \
        git \
        rsync \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Get the git sources
WORKDIR /git
RUN git config --global advice.detachedHead false
# Linux Kernel: Released 25 Apr 2021
RUN git clone --branch v5.12 --depth 1 https://kernel.googlesource.com/pub/scm/linux/kernel/git/torvalds/linux

# Get Go
ENV GOLANG_VERSION 1.16.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 be333ef18b3016e9d7cb7b1ff1fdb0cac800ca0be4cf2290fe613b3d069dfe0d

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV PATH /usr/local/go/bin:$PATH

# Linux build dependencies
RUN apt-get update && apt-get install -y  --no-install-recommends \
        bison gawk make python3 \
        gcc gcc-multilib \
        gettext texinfo \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Cross compilers (install recommended packages to get cross libc-dev)
RUN apt-get update && apt-get install -y \
        gcc-aarch64-linux-gnu       gcc-arm-linux-gnueabi     \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /build
ENTRYPOINT ["go", "run", "linux/mkall.go", "/git/linux"]

