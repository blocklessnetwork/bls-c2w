include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-riscv64-linux6.1

all: build-linux

build-linux:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/linux.tar.gz . 
	cp -rf ../cache/assets . 
	cp -rf ../cache/busybox-${BUSYBOX_VERSION}.tar.bz . 
	cp -rf ../cache/tini.tar.gz . 
	cp -rf ../cache/runc.tar.gz . 
	cp -rf ../cache/vmtouch.tar.gz . 
	cp -rf ../cache/riscv-pk . 
else
	if [[ ! -f cache/linux.tar.gz ]]; then \
		curl -Lo cache/linux.tar.gz https://github.com/torvalds/linux/archive/refs/tags/v6.1.tar.gz \
	; fi
	if [[ ! -d cache/assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm cache/assets; fi
	if [[ ! -f cache/busybox-${BUSYBOX_VERSION}.tar.bz ]]; then \
		curl -Lo cache/busybox-${BUSYBOX_VERSION}.tar.bz https://busybox.net/downloads/busybox-${BUSYBOX_VERSION}.tar.bz2 \
	; fi

	if [[ ! -f cache/tini.tar.gz ]]; then \
		curl -Lo cache/tini.tar.gz https://github.com/krallin/tini/archive/refs/tags/v0.19.0.tar.gz \
	; fi
	if [[ ! -f cache/runc.tar.gz ]]; then \
		curl -Lo cache/runc.tar.gz https://github.com/opencontainers/runc/archive/refs/tags/v1.1.12.tar.gz \
	; fi
	if [[ ! -f cache/vmtouch.tar.gz ]]; then \
		curl -Lo cache/vmtouch.tar.gz https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.tar.gz \
	; fi

	if [[ ! -d cache/riscv-pk ]]; then \
		git clone https://github.com/riscv-software-src/riscv-pk cache/riscv-pk && cd cache/riscv-pk && git checkout 7e9b671c0415dfd7b562ac934feb9380075d4aa2 \
	; fi
endif
	docker build --build-arg TARGETARCH=amd64 -t joing/${IMAGE_NAME} -f Dockerfile cache

push:
        docker push joing/${IMAGE_NAME}

clean:
        rm -rf assets  busybox-${BUSYBOX_VERSION}.tar.bz  linux.tar.gz  riscv-pk  runc.tar.gz  tini.tar.gz  vmtouch.tar.gz

