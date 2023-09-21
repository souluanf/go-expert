package service

import (
	"context"
	"github.com/souluanf/fullcycle-grpc/internal/database"
	"github.com/souluanf/fullcycle-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoryResponses []*pb.Category
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryList{Categories: categoryResponses}, nil
}
