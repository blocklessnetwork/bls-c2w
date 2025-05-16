# BLS-C2W

This project aims to compile docker image into webassembly format, leveraging pre-compilation techniques to significantly accelerate the entire compilation time.

## Download and install
Exceute the shell downloads the binary

```bash
curl -L https://raw.githubusercontent.com/blocklessnetwork/bls-c2w/refs/heads/main/download.sh |bash
```

after execute you should get output 

```
Installing Blockless C2W to $HOME/.blessnet/bin...
Please add follow line to your shell profile...
export PATH=$HOME/.blessnet/bin:$PATH
Install complete!
NAME:
   bls-c2w - container to wasm converter

USAGE:
   bls-c2w [options] image-name [output file]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dockerfile value   Custom location of Dockerfile (default: embedded to this command)
   --builder value      Bulider command to use (default: "docker")
   --target-arch value  target architecture of the source image to use (default: "amd64")
   --build-arg value    Additional build arguments
   --help, -h           show help
```

Set the PATH variable as instructed

## Make a wasm

### Hello world exmaple

Make a smallest `hello world` image.

```bash
mkdir -p hello_world
cd hello_world

cat > hello.go <<EOF
package main
func main() {
    println("hello world")
}
EOF
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello hello.go

cat > Dockerfile <<DEOF
FROM scratch
COPY hello /
CMD ["/hello"]
DEOF

docker build -t shello .
```

Convert the 'shello' image to wasm file

```bash
bls-c2w shello
```

After execute the  bls-c2w, the `bls-out.wasm` file will generated.

Before execute the wasm, download the bls-runtime

```bash
sh -c "curl https://raw.githubusercontent.com/blessnetwork/bls-runtime/refs/heads/main/install.sh | bash"

#execute the wasm file
bls-runtime bls-out.wasm
```



## Run with network

Make `http server` image.

```bash
mkdir -p http_server
cd http_server

cat > http.go <<EOF
package main

import (
    "fmt"
    "net/http"
)
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello world\n")
}

func main() {
    fmt.Println("http server.")
    go func() {
         mux := http.NewServeMux()
         mux.HandleFunc("/hello", hello)
         http.ListenAndServe(":8091", mux)
    }()
    mux := http.NewServeMux()

    mux.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", mux)
}
EOF
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o http http.go

cat > Dockerfile <<DEOF
FROM scratch
COPY http /
CMD ["/http"]
DEOF

docker build -t shttp .
```

Convert the 'shttp' image to wasm file

```bash
bls-c2w shttp
```


Make sure you have downloaded the bls-runtime before execute wasm with bls-c2wnet command

```bash
bls-c2wnet --invoke -p 0.0.0.0:8090:8090 bls-out.wasm --net=socket
```

Use curl for validate the result

```bash
curl http://localhost:8090/hello
```

