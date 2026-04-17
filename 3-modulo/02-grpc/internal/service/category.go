package service

import (
	"context"
	"io"

	"github.com/danielfmpc/pos-go-grpc/internal/database"
	"github.com/danielfmpc/pos-go-grpc/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{CategoryDb: categoryDb}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Create(in.Name, in.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryListResponse, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list categories")
	}
	var categoriesResponse []*pb.Category
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryListResponse{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Find(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get category")
	}
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := []*pb.Category{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CategoryListResponse{Categories: categories})
		}
		if err != nil {
			return err
		}
		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories = append(categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBi(stream pb.CategoryService_CreateCategoryStreamBiServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}

func (c *CategoryService) ListCategoriesStream(in *pb.Blank, stream pb.CategoryService_ListCategoriesStreamServer) error {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return status.Error(codes.Internal, "failed to get category")
	}
	for _, category := range categories {
		err = stream.Send(&pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
