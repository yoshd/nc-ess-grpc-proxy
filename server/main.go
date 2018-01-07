package main

import (
	"net"
	"flag"

	"google.golang.org/grpc"
	"github.com/golang/glog"

	"github.com/yoshd/nc-ess-grpc-proxy/server/ess"
	"github.com/yoshd/nc-ess-grpc-proxy/pb"
)

var (
	port = flag.String("port", ":13009", "port :port")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		glog.Fatal(err)
	}

	s := grpc.NewServer()
	proxy := ess.NewProxy()
	pb.RegisterESSProxyServer(s, proxy)
	s.Serve(lis)
}
