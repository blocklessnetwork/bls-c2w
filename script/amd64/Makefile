include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-amd64-linux6.1

all: build


build:
	mkdir cache -p
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/linux.tar.gz cache/ 
	cp -rf ../cache/assets cache/ 
	cp -rf ../cache/grub-${GRUB_VERSION}.tar.gz cache/ 
	cp -rf ../cache/busybox-${BUSYBOX_VERSION}.tar.bz cache/
	cp -rf ../cache/runc.tar.gz cache/
	cp -rf ../cache/tini.tar.gz cache/
	cp -rf ../cache/vmtouch.tar.gz cache/
	cp -rf ../cache/Bochs cache/
else
	if [[ ! -f cache/linux.tar.gz ]]; then \
		curl -Lo cache/linux.tar.gz https://github.com/torvalds/linux/archive/refs/tags/v6.1.tar.gz \
	; fi
	if [[ ! -d cache/assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm cache/assets; fi
	if [[ ! -f cache/grub-${GRUB_VERSION}.tar.gz ]]; then \
		curl -Lo cache/grub-${GRUB_VERSION}.tar.gz https://ftp.gnu.org/gnu/grub/grub-${GRUB_VERSION}.tar.gz \
	; fi
	if [[ ! -f cache/busybox-${BUSYBOX_VERSION}.tar.bz ]]; then \
		curl -Lo cache/busybox-${BUSYBOX_VERSION}.tar.bz https://busybox.net/downloads/busybox-${BUSYBOX_VERSION}.tar.bz2 \
	; fi
	if [[ ! -f cache/runc.tar.gz ]]; then \
		curl -Lo cache/runc.tar.gz https://github.com/opencontainers/runc/archive/refs/tags/v1.1.12.tar.gz \
	; fi
	if [[ ! -f cache/tini.tar.gz ]]; then \
		curl -Lo cache/tini.tar.gz https://github.com/krallin/tini/archive/refs/tags/v0.19.0.tar.gz \
	; fi
	if [[ ! -f cache/vmtouch.tar.gz ]]; then \
		curl -Lo cache/vmtouch.tar.gz https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.tar.gz \
	; fi

	if [[ ! -d cache/Bochs ]]; then \
		git clone ${BOCHS_REPO} cache/Bochs  &&\
		cd cache/Bochs &&\
		git checkout ${BOCHS_REPO_VERSION} \
	; fi
endif
	docker build -t blocklessnetwork/${IMAGE_NAME}  -f Dockerfile cache

push:
	docker push ghcr.io/blocklessnetwork/${IMAGE_NAME}

clean:
	rm -rf grub-${GRUB_VERSION}.tar.gz

