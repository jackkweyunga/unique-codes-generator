package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "unique-codes-generator/generator"
)

var (
	port = flag.Int("port", 5000, "The server port")
)

type generatorServer struct {
	pb.UnimplementedGeneratorServer
}

func (s *generatorServer) GenerateUniqueCode(ctx context.Context, options *pb.GenerateUniqueCodeRequest) (*pb.Code, error) {
	code, err := pb.GenIdWithSonyFlake()
	if err != nil {
		return nil, err
	}
	reply := &pb.Code{
		Code: code,
	}
	return reply, nil
}

func (s *generatorServer) GenerateUniqueCodes(ctx context.Context, options *pb.GenerateUniqueCodesRequest) (*pb.Codes, error) {
	codes, err := pb.GenIdsWithSonyFlake(options.Count)
	if err != nil {
		return nil, err
	}
	reply := &pb.Codes{
		Codes: codes,
	}
	return reply, nil
}

func newServer() *generatorServer {
	s := &generatorServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGeneratorServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	log.Printf("gPRC Server started at http://localhost:%v", port)
	if err != nil {
		return
	}
}
