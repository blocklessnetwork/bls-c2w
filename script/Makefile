include inc.mk
export PREPARED_PKG=true

all: amd c2w riscv64 tools wasi

prepare_pkg:
	mkdir cache -p
	cd cache
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
	if [[ ! -d cache/wasi-vfs ]]; then \
		git clone https://github.com/kateinoigakukun/wasi-vfs.git --recurse-submodules cache/wasi-vfs &&\
		cd cache/wasi-vfs &&\
		git checkout "${WIZER_VFS_VERSION}" \
	; fi
	if [[ ! -d cache/wizer ]]; then \
		git clone https://github.com/bytecodealliance/wizer cache/wizer &&\
		cd cache/wizer &&\
		git checkout "${WIZER_VERSION}" \
	; fi
	if [[ ! -f cache/wasi-sdk.tar.gz ]]; then \
		curl -o cache/wasi-sdk.tar.gz -fSL https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-${WASI_SDK_VERSION}/wasi-sdk-${WASI_SDK_VERSION_FULL}-linux.tar.gz \
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

amd: prepare_pkg
	$(MAKE) -C amd64
c2w: prepare_pkg
	$(MAKE) -C c2w
riscv64: prepare_pkg
	$(MAKE) -C riscv64
tools: prepare_pkg
	$(MAKE) -C tools
wasi: prepare_pkg
	$(MAKE) -C wasi