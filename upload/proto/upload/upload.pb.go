// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/upload/upload.proto

package fullbottle_srv_upload

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type UploadStatus int32

const (
	UploadStatus__         UploadStatus = 0
	UploadStatus_uploading UploadStatus = 1
	// 2, 3 wouldn't appear in normal case
	UploadStatus_manifest  UploadStatus = 2
	UploadStatus_weed_done UploadStatus = 3
	UploadStatus_finish    UploadStatus = 4
)

var UploadStatus_name = map[int32]string{
	0: "_",
	1: "uploading",
	2: "manifest",
	3: "weed_done",
	4: "finish",
}

var UploadStatus_value = map[string]int32{
	"_":         0,
	"uploading": 1,
	"manifest":  2,
	"weed_done": 3,
	"finish":    4,
}

func (x UploadStatus) String() string {
	return proto.EnumName(UploadStatus_name, int32(x))
}

func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{0}
}

type GenerateUploadTokenRequest struct {
	OwnerId              int64    `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	FolderId             int64    `protobuf:"varint,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Mime                 string   `protobuf:"bytes,6,opt,name=mime,proto3" json:"mime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateUploadTokenRequest) Reset()         { *m = GenerateUploadTokenRequest{} }
func (m *GenerateUploadTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateUploadTokenRequest) ProtoMessage()    {}
func (*GenerateUploadTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{0}
}

func (m *GenerateUploadTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateUploadTokenRequest.Unmarshal(m, b)
}
func (m *GenerateUploadTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateUploadTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateUploadTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateUploadTokenRequest.Merge(m, src)
}
func (m *GenerateUploadTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateUploadTokenRequest.Size(m)
}
func (m *GenerateUploadTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateUploadTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateUploadTokenRequest proto.InternalMessageInfo

func (m *GenerateUploadTokenRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *GenerateUploadTokenRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *GenerateUploadTokenRequest) GetFolderId() int64 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

func (m *GenerateUploadTokenRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *GenerateUploadTokenRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GenerateUploadTokenRequest) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

type GenerateUploadTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	NeedUpload           bool     `protobuf:"varint,2,opt,name=need_upload,json=needUpload,proto3" json:"need_upload,omitempty"`
	Uploaded             []int64  `protobuf:"varint,3,rep,packed,name=uploaded,proto3" json:"uploaded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateUploadTokenResponse) Reset()         { *m = GenerateUploadTokenResponse{} }
func (m *GenerateUploadTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateUploadTokenResponse) ProtoMessage()    {}
func (*GenerateUploadTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{1}
}

func (m *GenerateUploadTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateUploadTokenResponse.Unmarshal(m, b)
}
func (m *GenerateUploadTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateUploadTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateUploadTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateUploadTokenResponse.Merge(m, src)
}
func (m *GenerateUploadTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateUploadTokenResponse.Size(m)
}
func (m *GenerateUploadTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateUploadTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateUploadTokenResponse proto.InternalMessageInfo

func (m *GenerateUploadTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GenerateUploadTokenResponse) GetNeedUpload() bool {
	if m != nil {
		return m.NeedUpload
	}
	return false
}

func (m *GenerateUploadTokenResponse) GetUploaded() []int64 {
	if m != nil {
		return m.Uploaded
	}
	return nil
}

type UploadFileRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Raw                  []byte   `protobuf:"bytes,4,opt,name=raw,proto3" json:"raw,omitempty"`
	ChunkHash            string   `protobuf:"bytes,5,opt,name=chunk_hash,json=chunkHash,proto3" json:"chunk_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{2}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

func (m *UploadFileRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UploadFileRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *UploadFileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UploadFileRequest) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *UploadFileRequest) GetChunkHash() string {
	if m != nil {
		return m.ChunkHash
	}
	return ""
}

