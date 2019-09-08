package main

import (
	"context"
	"fmt"
	"net"

	"github.com/confus1on/go-protobuff/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func (s *Server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}

func (s *Server) Multiply(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &pb.Response{Result: result}, nil
}

func main() {
	fmt.Println("GRPC Server")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
