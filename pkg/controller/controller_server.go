package controller

import (
	api "sharepoint-connector/pkg/api"
	sidecar "sharepoint-connector/sidecar"
)

import (
	"context"
)

type Server struct {
	sidecar.UnimplementedServiceServer
}

func (s *Server) Sidecar(ctx context.Context, request *sidecar.SidecarRequest) (*sidecar.ResponseDto, error) {

	result, err := api.Sidecar(request.RequestDto, ctx)

	return result, err
}
