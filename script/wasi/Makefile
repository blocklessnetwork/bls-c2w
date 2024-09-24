include ../inc.mk
PREPARED_PKG?=false
IMAGE_NAME=bls-tinyemu-wasi-19

all: build

build:
	mkdir cache -p
	cd cache
ifeq ($(PREPARED_PKG), true)
	cp -rf ../cache/wasi-sdk.tar.gz . 
	cp -rf ../cache/wasi-vfs . 
	cp -rf ../cache/Bochs . 
	cp -rf ../cache/wizer.h . 
	cp -rf ../cache/tinyemu . 
	cp -rf ../cache/Bochs . 
	cp -rf ../cache/binaryen.tar.gz . 
else
	if [[ ! -f wasi-sdk.tar.gz ]]; then \
		curl -o wasi-sdk.tar.gz -fSL https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-${WASI_SDK_VERSION}/wasi-sdk-${WASI_SDK_VERSION_FULL}-linux.tar.gz \
	; fi

	if [[ ! -d wasi-vfs ]]; then \
		git clone https://github.com/kateinoigakukun/wasi-vfs.git --recurse-submodules &&\
		cd wasi-vfs &&\
		git checkout "${WASI_VFS_VERSION}" \
	; fi

	if [[ ! -d Bochs ]]; then \
		git clone ${BOCHS_REPO} Bochs &&\
		cd Bochs &&\
		git checkout ${BOCHS_REPO_VERSION} \
	; fi

	if [[ ! -f wizer.h ]]; then \
		curl -o wizer.h -L https://raw.githubusercontent.com/bytecodealliance/wizer/${WIZER_VERSION}/include/wizer.h \
	; fi

	if [[ ! -d tinyemu ]]; then \
		git clone ${TINYEMU_REPO} tinyemu  &&\
		cd tinyemu &&\
		git checkout ${TINYEMU_REPO_VERSION} \
	; fi

	if [[ ! -d Bochs ]]; then \
		git clone ${BOCHS_REPO} Bochs  &&\
		cd Bochs &&\
		git checkout ${BOCHS_REPO_VERSION} \
	; fi
	if [[ ! -f binaryen.tar.gz ]]; then \
		curl -o binaryen.tar.gz -L https://github.com/WebAssembly/binaryen/releases/download/version_${BINARYEN_VERSION}/binaryen-version_${BINARYEN_VERSION}-x86_64-linux.tar.gz \
	; fi
endif
	docker build -t joing/${IMAGE_NAME}  -f Dockerfile .

push:
	docker push joing/${IMAGE_NAME}

clean:
	rm -rf binaryen.tar.gz tinyemu wasi-vfs wasi-sdk.tar.gz wizer.h Bochs

