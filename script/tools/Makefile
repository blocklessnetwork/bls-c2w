include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-c2w-tools

all: build

build:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/assets cache/ 
	cp -rf ../cache/wasi-vfs cache/ 
	cp -rf ../cache/wizer cache/ 
else
	if [[ ! -d cache/assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm cache/assets; fi
	if [[ ! -d cache/wasi-vfs ]]; then \
		git clone https://github.com/kateinoigakukun/wasi-vfs.git  cache/wasi-vfs --recurse-submodules &&\
		cd  cache/wasi-vfs &&\
		git checkout "${WIZER_VFS_VERSION}" \
	; fi
	if [[ ! -d  cache/wizer ]]; then \
		git clone https://github.com/bytecodealliance/wizer  cache/wizer &&\
		cd cache/wizer &&\
		git checkout "${WIZER_VERSION}" \
	; fi
endif
	docker build -t ghcr.io/blocklessnetwork/${IMAGE_NAME}  -f Dockerfile cache

push:
	docker push --platform=linux/${ARCH} ghcr.io/blocklessnetwork/${IMAGE_NAME}
	docker tag  ghcr.io/blocklessnetwork/${IMAGE_NAME} ghcr.io/blocklessnetwork/${IMAGE_NAME}:${IMAGE_TAG_VERSION}
	docker push --platform=linux/${ARCH} ghcr.io/blocklessnetwork/${IMAGE_NAME}:${IMAGE_TAG_VERSION}

clean:
	rm -rf wizer