type UploadFileResponse struct {
	Status               UploadStatus `protobuf:"varint,1,opt,name=status,proto3,enum=fullbottle.srv.upload.UploadStatus" json:"status,omitempty"`
	Uploaded             []int64      `protobuf:"varint,2,rep,packed,name=uploaded,proto3" json:"uploaded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{3}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetStatus() UploadStatus {
	if m != nil {
		return m.Status
	}
	return UploadStatus__
}

func (m *UploadFileResponse) GetUploaded() []int64 {
	if m != nil {
		return m.Uploaded
	}
	return nil
}

type GetFileUploadedChunksRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileUploadedChunksRequest) Reset()         { *m = GetFileUploadedChunksRequest{} }
func (m *GetFileUploadedChunksRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileUploadedChunksRequest) ProtoMessage()    {}
func (*GetFileUploadedChunksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{4}
}

func (m *GetFileUploadedChunksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileUploadedChunksRequest.Unmarshal(m, b)
}
func (m *GetFileUploadedChunksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileUploadedChunksRequest.Marshal(b, m, deterministic)
}
func (m *GetFileUploadedChunksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileUploadedChunksRequest.Merge(m, src)
}
func (m *GetFileUploadedChunksRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileUploadedChunksRequest.Size(m)
}
func (m *GetFileUploadedChunksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileUploadedChunksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileUploadedChunksRequest proto.InternalMessageInfo

func (m *GetFileUploadedChunksRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetFileUploadedChunksRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

type GetFileUploadedChunksResponse struct {
	Uploaded             []int64  `protobuf:"varint,1,rep,packed,name=uploaded,proto3" json:"uploaded,omitempty"`
	NeedUpload           bool     `protobuf:"varint,2,opt,name=need_upload,json=needUpload,proto3" json:"need_upload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileUploadedChunksResponse) Reset()         { *m = GetFileUploadedChunksResponse{} }
func (m *GetFileUploadedChunksResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileUploadedChunksResponse) ProtoMessage()    {}
func (*GetFileUploadedChunksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{5}
}

func (m *GetFileUploadedChunksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileUploadedChunksResponse.Unmarshal(m, b)
}
func (m *GetFileUploadedChunksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileUploadedChunksResponse.Marshal(b, m, deterministic)
}
func (m *GetFileUploadedChunksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileUploadedChunksResponse.Merge(m, src)
}
func (m *GetFileUploadedChunksResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileUploadedChunksResponse.Size(m)
}
func (m *GetFileUploadedChunksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileUploadedChunksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileUploadedChunksResponse proto.InternalMessageInfo

func (m *GetFileUploadedChunksResponse) GetUploaded() []int64 {
	if m != nil {
		return m.Uploaded
	}
	return nil
}

func (m *GetFileUploadedChunksResponse) GetNeedUpload() bool {
	if m != nil {
		return m.NeedUpload
	}
	return false
}

type CancelFileUploadRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	OwnerId              int64    `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelFileUploadRequest) Reset()         { *m = CancelFileUploadRequest{} }
func (m *CancelFileUploadRequest) String() string { return proto.CompactTextString(m) }
func (*CancelFileUploadRequest) ProtoMessage()    {}
func (*CancelFileUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{6}
}

func (m *CancelFileUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelFileUploadRequest.Unmarshal(m, b)
}
func (m *CancelFileUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelFileUploadRequest.Marshal(b, m, deterministic)
}
func (m *CancelFileUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelFileUploadRequest.Merge(m, src)
}
func (m *CancelFileUploadRequest) XXX_Size() int {
	return xxx_messageInfo_CancelFileUploadRequest.Size(m)
}
func (m *CancelFileUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelFileUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelFileUploadRequest proto.InternalMessageInfo

func (m *CancelFileUploadRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CancelFileUploadRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

type CancelFileUploadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelFileUploadResponse) Reset()         { *m = CancelFileUploadResponse{} }
func (m *CancelFileUploadResponse) String() string { return proto.CompactTextString(m) }
func (*CancelFileUploadResponse) ProtoMessage()    {}
func (*CancelFileUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5508b024e4d0ced2, []int{7}
}

func (m *CancelFileUploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelFileUploadResponse.Unmarshal(m, b)
}
func (m *CancelFileUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelFileUploadResponse.Marshal(b, m, deterministic)
}
func (m *CancelFileUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelFileUploadResponse.Merge(m, src)
}
func (m *CancelFileUploadResponse) XXX_Size() int {
	return xxx_messageInfo_CancelFileUploadResponse.Size(m)
}
func (m *CancelFileUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelFileUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelFileUploadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fullbottle.srv.upload.UploadStatus", UploadStatus_name, UploadStatus_value)
	proto.RegisterType((*GenerateUploadTokenRequest)(nil), "fullbottle.srv.upload.GenerateUploadTokenRequest")
	proto.RegisterType((*GenerateUploadTokenResponse)(nil), "fullbottle.srv.upload.GenerateUploadTokenResponse")
	proto.RegisterType((*UploadFileRequest)(nil), "fullbottle.srv.upload.UploadFileRequest")
	proto.RegisterType((*UploadFileResponse)(nil), "fullbottle.srv.upload.UploadFileResponse")
	proto.RegisterType((*GetFileUploadedChunksRequest)(nil), "fullbottle.srv.upload.GetFileUploadedChunksRequest")
	proto.RegisterType((*GetFileUploadedChunksResponse)(nil), "fullbottle.srv.upload.GetFileUploadedChunksResponse")
	proto.RegisterType((*CancelFileUploadRequest)(nil), "fullbottle.srv.upload.CancelFileUploadRequest")
	proto.RegisterType((*CancelFileUploadResponse)(nil), "fullbottle.srv.upload.CancelFileUploadResponse")
}

func init() { proto.RegisterFile("proto/upload/upload.proto", fileDescriptor_5508b024e4d0ced2) }

var fileDescriptor_5508b024e4d0ced2 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xeb, 0x38, 0x31, 0xc9, 0x10, 0x90, 0x19, 0x28, 0xb8, 0x2e, 0x15, 0x95, 0xb9, 0x04,
	0x0e, 0xae, 0x48, 0xb9, 0x71, 0xac, 0x44, 0x29, 0x12, 0x42, 0x32, 0xf4, 0x86, 0x14, 0xb9, 0xf1,
	0x98, 0x58, 0x75, 0x76, 0x83, 0x77, 0x4d, 0x24, 0x0e, 0x9c, 0xb8, 0xf2, 0x26, 0xbc, 0x19, 0x2f,
	0x81, 0x76, 0xd7, 0x49, 0x9d, 0x62, 0x43, 0xdb, 0x53, 0x76, 0xfe, 0x3b, 0x1f, 0xbf, 0x99, 0x1d,
	0x07, 0x76, 0x16, 0x05, 0x97, 0xfc, 0xa0, 0x5c, 0xe4, 0x3c, 0x4e, 0xaa, 0x9f, 0x50, 0x6b, 0xb8,
	0x9d, 0x96, 0x79, 0x7e, 0xc6, 0xa5, 0xcc, 0x29, 0x14, 0xc5, 0xd7, 0xd0, 0x5c, 0x06, 0xbf, 0x2c,
	0xf0, 0x8f, 0x89, 0x51, 0x11, 0x4b, 0x3a, 0xd5, 0xd2, 0x47, 0x7e, 0x4e, 0x2c, 0xa2, 0x2f, 0x25,
	0x09, 0x89, 0x3b, 0xd0, 0xe7, 0x4b, 0x46, 0xc5, 0x24, 0x4b, 0x3c, 0x6b, 0xdf, 0x1a, 0xd9, 0xd1,
	0x2d, 0x6d, 0x9f, 0x24, 0xe8, 0x43, 0x3f, 0xcd, 0x72, 0x62, 0xf1, 0x9c, 0xbc, 0xce, 0xbe, 0x35,
	0x1a, 0x44, 0x6b, 0x1b, 0x77, 0x61, 0x90, 0xf2, 0x3c, 0x31, 0x71, 0xb6, 0x8e, 0xeb, 0x1b, 0xe1,
	0x24, 0x41, 0x84, 0xee, 0x2c, 0x16, 0x33, 0xaf, 0xab, 0x83, 0xf4, 0x59, 0x69, 0x22, 0xfb, 0x46,
	0x5e, 0x4f, 0xfb, 0xea, 0xb3, 0xd2, 0xe6, 0xd9, 0x9c, 0x3c, 0xc7, 0xf8, 0xa9, 0x73, 0xb0, 0x80,
	0xdd, 0x46, 0x5a, 0xb1, 0xe0, 0x4c, 0x10, 0x3e, 0x80, 0x9e, 0x54, 0x82, 0x66, 0x1d, 0x44, 0xc6,
	0xc0, 0x27, 0x70, 0x9b, 0x11, 0x25, 0x13, 0xd3, 0xb2, 0x86, 0xed, 0x47, 0xa0, 0x24, 0x93, 0x43,
	0xb5, 0x62, 0xee, 0x48, 0xd1, 0xda, 0x8a, 0x76, 0x65, 0x07, 0x3f, 0x2d, 0xb8, 0x67, 0xdc, 0x5e,
	0x67, 0x39, 0xad, 0xe6, 0xd2, 0x5c, 0xa8, 0x3e, 0xad, 0xce, 0xe6, 0xb4, 0x1e, 0x82, 0xc3, 0xd3,
	0x54, 0x90, 0xac, 0xc6, 0x51, 0x59, 0xe8, 0x82, 0x5d, 0xc4, 0x4b, 0x3d, 0x8b, 0x61, 0xa4, 0x8e,
	0xb8, 0x07, 0x30, 0x9d, 0x95, 0xec, 0x7c, 0xa2, 0x87, 0xd4, 0xd3, 0xf9, 0x07, 0x5a, 0x79, 0x13,
	0x8b, 0x59, 0x30, 0x07, 0xac, 0xe3, 0x54, 0x8d, 0xbf, 0x02, 0x47, 0xc8, 0x58, 0x96, 0x42, 0x03,
	0xdd, 0x1d, 0x3f, 0x0d, 0x1b, 0x9f, 0x3b, 0x34, 0xa1, 0x1f, 0xb4, 0x6b, 0x54, 0x85, 0x6c, 0xb4,
	0xdf, 0xb9, 0xd4, 0xfe, 0x7b, 0x78, 0x7c, 0x4c, 0x52, 0xd5, 0x3a, 0xad, 0xa4, 0x23, 0x85, 0x22,
	0x6e, 0x3a, 0x88, 0xe0, 0x13, 0xec, 0xb5, 0x24, 0xac, 0x5a, 0xa9, 0xd3, 0x58, 0x9b, 0x34, 0xff,
	0x7d, 0xc9, 0xe0, 0x2d, 0x3c, 0x3a, 0x8a, 0xd9, 0x94, 0xf2, 0x8b, 0x02, 0x37, 0x26, 0xf5, 0xc1,
	0xfb, 0x3b, 0x97, 0x81, 0x7c, 0xfe, 0x0e, 0x86, 0xf5, 0x51, 0x62, 0x0f, 0xac, 0x89, 0xbb, 0x85,
	0x77, 0x60, 0x60, 0xd0, 0x32, 0xf6, 0xd9, 0xb5, 0x70, 0x08, 0xfd, 0x79, 0xcc, 0xb2, 0x94, 0x84,
	0x74, 0x3b, 0xea, 0x72, 0xa9, 0xe0, 0x13, 0xce, 0xc8, 0xb5, 0x11, 0xc0, 0x49, 0x33, 0x96, 0x89,
	0x99, 0xdb, 0x1d, 0xff, 0xb6, 0xc1, 0xa9, 0x76, 0xf1, 0x3b, 0xdc, 0x6f, 0xd8, 0x70, 0x7c, 0xd1,
	0xf2, 0xa0, 0xed, 0xdf, 0xae, 0x3f, 0xbe, 0x4e, 0x88, 0xe9, 0x2b, 0xd8, 0xc2, 0x29, 0xc0, 0xc5,
	0x7e, 0xe1, 0xe8, 0x9f, 0x7b, 0x54, 0xfb, 0x22, 0xfc, 0x67, 0x57, 0xf0, 0x5c, 0x17, 0xf9, 0x61,
	0xc1, 0x76, 0xe3, 0x16, 0xe0, 0x61, 0x2b, 0x74, 0xfb, 0x12, 0xfa, 0x2f, 0xaf, 0x17, 0xb4, 0xc6,
	0x28, 0xc1, 0xbd, 0xfc, 0xc2, 0x18, 0xb6, 0xe4, 0x6a, 0x59, 0x2b, 0xff, 0xe0, 0xca, 0xfe, 0xab,
	0xb2, 0x67, 0x8e, 0xfe, 0x47, 0x3e, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x07, 0x34, 0x1f, 0xac,
	0xae, 0x05, 0x00, 0x00,
}
