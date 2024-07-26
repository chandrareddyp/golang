package main

import (
	"net"

	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}