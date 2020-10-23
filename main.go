package main

import (
	"log"

	"github.com/WingkaiHo/go-grpc-project-template/app"
	"github.com/WingkaiHo/go-grpc-project-template/config"
)

func main() {
	cfg := config.GetConfigByEnv()
	log.Println(cfg)
	app.StartApp(cfg)
}
