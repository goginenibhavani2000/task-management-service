package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/goginenibhavani2000/task-management-service/internal/domain"
	"github.com/goginenibhavani2000/task-management-service/internal/repository"
	"github.com/goginenibhavani2000/task-management-service/internal/service"
	transport "github.com/goginenibhavani2000/task-management-service/internal/transport/grpc"
	"github.com/goginenibhavani2000/task-management-service/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=task_management port=5432 sslmode=disable" // 1. Initialize DB (GORM)
	//1.Connect to DB
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&domain.Task{})

	// 2. Initialize Layers
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	hdl := transport.NewTaskHandler(svc)

	// 3. Create a listener for gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 4. Start gRPC Server in a Goroutine
	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, hdl) // 'hdl' is your transport handler

	go func() {
		log.Println("gRPC Server running on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// 5. Start gRPC-Gateway (HTTP)
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Dial the gRPC server we just started above
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// 6. Start HTTP Server
	log.Println("HTTP Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
