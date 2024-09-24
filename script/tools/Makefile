include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-c2w-tools

all: build

build:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/assets . 
	cp -rf ../cache/wasi-vfs . 
	cp -rf ../cache/wizer . 
else
	if [[ ! -d assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm assets; fi
	if [[ ! -d wasi-vfs ]]; then \
		git clone https://github.com/kateinoigakukun/wasi-vfs.git --recurse-submodules &&\
		cd wasi-vfs &&\
		git checkout "${WIZER_VFS_VERSION}" \
	; fi
	if [[ ! -d wizer ]]; then \
		git clone https://github.com/bytecodealliance/wizer &&\
		cd wizer &&\
		git checkout "${WIZER_VERSION}" \
	; fi
endif
	docker build -t joing/${IMAGE_NAME}  -f Dockerfile . --progress=plain

push:
	docker push joing/${IMAGE_NAME}

clean:
	rm -rf wizer

