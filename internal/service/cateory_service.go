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

func (c *CategoryService) ListCategory(ctx context.Context, input *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	categories, err := c.CategoryDB.ListCategory()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category
	var total int32

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	total = int32(len(categoriesResponse))

	return &pb.ListCategoryResponse{
		Categories: categoriesResponse,
		Total:      total,
	}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	category, err := c.CategoryDB.GetCategory(input.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetCategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}
