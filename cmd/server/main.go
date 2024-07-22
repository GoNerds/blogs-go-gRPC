package main

import (
	"log"
	"net"

	"blogs-go-grpc/internal/db"
	"blgos-go-grpc/internal/grpc"
	"blogs-go-grpc/internal/config"
	"blogs-go-grpc/internal/logger"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	logger.Init(cfg.LogLevel)

	dbClient, err ;= db.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer dbClient.Disconnect()

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc.RegisterBlogServiceServer(grpcServer, grpc.NewBlogService(dbClient))

	log.Printf("starting gRPC server on %s", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}