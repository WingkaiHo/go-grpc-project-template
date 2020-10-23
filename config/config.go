package config

// AppConfig is app config
type AppConfig struct {
	App  string
	Port int
	Env  string
}

// Config values
var Config = &AppConfig{
	App:  "go-grpc-project-template",
	Port: 55006,
	Env:  "development",
}
