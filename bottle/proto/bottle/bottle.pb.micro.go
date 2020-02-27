// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bottle/bottle.proto

package fullbottle_srv_bottle

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

// Client API for BottleService service

type BottleService interface {
	GetBottleMetadata(ctx context.Context, in *GetBottleMetadataRequest, opts ...client.CallOption) (*GetBottleMetadataResponse, error)
	InitBottle(ctx context.Context, in *InitBottleRequest, opts ...client.CallOption) (*InitBottleResponse, error)
	UpdateBottle(ctx context.Context, in *UpdateBottleRequest, opts ...client.CallOption) (*UpdateBottleResponse, error)
	GetFolderInfo(ctx context.Context, in *GetFolderInfoRequest, opts ...client.CallOption) (*GetFolderInfoResponse, error)
	CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...client.CallOption) (*CreateFolderResponse, error)
	UpdateFolder(ctx context.Context, in *UpdateFolderRequest, opts ...client.CallOption) (*UpdateFolderResponse, error)
	RemoveFolder(ctx context.Context, in *RemoveFolderRequest, opts ...client.CallOption) (*RemoveFolderResponse, error)
	GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...client.CallOption) (*GetFileInfoResponse, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...client.CallOption) (*UpdateFileResponse, error)
	RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...client.CallOption) (*RemoveFileResponse, error)
	GenerateUploadToken(ctx context.Context, in *GenerateUploadTokenRequest, opts ...client.CallOption) (*GenerateUploadTokenResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...client.CallOption) (*UploadFileResponse, error)
}

type bottleService struct {
	c    client.Client
	name string
}

func NewBottleService(name string, c client.Client) BottleService {
	return &bottleService{
		c:    c,
		name: name,
	}
}

