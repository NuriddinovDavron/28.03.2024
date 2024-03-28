package services

import (
	"fmt"

	"api_exam/config"
	pbp "api_exam/genproto/product_exam"
	pbu "api_exam/genproto/user_exam"
	mock "api_exam/mock-test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	MockUserService() mock.UserServiceClient
	ProductService() pbp.ProductServiceClient
	MockProductService() mock.ProductServiceClient
}

type serviceManager struct {
	userService        pbu.UserServiceClient
	mockUserService    mock.UserServiceClient
	productService     pbp.ProductServiceClient
	mockProductService mock.ProductServiceClient
}

// MockProductService implements IServiceManager.
func (s *serviceManager) MockProductService() mock.ProductServiceClient {
	return s.mockProductService
}

// MockUserService implements IServiceManager.
func (s *serviceManager) MockUserService() mock.UserServiceClient {
	return s.mockUserService
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProductServiceHost, conf.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pbu.NewUserServiceClient(connUser),
		mockUserService: mock.NewUserServiceClient(),
		productService: pbp.NewProductServiceClient(connPost),
		mockProductService: mock.NewProductServiceClient(),
	}

	return serviceManager, nil
}
