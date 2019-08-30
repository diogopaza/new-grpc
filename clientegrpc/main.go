package main

import (
	pb "clientegrpc/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	fmt.Println("Rodando na porta ", 4040)
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAddServiceClienteServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) VerCliente(ctx context.Context, in *pb.RequisicaoCliente) (*pb.RetornoCliente, error) {
	/*
		clientes := []*pb.RetornoCliente{
			&pb.RetornoCliente{
				IdResultado:   1,
				NomeResultado: "vitor",
			},
			&pb.RetornoCliente{
				IdResultado:   2,
				NomeResultado: "beto",
			},
		}*/
	fmt.Println(in.GetNome())
	nome := "antonio grpc"
	fmt.Println("estou grpc ver cliente")

	return &pb.RetornoCliente{NomeResultado: nome}, nil

}
