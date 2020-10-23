package main

import (
	"github.com/WingkaiHo/go-grpc-project-template/app"
	"github.com/WingkaiHo/go-grpc-project-template/config"
)

func main() {
	app.StartApp(config.GetConfigByEnv())
}
