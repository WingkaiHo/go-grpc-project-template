package controller

import (
	"context"

	"github.com/WingkaiHo/go-grpc-project-template/api/proto/demo"
	"github.com/WingkaiHo/go-grpc-project-template/manager"
	"google.golang.org/grpc"
)

type detailsController struct {
}

// DetailsController 的实现
var DetailsController detailsController

// RegisterDetailsServer register detail server func
func RegisterDetailsServer(s *grpc.Server) {
	demo.RegisterDetailsServer(s, &DetailsController)
}

// Get detail information
func (detailController *detailsController) Get(ctx context.Context, in *demo.GetDetailRequest) (*demo.GetDetailResponse, error) {

	id := in.GetId()
	_, name := manager.DetailsManager.Get(id)

	var response demo.GetDetailResponse

	response.Name = name
	response.Id = id

	return &response, nil
}
