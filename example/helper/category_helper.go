package helper

import (
	"example/gen/category/v1"
	"example/repository"
)

func CategoryToProto(c *repository.Category) *categoryv1.Category {
	return &categoryv1.Category{
		Id:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
