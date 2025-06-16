package helper

import (
	"example/models"
	pb "example/pb/proto/category/v1"
)

func CategoryToProto(c *models.Category) *pb.Category {
	return &pb.Category{
		Id:   c.ID,
		Name: c.Name,
	}
}
