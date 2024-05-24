package main

import (
	"log"
	"net/http"

	"github.com/GustavoCesarSantos/go-expert/complete-microservices/common"
	pb "github.com/GustavoCesarSantos/go-expert/complete-microservices/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
    httpAddr = common.EnvString("HTTP_ADDR", ":3000")
    orderServiceAddr = "localhost:3001"
)

func main() {
    conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to dial server: %v", err)
    }
    defer conn.Close()
    log.Println("Dialing orders service at ", orderServiceAddr)
    orderClient := pb.NewOrderServiceClient(conn)
    mux := http.NewServeMux()
    handler := NewHandler(orderClient)
    handler.registerRoutes(mux)
    log.Printf("Starting HTTP server at %s", httpAddr)
    if err := http.ListenAndServe(httpAddr, mux); err != nil {
        log.Fatal("Failed to start http server")
    }
}
