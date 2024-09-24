include inc.mk
export PREPARED_PKG=true

all: amd

prepare_pkg:
	mkdir cache -p
	cd cache
	# if [[ ! -f linux.tar.gz ]]; then \
	# 	curl -Lo linux.tar.gz https://github.com/torvalds/linux/archive/refs/tags/v6.1.tar.gz \
	# ; fi
	# if [[ ! -d assets ]]; then git clone -b v0.6.5 https://github.com/ktock/container2wasm assets; fi
	# if [[ ! -f grub-${GRUB_VERSION}.tar.gz ]]; then \
	# 	curl -Lo grub-${GRUB_VERSION}.tar.gz https://ftp.gnu.org/gnu/grub/grub-${GRUB_VERSION}.tar.gz \
	# ; fi
	# if [[ ! -f busybox-${BUSYBOX_VERSION}.tar.bz ]]; then \
	# 	curl -Lo busybox-${BUSYBOX_VERSION}.tar.bz https://busybox.net/downloads/busybox-${BUSYBOX_VERSION}.tar.bz2 \
	# ; fi
	# if [[ ! -f runc.tar.gz ]]; then \
	# 	curl -Lo runc.tar.gz https://github.com/opencontainers/runc/archive/refs/tags/v1.1.12.tar.gz \
	# ; fi
	# if [[ ! -f tini.tar.gz ]]; then \
	# 	curl -Lo tini.tar.gz https://github.com/krallin/tini/archive/refs/tags/v0.19.0.tar.gz \
	# ; fi
	# if [[ ! -f vmtouch.tar.gz ]]; then \
	# 	curl -Lo vmtouch.tar.gz https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.tar.gz \
	# ; fi

	# if [[ ! -d Bochs ]]; then \
	# 	git clone ${BOCHS_REPO} Bochs  &&\
	# 	cd Bochs &&\
	# 	git checkout ${BOCHS_REPO_VERSION} \
	# ; fi
	# if [[ ! -d wasi-vfs ]]; then \
	# 	git clone https://github.com/kateinoigakukun/wasi-vfs.git --recurse-submodules &&\
	# 	cd wasi-vfs &&\
	# 	git checkout "${WIZER_VFS_VERSION}" \
	# ; fi
	# if [[ ! -d wizer ]]; then \
	# 	git clone https://github.com/bytecodealliance/wizer &&\
	# 	cd wizer &&\
	# 	git checkout "${WIZER_VERSION}" \
	# ; fi
	# if [[ ! -f wasi-sdk.tar.gz ]]; then \
	# 	curl -o wasi-sdk.tar.gz -fSL https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-${WASI_SDK_VERSION}/wasi-sdk-${WASI_SDK_VERSION_FULL}-linux.tar.gz \
	# ; fi
	# if [[ ! -f wizer.h ]]; then \
	# 	curl -o wizer.h -L https://raw.githubusercontent.com/bytecodealliance/wizer/${WIZER_VERSION}/include/wizer.h \
	# ; fi
	# if [[ ! -d tinyemu ]]; then \
	# 	git clone ${TINYEMU_REPO} tinyemu  &&\
	# 	cd tinyemu &&\
	# 	git checkout ${TINYEMU_REPO_VERSION} \
	# ; fi
	# if [[ ! -f binaryen.tar.gz ]]; then \
	# 	curl -o binaryen.tar.gz -L https://github.com/WebAssembly/binaryen/releases/download/version_${BINARYEN_VERSION}/binaryen-version_${BINARYEN_VERSION}-x86_64-linux.tar.gz \
	# ; fi
amd:
	$(MAKE) -C amd64