func (c *bottleService) GetBottleMetadata(ctx context.Context, in *GetBottleMetadataRequest, opts ...client.CallOption) (*GetBottleMetadataResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.GetBottleMetadata", in)
	out := new(GetBottleMetadataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) InitBottle(ctx context.Context, in *InitBottleRequest, opts ...client.CallOption) (*InitBottleResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.InitBottle", in)
	out := new(InitBottleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) UpdateBottle(ctx context.Context, in *UpdateBottleRequest, opts ...client.CallOption) (*UpdateBottleResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.UpdateBottle", in)
	out := new(UpdateBottleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) GetFolderInfo(ctx context.Context, in *GetFolderInfoRequest, opts ...client.CallOption) (*GetFolderInfoResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.GetFolderInfo", in)
	out := new(GetFolderInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...client.CallOption) (*CreateFolderResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.CreateFolder", in)
	out := new(CreateFolderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) UpdateFolder(ctx context.Context, in *UpdateFolderRequest, opts ...client.CallOption) (*UpdateFolderResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.UpdateFolder", in)
	out := new(UpdateFolderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) RemoveFolder(ctx context.Context, in *RemoveFolderRequest, opts ...client.CallOption) (*RemoveFolderResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.RemoveFolder", in)
	out := new(RemoveFolderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...client.CallOption) (*GetFileInfoResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.GetFileInfo", in)
	out := new(GetFileInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...client.CallOption) (*UpdateFileResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.UpdateFile", in)
	out := new(UpdateFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...client.CallOption) (*RemoveFileResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.RemoveFile", in)
	out := new(RemoveFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) GenerateUploadToken(ctx context.Context, in *GenerateUploadTokenRequest, opts ...client.CallOption) (*GenerateUploadTokenResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.GenerateUploadToken", in)
	out := new(GenerateUploadTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bottleService) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...client.CallOption) (*UploadFileResponse, error) {
	req := c.c.NewRequest(c.name, "BottleService.UploadFile", in)
	out := new(UploadFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BottleService service

type BottleServiceHandler interface {
	GetBottleMetadata(context.Context, *GetBottleMetadataRequest, *GetBottleMetadataResponse) error
	InitBottle(context.Context, *InitBottleRequest, *InitBottleResponse) error
	UpdateBottle(context.Context, *UpdateBottleRequest, *UpdateBottleResponse) error
	GetFolderInfo(context.Context, *GetFolderInfoRequest, *GetFolderInfoResponse) error
	CreateFolder(context.Context, *CreateFolderRequest, *CreateFolderResponse) error
	UpdateFolder(context.Context, *UpdateFolderRequest, *UpdateFolderResponse) error
	RemoveFolder(context.Context, *RemoveFolderRequest, *RemoveFolderResponse) error
	GetFileInfo(context.Context, *GetFileInfoRequest, *GetFileInfoResponse) error
	UpdateFile(context.Context, *UpdateFileRequest, *UpdateFileResponse) error
	RemoveFile(context.Context, *RemoveFileRequest, *RemoveFileResponse) error
	GenerateUploadToken(context.Context, *GenerateUploadTokenRequest, *GenerateUploadTokenResponse) error
	UploadFile(context.Context, *UploadFileRequest, *UploadFileResponse) error
}

func RegisterBottleServiceHandler(s server.Server, hdlr BottleServiceHandler, opts ...server.HandlerOption) error {
	type bottleService interface {
		GetBottleMetadata(ctx context.Context, in *GetBottleMetadataRequest, out *GetBottleMetadataResponse) error
		InitBottle(ctx context.Context, in *InitBottleRequest, out *InitBottleResponse) error
		UpdateBottle(ctx context.Context, in *UpdateBottleRequest, out *UpdateBottleResponse) error
		GetFolderInfo(ctx context.Context, in *GetFolderInfoRequest, out *GetFolderInfoResponse) error
		CreateFolder(ctx context.Context, in *CreateFolderRequest, out *CreateFolderResponse) error
		UpdateFolder(ctx context.Context, in *UpdateFolderRequest, out *UpdateFolderResponse) error
		RemoveFolder(ctx context.Context, in *RemoveFolderRequest, out *RemoveFolderResponse) error
		GetFileInfo(ctx context.Context, in *GetFileInfoRequest, out *GetFileInfoResponse) error
		UpdateFile(ctx context.Context, in *UpdateFileRequest, out *UpdateFileResponse) error
		RemoveFile(ctx context.Context, in *RemoveFileRequest, out *RemoveFileResponse) error
		GenerateUploadToken(ctx context.Context, in *GenerateUploadTokenRequest, out *GenerateUploadTokenResponse) error
		UploadFile(ctx context.Context, in *UploadFileRequest, out *UploadFileResponse) error
	}
	type BottleService struct {
		bottleService
	}
	h := &bottleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BottleService{h}, opts...))
}

type bottleServiceHandler struct {
	BottleServiceHandler
}

func (h *bottleServiceHandler) GetBottleMetadata(ctx context.Context, in *GetBottleMetadataRequest, out *GetBottleMetadataResponse) error {
	return h.BottleServiceHandler.GetBottleMetadata(ctx, in, out)
}

func (h *bottleServiceHandler) InitBottle(ctx context.Context, in *InitBottleRequest, out *InitBottleResponse) error {
	return h.BottleServiceHandler.InitBottle(ctx, in, out)
}

func (h *bottleServiceHandler) UpdateBottle(ctx context.Context, in *UpdateBottleRequest, out *UpdateBottleResponse) error {
	return h.BottleServiceHandler.UpdateBottle(ctx, in, out)
}

func (h *bottleServiceHandler) GetFolderInfo(ctx context.Context, in *GetFolderInfoRequest, out *GetFolderInfoResponse) error {
	return h.BottleServiceHandler.GetFolderInfo(ctx, in, out)
}

func (h *bottleServiceHandler) CreateFolder(ctx context.Context, in *CreateFolderRequest, out *CreateFolderResponse) error {
	return h.BottleServiceHandler.CreateFolder(ctx, in, out)
}

func (h *bottleServiceHandler) UpdateFolder(ctx context.Context, in *UpdateFolderRequest, out *UpdateFolderResponse) error {
	return h.BottleServiceHandler.UpdateFolder(ctx, in, out)
}

func (h *bottleServiceHandler) RemoveFolder(ctx context.Context, in *RemoveFolderRequest, out *RemoveFolderResponse) error {
	return h.BottleServiceHandler.RemoveFolder(ctx, in, out)
}

func (h *bottleServiceHandler) GetFileInfo(ctx context.Context, in *GetFileInfoRequest, out *GetFileInfoResponse) error {
	return h.BottleServiceHandler.GetFileInfo(ctx, in, out)
}

func (h *bottleServiceHandler) UpdateFile(ctx context.Context, in *UpdateFileRequest, out *UpdateFileResponse) error {
	return h.BottleServiceHandler.UpdateFile(ctx, in, out)
}

func (h *bottleServiceHandler) RemoveFile(ctx context.Context, in *RemoveFileRequest, out *RemoveFileResponse) error {
	return h.BottleServiceHandler.RemoveFile(ctx, in, out)
}

func (h *bottleServiceHandler) GenerateUploadToken(ctx context.Context, in *GenerateUploadTokenRequest, out *GenerateUploadTokenResponse) error {
	return h.BottleServiceHandler.GenerateUploadToken(ctx, in, out)
}

func (h *bottleServiceHandler) UploadFile(ctx context.Context, in *UploadFileRequest, out *UploadFileResponse) error {
	return h.BottleServiceHandler.UploadFile(ctx, in, out)
}
