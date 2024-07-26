package main

import (
	"context"
	"log"
	"net"

	grpc2 "github.com/chandrareddyp/golang/19grpc/grpc2/invoicer"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	ServiceRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}
	grpc2.RegisterInvoicerServer(ServiceRegistrar, service)
	
	if err := ServiceRegistrar.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}

type myInvoiceServer struct {
	grpc2.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(ctx context.Context, req *grpc2.CreateRequest) (*grpc2.CreateResponse, error) {
	return &grpc2.CreateResponse{
		Pdf: []byte("some pdf"),
		Docx: []byte("some docx"),
		}, nil
}