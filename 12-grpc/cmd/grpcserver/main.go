package main

import (
	"database/sql"
	"github.com/souluanf/fullcycle-grpc/internal/database"
	"github.com/souluanf/fullcycle-grpc/internal/pb"
	"github.com/souluanf/fullcycle-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, error := sql.Open("sqlite3", "db.sqlite3")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
