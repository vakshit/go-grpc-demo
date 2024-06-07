package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/demo"
	"net"
)

type DemoServer struct {
	demo.UnimplementedDemoServer
}

func (s DemoServer) Create(ctx context.Context, req *demo.DemoRequest) (*demo.DemoResponse, error) {
	return &demo.DemoResponse{
		Pdf:  req.To,
		Docx: req.Amount,
	}, nil
}

func (s DemoServer) Delulu(ctx context.Context, req *demo.DeluluRequest) (*demo.DeluluRequest, error) {
	return req, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	//defer lis.Close();
	serverRegisterer := grpc.NewServer()
	service := &DemoServer{}
	demo.RegisterDemoServer(serverRegisterer, service)
	err = serverRegisterer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
