package helper

import (
	pb "example/pb/proto/v1"
	"example/models"
)

func CategoryToProto(c *models.Category) *pb.Category {
	return &pb.Category{
		Id:          c.ID,
		Name:        c.Name,
	}
}
