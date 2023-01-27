// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_MESSAGE_TYPE_UNKNOWN                MessageType = 0
	MessageType_MESSAGE_TYPE_UPLOAD_FILE_REQUEST    MessageType = 1
	MessageType_MESSAGE_TYPE_UPLOAD_FILE_RESPONSE   MessageType = 2
	MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_REQUEST  MessageType = 3
	MessageType_MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE MessageType = 4
	MessageType_MESSAGE_TYPE_LIST_FILES_REQUEST     MessageType = 5
	MessageType_MESSAGE_TYPE_LIST_FILES_RESPONSE    MessageType = 6
	MessageType_MESSAGE_TYPE_DELETE_FILE_REQUEST    MessageType = 7
	MessageType_MESSAGE_TYPE_DELETE_FILE_RESPONSE   MessageType = 8
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MESSAGE_TYPE_UNKNOWN",
		1: "MESSAGE_TYPE_UPLOAD_FILE_REQUEST",
		2: "MESSAGE_TYPE_UPLOAD_FILE_RESPONSE",
		3: "MESSAGE_TYPE_DOWNLOAD_FILE_REQUEST",
		4: "MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE",
		5: "MESSAGE_TYPE_LIST_FILES_REQUEST",
		6: "MESSAGE_TYPE_LIST_FILES_RESPONSE",
		7: "MESSAGE_TYPE_DELETE_FILE_REQUEST",
		8: "MESSAGE_TYPE_DELETE_FILE_RESPONSE",
	}
	MessageType_value = map[string]int32{
		"MESSAGE_TYPE_UNKNOWN":                0,
		"MESSAGE_TYPE_UPLOAD_FILE_REQUEST":    1,
		"MESSAGE_TYPE_UPLOAD_FILE_RESPONSE":   2,
		"MESSAGE_TYPE_DOWNLOAD_FILE_REQUEST":  3,
		"MESSAGE_TYPE_DOWNLOAD_FILE_RESPONSE": 4,
		"MESSAGE_TYPE_LIST_FILES_REQUEST":     5,
		"MESSAGE_TYPE_LIST_FILES_RESPONSE":    6,
		"MESSAGE_TYPE_DELETE_FILE_REQUEST":    7,
		"MESSAGE_TYPE_DELETE_FILE_RESPONSE":   8,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type UploadFileResponseCode int32

const (
	UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY   UploadFileResponseCode = 0
	UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_SUCCESS UploadFileResponseCode = 1
	UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_FAILED  UploadFileResponseCode = 2
	UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_EXISTS  UploadFileResponseCode = 3
)

// Enum value maps for UploadFileResponseCode.
var (
	UploadFileResponseCode_name = map[int32]string{
		0: "UPLOAD_FILE_RESPONSE_CODE_READY",
		1: "UPLOAD_FILE_RESPONSE_CODE_SUCCESS",
		2: "UPLOAD_FILE_RESPONSE_CODE_FAILED",
		3: "UPLOAD_FILE_RESPONSE_CODE_EXISTS",
	}
	UploadFileResponseCode_value = map[string]int32{
		"UPLOAD_FILE_RESPONSE_CODE_READY":   0,
		"UPLOAD_FILE_RESPONSE_CODE_SUCCESS": 1,
		"UPLOAD_FILE_RESPONSE_CODE_FAILED":  2,
		"UPLOAD_FILE_RESPONSE_CODE_EXISTS":  3,
	}
)

func (x UploadFileResponseCode) Enum() *UploadFileResponseCode {
	p := new(UploadFileResponseCode)
	*p = x
	return p
}

func (x UploadFileResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadFileResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (UploadFileResponseCode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x UploadFileResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadFileResponseCode.Descriptor instead.
func (UploadFileResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type DownLoadFileResponseCode int32

const (
	DownLoadFileResponseCode_DOWNLOAD_FILE_RESPONSE_CODE_NOT_FOUND DownLoadFileResponseCode = 0
	DownLoadFileResponseCode_DOWNLOAD_FILE_RESPONSE_CODE_OK        DownLoadFileResponseCode = 1
)

// Enum value maps for DownLoadFileResponseCode.
var (
	DownLoadFileResponseCode_name = map[int32]string{
		0: "DOWNLOAD_FILE_RESPONSE_CODE_NOT_FOUND",
		1: "DOWNLOAD_FILE_RESPONSE_CODE_OK",
	}
	DownLoadFileResponseCode_value = map[string]int32{
		"DOWNLOAD_FILE_RESPONSE_CODE_NOT_FOUND": 0,
		"DOWNLOAD_FILE_RESPONSE_CODE_OK":        1,
	}
)

func (x DownLoadFileResponseCode) Enum() *DownLoadFileResponseCode {
	p := new(DownLoadFileResponseCode)
	*p = x
	return p
}

func (x DownLoadFileResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownLoadFileResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[2].Descriptor()
}

func (DownLoadFileResponseCode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[2]
}

func (x DownLoadFileResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownLoadFileResponseCode.Descriptor instead.
func (DownLoadFileResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type DeleteFileResponseCode int32

const (
	DeleteFileResponseCode_DELETE_FILE_RESPONSE_CODE_NOT_FOUND DeleteFileResponseCode = 0
	DeleteFileResponseCode_DELETE_FILE_RESPONSE_CODE_OK        DeleteFileResponseCode = 1
)

// Enum value maps for DeleteFileResponseCode.
var (
	DeleteFileResponseCode_name = map[int32]string{
		0: "DELETE_FILE_RESPONSE_CODE_NOT_FOUND",
		1: "DELETE_FILE_RESPONSE_CODE_OK",
	}
	DeleteFileResponseCode_value = map[string]int32{
		"DELETE_FILE_RESPONSE_CODE_NOT_FOUND": 0,
		"DELETE_FILE_RESPONSE_CODE_OK":        1,
	}
)

func (x DeleteFileResponseCode) Enum() *DeleteFileResponseCode {
	p := new(DeleteFileResponseCode)
	*p = x
	return p
}

func (x DeleteFileResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteFileResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[3].Descriptor()
}

func (DeleteFileResponseCode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[3]
}

func (x DeleteFileResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteFileResponseCode.Descriptor instead.
func (DeleteFileResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdCode string `protobuf:"bytes,1,opt,name=id_code,json=idCode,proto3" json:"id_code,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size   uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetIdCode() string {
	if x != nil {
		return x.IdCode
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize string `protobuf:"bytes,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRequest) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId       string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	ResponseCode UploadFileResponseCode `protobuf:"varint,2,opt,name=response_code,json=responseCode,proto3,enum=proto.UploadFileResponseCode" json:"response_code,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UploadFileResponse) GetResponseCode() UploadFileResponseCode {
	if x != nil {
		return x.ResponseCode
	}
	return UploadFileResponseCode_UPLOAD_FILE_RESPONSE_CODE_READY
}

type DownLoadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *DownLoadFileRequest) Reset() {
	*x = DownLoadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownLoadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownLoadFileRequest) ProtoMessage() {}

func (x *DownLoadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownLoadFileRequest.ProtoReflect.Descriptor instead.
func (*DownLoadFileRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *DownLoadFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownLoadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize string `protobuf:"bytes,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
}

func (x *DownLoadFileResponse) Reset() {
	*x = DownLoadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownLoadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownLoadFileResponse) ProtoMessage() {}

func (x *DownLoadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownLoadFileResponse.ProtoReflect.Descriptor instead.
func (*DownLoadFileResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *DownLoadFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DownLoadFileResponse) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseCode *DeleteFileResponse `protobuf:"bytes,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFileResponse) GetResponseCode() *DeleteFileResponse {
	if x != nil {
		return x.ResponseCode
	}
	return nil
}

type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenKey string `protobuf:"bytes,1,opt,name=token_key,json=tokenKey,proto3" json:"token_key,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListFilesRequest) GetTokenKey() string {
	if x != nil {
		return x.TokenKey
	}
	return ""
}

type ListFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalFiles uint32 `protobuf:"varint,1,opt,name=total_files,json=totalFiles,proto3" json:"total_files,omitempty"`
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListFilesResponse) GetTotalFiles() uint32 {
	if x != nil {
		return x.TotalFiles
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x71, 0x0a, 0x12, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a,
	0x13, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x50, 0x0a,
	0x14, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x4b, 0x65, 0x79, 0x22, 0x34, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2a, 0xdd, 0x02, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41,
	0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10,
	0x02, 0x12, 0x26, 0x0a, 0x22, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x49, 0x4c,
	0x45, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x06, 0x12, 0x24, 0x0a,
	0x20, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x08, 0x2a, 0xb0, 0x01, 0x0a, 0x16, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x55, 0x50,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x50, 0x4c, 0x4f, 0x41,
	0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x2a, 0x69, 0x0a,
	0x18, 0x44, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x44, 0x4f, 0x57,
	0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44,
	0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x2a, 0x63, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c,
	0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x42, 0x4f, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x05, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_goTypes = []interface{}{
	(MessageType)(0),              // 0: proto.MessageType
	(UploadFileResponseCode)(0),   // 1: proto.UploadFileResponseCode
	(DownLoadFileResponseCode)(0), // 2: proto.DownLoadFileResponseCode
	(DeleteFileResponseCode)(0),   // 3: proto.DeleteFileResponseCode
	(*File)(nil),                  // 4: proto.File
	(*UploadFileRequest)(nil),     // 5: proto.UploadFileRequest
	(*UploadFileResponse)(nil),    // 6: proto.UploadFileResponse
	(*DownLoadFileRequest)(nil),   // 7: proto.DownLoadFileRequest
	(*DownLoadFileResponse)(nil),  // 8: proto.DownLoadFileResponse
	(*DeleteFileRequest)(nil),     // 9: proto.DeleteFileRequest
	(*DeleteFileResponse)(nil),    // 10: proto.DeleteFileResponse
	(*ListFilesRequest)(nil),      // 11: proto.ListFilesRequest
	(*ListFilesResponse)(nil),     // 12: proto.ListFilesResponse
}
var file_api_proto_depIdxs = []int32{
	1,  // 0: proto.UploadFileResponse.response_code:type_name -> proto.UploadFileResponseCode
	10, // 1: proto.DeleteFileResponse.response_code:type_name -> proto.DeleteFileResponse
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownLoadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownLoadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
