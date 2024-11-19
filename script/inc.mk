SHELL := /bin/bash
PREPARED_PKG = true
GRUB_VERSION=2.06
BUSYBOX_VERSION=1.36.1
BOCHS_REPO=https://github.com/blocklessnetwork/bls-bochs
BOCHS_REPO_VERSION=67eab5061bb69c9f87e51091a29c6344034e0581
TINYEMU_REPO=https://github.com/ktock/tinyemu-c2w.git
TINYEMU_REPO_VERSION=e4e9bd198f9c0505ab4c77a6a9d038059cd1474a
WIZER_VERSION=04e49c989542f2bf3a112d60fbf88a62cce2d0d0
WASI_SDK_VERSION=19
WASI_SDK_VERSION_FULL=${WASI_SDK_VERSION}.0
WASI_VFS_VERSION=v0.3.0
BINARYEN_VERSION=114
OS_ARCH=$(shell uname -m)
ARCH=amd64

ifeq ($(OS_ARCH), "x86_64")
	ARCH=amd64
endif

ifeq ($(OS_ARCH), "aarch64")
	ARCH=arm64
endif

ifeq ($(OS_ARCH), "arm64")
	ARCH=arm64
endif

IMAGE_TAG_VERSION=v0.0.1
