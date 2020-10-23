package app

import (
	"fmt"
	"log"
	"net"

	"github.com/WingkaiHo/go-grpc-project-template/config"
	"github.com/WingkaiHo/go-grpc-project-template/controller"
	"google.golang.org/grpc"
)

// StartApp start grpc server
func StartApp(cfg config.AppConfig) {
	server := grpc.NewServer()
	controller.RegisterAllServer(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server is start at %v\n", cfg.Port)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
