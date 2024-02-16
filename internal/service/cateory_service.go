package service

import (
	"context"

	"github.com/lyracampos/go-gRPC/internal/database"
	"github.com/lyracampos/go-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	category, err := c.CategoryDB.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}
