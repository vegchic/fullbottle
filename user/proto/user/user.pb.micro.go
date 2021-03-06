// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package fullbottle_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	GetUserInfo(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
	ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...client.CallOption) (*ModifyUserResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, opts ...client.CallOption) (*GetUserAvatarResponse, error)
	UploadUserAvatar(ctx context.Context, in *UploadUserAvatarRequest, opts ...client.CallOption) (*UploadUserAvatarResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetUserInfo(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUserInfo", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...client.CallOption) (*ModifyUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.ModifyUser", in)
	out := new(ModifyUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserLogin", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, opts ...client.CallOption) (*GetUserAvatarResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUserAvatar", in)
	out := new(GetUserAvatarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UploadUserAvatar(ctx context.Context, in *UploadUserAvatarRequest, opts ...client.CallOption) (*UploadUserAvatarResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UploadUserAvatar", in)
	out := new(UploadUserAvatarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	GetUserInfo(context.Context, *GetUserRequest, *GetUserResponse) error
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
	ModifyUser(context.Context, *ModifyUserRequest, *ModifyUserResponse) error
	UserLogin(context.Context, *UserLoginRequest, *UserLoginResponse) error
	GetUserAvatar(context.Context, *GetUserAvatarRequest, *GetUserAvatarResponse) error
	UploadUserAvatar(context.Context, *UploadUserAvatarRequest, *UploadUserAvatarResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		GetUserInfo(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error
		ModifyUser(ctx context.Context, in *ModifyUserRequest, out *ModifyUserResponse) error
		UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, out *GetUserAvatarResponse) error
		UploadUserAvatar(ctx context.Context, in *UploadUserAvatarRequest, out *UploadUserAvatarResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) GetUserInfo(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserServiceHandler.GetUserInfo(ctx, in, out)
}

func (h *userServiceHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *userServiceHandler) ModifyUser(ctx context.Context, in *ModifyUserRequest, out *ModifyUserResponse) error {
	return h.UserServiceHandler.ModifyUser(ctx, in, out)
}

func (h *userServiceHandler) UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserServiceHandler.UserLogin(ctx, in, out)
}

func (h *userServiceHandler) GetUserAvatar(ctx context.Context, in *GetUserAvatarRequest, out *GetUserAvatarResponse) error {
	return h.UserServiceHandler.GetUserAvatar(ctx, in, out)
}

func (h *userServiceHandler) UploadUserAvatar(ctx context.Context, in *UploadUserAvatarRequest, out *UploadUserAvatarResponse) error {
	return h.UserServiceHandler.UploadUserAvatar(ctx, in, out)
}
