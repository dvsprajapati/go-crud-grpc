package main

import (
	"go-crud-grpc/db"
	"go-crud-grpc/server"
	"log"
	"net"

	pb "go-crud-grpc/go_crud_grpc/proto"

	"google.golang.org/grpc"
)

// used command to create proto file  => protoc --go_out=. --go-grpc_out=. proto/user.proto
func main() {
	err := db.InitMongo()

	if err != nil {
		log.Fatal("Mongo Connection Failed", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})
	log.Println("Grpc server is running on 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve the grpc server", err)
	}
}
