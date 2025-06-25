package helper

import (
	repo "example/internal/repository"
	pb "example/pb/proto/product/v1"
)

func ProductToProto(c repo.Product) *pb.Product {
	return &pb.Product{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CategoryId:  c.CategoryID,
		Price:       c.Price,
	}
}
