package controller

import "google.golang.org/grpc"

// RegisterAllServer register all controller
func RegisterAllServer(s *grpc.Server) {
	RegisterDetailsServer(s)
}
