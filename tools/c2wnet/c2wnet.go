package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	gvntypes "github.com/containers/gvisor-tap-vsock/pkg/types"
	gvnvirtualnetwork "github.com/containers/gvisor-tap-vsock/pkg/virtualnetwork"
	"golang.org/x/net/websocket"
)

const (
	gatewayIP = "192.168.127.1"
	vmIP      = "192.168.127.3"
	vmMAC     = "02:00:00:00:00:01"
)

func main() {
	var (
		portFlags   sliceFlags
		volumeFlags sliceFlags
	)
	flag.Var(&portFlags, "p", "map port between host and guest (host:guest). -mac must be set correctly.")
	flag.Var(&volumeFlags, "v", "map directory between host and guest (host_dir::guest_dir or host_dir)")
	var (
		debug    = flag.Bool("debug", false, "enable debug print")
		listenWS = flag.Bool("listen-ws", false, "listen on a websocket port specified as argument")
		invoke   = flag.Bool("invoke", false, "invoke the container with NW support")
		mac      = flag.String("mac", vmMAC, "mac address assigned to the container")
		wasiAddr = flag.String("wasi-addr", "127.0.0.1:1234", "IP address used to communicate between wasi and network stack (valid only with invoke flag)") // TODO: automatically use empty random port or unix socket
		envFile  = flag.String("env-file", "", "path to environment file")
		listenfd  = flag.String("listenfd", "", "listenfd is listen file descriptor in wasi.")
		runtimeArgs  = flag.String("runtime-args", "", "arguments for the runtime.")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		panic("specify args")
	}
	socketAddr := args[0]
	forwards := make(map[string]string)
	for _, p := range portFlags {
		parts := strings.Split(p, ":")
		switch len(parts) {
		case 3:
			// IP:PORT1:PORT2
			forwards[strings.Join(parts[0:2], ":")] = strings.Join([]string{vmIP, parts[2]}, ":")
		case 2:
			// PORT1:PORT2
			forwards["0.0.0.0:"+parts[0]] = vmIP + ":" + parts[1]
		}
	}
	if *debug {
		fmt.Fprintf(os.Stderr, "port mapping: %+v\n", forwards)
		fmt.Fprintf(os.Stderr, "volume mapping: %+v\n", volumeFlags)
	}
	config := &gvntypes.Configuration{
		Debug:             *debug,
		MTU:               1500,
		Subnet:            "192.168.127.0/24",
		GatewayIP:         gatewayIP,
		GatewayMacAddress: "5a:94:ef:e4:0c:dd",
		DHCPStaticLeases: map[string]string{
			vmIP: *mac,
		},
		Forwards: forwards,
		NAT: map[string]string{
			"192.168.127.254": "127.0.0.1",
		},
		GatewayVirtualIPs: []string{"192.168.127.254"},
		Protocol:          gvntypes.QemuProtocol,
	}
	vn, err := gvnvirtualnetwork.New(config)
	if err != nil {
		panic(err)
	}
	if *invoke {
		go func() {
			var conn net.Conn
			for i := 0; i < 10; i++ {
				time.Sleep(1 * time.Second)
				fmt.Fprintf(os.Stderr, "connecting to NW...\n")
				conn, err = net.Dial("tcp", *wasiAddr)
				if err == nil {
					break
				}
				fmt.Fprintf(os.Stderr, "failed connecting to NW: %v\n", err)
			}
			if conn == nil {
				panic("failed to connect to vm")
			}
			// We register our VM network as a qemu "-netdev socket".
			if err := vn.AcceptQemu(context.TODO(), conn); err != nil {
				fmt.Fprintf(os.Stderr, "failed AcceptQemu: %v\n", err)
			}
		}()
		var tcplisten string = "--tcplisten=" + *wasiAddr;
		if *listenfd != "" {
			tcplisten = tcplisten + "::" + *listenfd
		}
		cmdArgs := []string{tcplisten, "--env='LISTEN_FDS=1'"}
		if *runtimeArgs != "" {
			args := strings.Split(*runtimeArgs, " ")
			cmdArgs = append(cmdArgs, args...)
		}
		// Add env-file parameter if provided
		if *envFile != "" {
			cmdArgs = append(cmdArgs, "--env-file="+*envFile)
		}
		// Add volume mounts
		for _, volume := range volumeFlags {
			// The volume string can be either "host_dir" or "host_dir::guest_dir"
			cmdArgs = append(cmdArgs, "--dir="+volume)
		}
		// Append the remaining arguments
		cmdArgs = append(cmdArgs, "--")

		//network listenfd.
		nargs := []string{}
		for _, v := range args {
			if v == "--net=socket" {
				var params string = "--net=socket"
				if *listenfd != "" {
					params = params + *listenfd
				}
				nargs = append(nargs, params)
			} else {
				nargs = append(nargs, v)
			}
		}
		cmdArgs = append(cmdArgs, nargs...)

		if *debug {
			fmt.Fprintf(os.Stderr, "executing command:\nbls-runtime %s\n\n", strings.Join(cmdArgs, " "))
		}

		var cmd = exec.Command("bls-runtime", cmdArgs...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		return
	}
	if *listenWS {
		http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
			ws.PayloadType = websocket.BinaryFrame
			if err := vn.AcceptQemu(context.TODO(), ws); err != nil {
				fmt.Fprintf(os.Stderr, "forwarding finished: %v\n", err)
			}
		}))
		if err := http.ListenAndServe(socketAddr, nil); err != nil {
			panic(err)
		}
		return
	}
	conn, err := net.Dial("tcp", socketAddr)
	if err != nil {
		panic(err)
	}
	// We register our VM network as a qemu "-netdev socket".
	if err := vn.AcceptQemu(context.TODO(), conn); err != nil {
		panic(err)
	}
}

type sliceFlags []string

func (f *sliceFlags) String() string {
	var s []string = *f
	return fmt.Sprintf("%v", s)
}

func (f *sliceFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}
