package v1

import (
	"context"

	structurav1 "github.com/MaxFando/structura/api/grpc/gen/go/structura/v1"
)

func (s *StructuraServer) GetPageBlocks(ctx context.Context, req *structurav1.GetPageBlocksRequest) (*structurav1.GetPageBlocksResponse, error) {
	return &structurav1.GetPageBlocksResponse{
		Blocks: []*structurav1.Block{
			{
				Id:    "block-1",
				Index: 0,
				Type:  "text",
			},
			{
				Id:    "block-2",
				Index: 1,
				Type:  "image",
			},
		},
	}, nil
}
