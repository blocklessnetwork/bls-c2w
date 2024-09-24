include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-amd64-linux6.1

all: build1

build1:
	echo 1
build:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/linux.tar.gz . 
	cp -rf ../cache/assets . 
	cp -rf ../cache/grub-${GRUB_VERSION}.tar.gz . 
	cp -rf ../cache/busybox-${BUSYBOX_VERSION}.tar.bz .
	cp -rf ../cache/runc.tar.gz .
	cp -rf ../cache/tini.tar.gz .
	cp -rf ../cache/vmtouch.tar.gz .
	cp -rf ../cache/Bochs .
else
	if [[ ! -f linux.tar.gz ]]; then \
		curl -Lo linux.tar.gz https://github.com/torvalds/linux/archive/refs/tags/v6.1.tar.gz \
	; fi
	if [[ ! -d assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm assets; fi
	if [[ ! -f grub-${GRUB_VERSION}.tar.gz ]]; then \
		curl -Lo grub-${GRUB_VERSION}.tar.gz https://ftp.gnu.org/gnu/grub/grub-${GRUB_VERSION}.tar.gz \
	; fi
	if [[ ! -f busybox-${BUSYBOX_VERSION}.tar.bz ]]; then \
		curl -Lo busybox-${BUSYBOX_VERSION}.tar.bz https://busybox.net/downloads/busybox-${BUSYBOX_VERSION}.tar.bz2 \
	; fi
	if [[ ! -f runc.tar.gz ]]; then \
		curl -Lo runc.tar.gz https://github.com/opencontainers/runc/archive/refs/tags/v1.1.12.tar.gz \
	; fi
	if [[ ! -f tini.tar.gz ]]; then \
		curl -Lo tini.tar.gz https://github.com/krallin/tini/archive/refs/tags/v0.19.0.tar.gz \
	; fi
	if [[ ! -f vmtouch.tar.gz ]]; then \
		curl -Lo vmtouch.tar.gz https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.tar.gz \
	; fi

	if [[ ! -d Bochs ]]; then \
		git clone ${BOCHS_REPO} Bochs  &&\
		cd Bochs &&\
		git checkout ${BOCHS_REPO_VERSION} \
	; fi
endif
	docker build -t joing/${IMAGE_NAME}  -f Dockerfile cache

push:
	docker push joing/${IMAGE_NAME}

clean:
	rm -rf grub-${GRUB_VERSION}.tar.gz

