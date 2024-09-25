include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-tinyemu-wasi-19

all: build

build:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/wasi-sdk.tar.gz cache/ 
	cp -rf ../cache/wasi-vfs cache/ 
	cp -rf ../cache/Bochs cache/ 
	cp -rf ../cache/wizer.h cache/ 
	cp -rf ../cache/tinyemu cache/ 
	cp -rf ../cache/Bochs cache/ 
	cp -rf ../cache/binaryen.tar.gz cache/
else
	if [[ ! -f cache/wasi-sdk.tar.gz ]]; then \
		curl -o cache/wasi-sdk.tar.gz -fSL https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-${WASI_SDK_VERSION}/wasi-sdk-${WASI_SDK_VERSION_FULL}-linux.tar.gz \
	; fi

	if [[ ! -d cache/wasi-vfs ]]; then \
		git clone https://github.com/kateinoigakukun/wasi-vfs.git cache/wasi-vfs --recurse-submodules  &&\
		cd cache/wasi-vfs &&\
		git checkout "${WASI_VFS_VERSION}" \
	; fi

	if [[ ! -d cache/Bochs ]]; then \
		git clone ${BOCHS_REPO} cache/Bochs &&\
		cd cache/Bochs &&\
		git checkout ${BOCHS_REPO_VERSION} \
	; fi

	if [[ ! -f cache/wizer.h ]]; then \
		curl -o cache/wizer.h -L https://raw.githubusercontent.com/bytecodealliance/wizer/${WIZER_VERSION}/include/wizer.h \
	; fi

	if [[ ! -d cache/tinyemu ]]; then \
		git clone ${TINYEMU_REPO} cache/tinyemu  &&\
		cd cache/tinyemu &&\
		git checkout ${TINYEMU_REPO_VERSION} \
	; fi

	
	if [[ ! -f cache/binaryen.tar.gz ]]; then \
		curl -o cache/binaryen.tar.gz -L https://github.com/WebAssembly/binaryen/releases/download/version_${BINARYEN_VERSION}/binaryen-version_${BINARYEN_VERSION}-x86_64-linux.tar.gz \
	; fi
endif
	docker build -t ghcr.io/blocklessnetwork/${IMAGE_NAME}  -f Dockerfile cache

push:
	docker push ghcr.io/blocklessnetwork/${IMAGE_NAME}

clean:
	rm -rf binaryen.tar.gz tinyemu wasi-vfs wasi-sdk.tar.gz wizer.h Bochs

