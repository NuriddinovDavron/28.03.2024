package mocktest

import (
	pbu "api_exam/genproto/user_exam"
	"context"

	"google.golang.org/grpc"
)

type UserServiceClient interface {
	CreateUser(ctx context.Context, in pbu.CreateUserRequest, opts ...grpc.CallOption) (*pbu.UserApi, error)
	GetUserById(ctx context.Context, in *pbu.GetUserByIdRequest, opts ...grpc.CallOption) (*pbu.UserApi, error)
	GetAllUser(ctx context.Context, in *pbu.GetAllUserRequest, opts ...grpc.CallOption) (*pbu.GetAllUserResponse, error)
	UpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.UserApi, error)
	DeleteUser(ctx context.Context, in *pbu.GetUserByIdRequest, opts ...grpc.CallOption) error
	CheckField(ctx context.Context, in *pbu.CheckUser, opts ...grpc.CallOption) (*pbu.CheckRes, error)
	GetUserByEmail(ctx context.Context, in *pbu.EmailRequest, opts ...grpc.CallOption) (*pbu.UserApi, error)
	GetUserByRefreshToken(ctx context.Context, in *pbu.UserToken, opts ...grpc.CallOption) (*pbu.UserApi, error)
}

type userServiceClient struct {
}

// CheckField implements UserServiceClient.
func (u *userServiceClient) CheckField(ctx context.Context, in *pbu.CheckUser, opts ...grpc.CallOption) (*pbu.CheckRes, error) {
	return &pbu.CheckRes{
		Exists: false,
	}, nil
}

// CreateUser implements UserServiceClient.
func (u *userServiceClient) CreateUser(ctx context.Context, in pbu.CreateUserRequest, opts ...grpc.CallOption) (*pbu.UserApi, error) {
	return &pbu.UserApi{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron",
		LastName:  "Nuriddinov",
		Email:     "nuriddinovdavron2003@gmail.com",
	}, nil
}

// DeleteUser implements UserServiceClient.
func (u *userServiceClient) DeleteUser(ctx context.Context, in *pbu.GetUserByIdRequest, opts ...grpc.CallOption) error {
	return nil
}

// GetAllUser implements UserServiceClient.
func (u *userServiceClient) GetAllUser(ctx context.Context, in *pbu.GetAllUserRequest, opts ...grpc.CallOption) (*pbu.GetAllUserResponse, error) {
	user := pbu.User{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron",
		LastName:  "Nuriddinov",
		Email:     "nuriddinovdavron2003@gmail.com",
	}

	return &pbu.GetAllUserResponse{
		Users: []*pbu.User{
			&user,
			&user,
		},
	}, nil
}

// GetUserByEmail implements UserServiceClient.
func (u *userServiceClient) GetUserByEmail(ctx context.Context, in *pbu.EmailRequest, opts ...grpc.CallOption) (*pbu.UserApi, error) {
	return &pbu.UserApi{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron",
		LastName:  "Nuriddinov",
		Email:     "nuriddinovdavron2003@gmail.com",
	}, nil
}

// GetUserById implements UserServiceClient.
func (u *userServiceClient) GetUserById(ctx context.Context, in *pbu.GetUserByIdRequest, opts ...grpc.CallOption) (*pbu.UserApi, error) {
	return &pbu.UserApi{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron",
		LastName:  "Nuriddinov",
		Email:     "nuriddinovdavron2003@gmail.com",
	}, nil
}

// GetUserByRefreshToken implements UserServiceClient.
func (u *userServiceClient) GetUserByRefreshToken(ctx context.Context, in *pbu.UserToken, opts ...grpc.CallOption) (*pbu.UserApi, error) {
	return &pbu.UserApi{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron",
		LastName:  "Nuriddinov",
		Email:     "nuriddinovdavron2003@gmail.com",
	}, nil
}

// UpdateUser implements UserServiceClient.
func (u *userServiceClient) UpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.UserApi, error) {
	return &pbu.UserApi{
		Id:        "bef75c7c-39f1-4d27-81ba-a06b76c5dc6b",
		FirstName: "Davron Update",
		LastName:  "Nuriddinov Update",
		Email:     "nuriddinovdavron2003@gmail.com",
	}, nil
}

func NewUserServiceClient() UserServiceClient {
	return &userServiceClient{}
}
