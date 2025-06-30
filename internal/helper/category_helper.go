package helper

import (
	repo "example/internal/repository"
	pb "example/pb/proto/category/v1"
)

func CategoryToProto(c repo.Category) *pb.Category {
	return &pb.Category{
		Id:   c.ID,
		Name: c.Name,
	}
}
