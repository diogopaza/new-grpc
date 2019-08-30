package main

import (
	"net"
	pb "servergrpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAddServicePessoaServer(srv, &pb.Server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
