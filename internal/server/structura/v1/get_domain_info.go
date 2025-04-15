package v1

import (
	"context"

	structurav1 "github.com/MaxFando/structura/api/grpc/gen/go/structura/v1"
)

func (s *StructuraServer) GetDomainInfo(ctx context.Context, req *structurav1.GetDomainInfoRequest) (*structurav1.GetDomainInfoResponse, error) {
	return &structurav1.GetDomainInfoResponse{
		PagesCount: 0,
		Domain:     req.GetDomain(),
		Pages:      nil,
	}, nil
}
