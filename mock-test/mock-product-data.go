package mocktest

import (
	pb "api_exam/genproto/product_exam"
	"context"
	"time"

	"github.com/google/uuid"
)

type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.Product, error)
	GetProductById(ctx context.Context, in *pb.GetProductByIdRequest) (*pb.Product, error)
	GetAllProduct(ctx context.Context, in *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
	UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.Product, error)
	DeleteProduct(ctx context.Context, in *pb.GetProductByIdRequest) error
}

type productServiceClient struct {
}

// CreateProduct implements ProductServiceClient.
func (p *productServiceClient) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.Product, error) {
	return &pb.Product{
		Id:          uuid.NewString(),
		OwnerId:     in.OwnerId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
		DeletedAt:   "",
	}, nil
}

// DeleteProduct implements ProductServiceClient.
func (p *productServiceClient) DeleteProduct(ctx context.Context, in *pb.GetProductByIdRequest) error {
	return nil
}

// GetAllProduct implements ProductServiceClient.
func (p *productServiceClient) GetAllProduct(ctx context.Context, in *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error) {
	product := pb.Product{
		Id:          "9d0a02e5-10a7-42f0-9fad-0bc91d2045d1",
		OwnerId:     "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		Name:        "Car",
		Description: "Chevrolet",
		Price:       54,
		CreatedAt:   "",
		UpdatedAt:   "",
		DeletedAt:   "",
	}
	return &pb.GetAllProductResponse{
		Products: []*pb.Product{
			&product,
			&product,
			&product},
	}, nil
}

// GetProductById implements ProductServiceClient.
func (p *productServiceClient) GetProductById(ctx context.Context, in *pb.GetProductByIdRequest) (*pb.Product, error) {
	return &pb.Product{
		Id:          in.ProductId,
		OwnerId:     "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		Name:        "Car",
		Description: "Chevrolet",
		Price:       54,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
		DeletedAt:   "",
	}, nil
}

// UpdateProduct implements ProductServiceClient.
func (p *productServiceClient) UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.Product, error) {
	return &pb.Product{
		Id:          in.Id,
		OwnerId:     in.OwnerId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
		DeletedAt:   "",
	}, nil
}

func NewProductServiceClient() ProductServiceClient {
	return &productServiceClient{}
}
