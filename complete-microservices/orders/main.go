package main

import (
	"context"
	"log"
	"net"

	"github.com/GustavoCesarSantos/go-expert/complete-microservices/common"
	"google.golang.org/grpc"
)

var (
    grpcAddr = common.EnvString("GRPC_ADDR", "localhost:3001")
)

func main() {
    grpcServer := grpc.NewServer()
    l, err := net.Listen("tcp", grpcAddr)
    if err != nil {
        log.Fatalf("failed to listem: %v", err)
    }
    defer l.Close()
    store := NewStore()
    svc := NewService(store)
    NewGRPCHandler(grpcServer, svc)
    svc.CreateOrder(context.Background())
    log.Println("GRPC Server started at ", grpcAddr)
    if err := grpcServer.Serve(l); err != nil {
        log.Fatal(err.Error())
    }
}
