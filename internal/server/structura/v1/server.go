package v1

import structurav1 "github.com/MaxFando/structura/api/grpc/gen/go/structura/v1"

type StructuraServer struct {
	structurav1.StructuraServiceServer
}

func NewStructuraServer() *StructuraServer {
	return &StructuraServer{}
}
