package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/yoshd/nc-ess-grpc-proxy/server/ess"
	"github.com/yoshd/nc-ess-grpc-proxy/pb"
)

func main() {
	lis, err := net.Listen("tcp", ":13009")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proxy := ess.NewProxy()
	pb.RegisterESSProxyServer(s, proxy)
	s.Serve(lis)
}
