# BLS-C2W

This project aims to compile docker image into webassembly format, leveraging pre-compilation techniques to significantly accelerate the entire compilation time.

## Make a wasm

```bash
$ bls-c2w image-name
```

## Run with network

```bash
$ bls-c2wnet --invoke -p 0.0.0.0:8090:8090 bls-out.wasm --net=socket
```

