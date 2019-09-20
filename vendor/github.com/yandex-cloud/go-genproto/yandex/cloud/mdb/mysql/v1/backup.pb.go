// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1/backup.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
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

// A MySQL backup. For more information, see
// the [documentation](/docs/managed-mysql/concepts/backup).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the MySQL cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9951bfdbe6ddc67, []int{0}
}

func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (m *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(m, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.mysql.v1.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1/backup.proto", fileDescriptor_a9951bfdbe6ddc67)
}

var fileDescriptor_a9951bfdbe6ddc67 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0x50, 0x2a, 0x62, 0x06, 0x44, 0xc4, 0x10, 0x82, 0x10, 0x15, 0x03, 0xaa, 0x90,
	0x6a, 0x2b, 0x30, 0x21, 0xa6, 0xa6, 0x62, 0x60, 0x8d, 0x98, 0x58, 0x22, 0xff, 0x6b, 0xb0, 0xb0,
	0xe3, 0x92, 0x9c, 0x23, 0xfa, 0x84, 0x3c, 0x05, 0x3b, 0x8f, 0x81, 0xb0, 0xd3, 0x81, 0x01, 0xb1,
	0xd9, 0x77, 0xdf, 0xfd, 0x3e, 0xe9, 0x0e, 0x5d, 0x6d, 0x69, 0x2b, 0xe4, 0x3b, 0xe1, 0xda, 0x3a,
	0x41, 0x8c, 0x60, 0xc4, 0x6c, 0xfb, 0x37, 0x4d, 0x86, 0x82, 0x30, 0xca, 0x5f, 0xdd, 0x06, 0x6f,
	0x3a, 0x0b, 0x36, 0x3d, 0x0d, 0x1c, 0xf6, 0x1c, 0x36, 0x82, 0x61, 0xcf, 0xe1, 0xa1, 0xc8, 0xcf,
	0x7f, 0x45, 0x0c, 0x54, 0x2b, 0x41, 0x41, 0xd9, 0x36, 0x4c, 0xe6, 0x17, 0x8d, 0xb5, 0x8d, 0x96,
	0xc4, 0xff, 0x98, 0x5b, 0x13, 0x50, 0x46, 0xf6, 0x40, 0xcd, 0x18, 0x7d, 0xf9, 0x19, 0xa1, 0x69,
	0xe9, 0x5d, 0xe9, 0x09, 0x8a, 0x95, 0xc8, 0xa2, 0x59, 0x34, 0x4f, 0xca, 0xc9, 0xd7, 0x47, 0x11,
	0x55, 0xb1, 0x12, 0xe9, 0x19, 0x4a, 0xd6, 0x56, 0x0b, 0xd9, 0xd5, 0x4a, 0x64, 0xf1, 0x4f, 0xb3,
	0x3a, 0x08, 0x85, 0x47, 0x91, 0xde, 0x21, 0xc4, 0x3b, 0x49, 0x41, 0x8a, 0x9a, 0x42, 0xb6, 0x37,
	0x8b, 0xe6, 0x87, 0x37, 0x39, 0x0e, 0x4e, 0xbc, 0x73, 0xe2, 0xa7, 0x9d, 0xb3, 0x4a, 0x46, 0x7a,
	0x09, 0xe9, 0x35, 0x3a, 0xee, 0xad, 0xeb, 0xb8, 0xac, 0xb9, 0x76, 0x3d, 0x84, 0xfc, 0x89, 0xcf,
	0x3f, 0x0a, 0x8d, 0x55, 0xa8, 0x07, 0x4d, 0x0f, 0xb4, 0x1b, 0x35, 0xfb, 0xff, 0x6b, 0x46, 0x7a,
	0x09, 0xe5, 0xc3, 0xf3, 0xaa, 0x51, 0xf0, 0xe2, 0x18, 0xe6, 0xd6, 0x90, 0xb0, 0xac, 0x45, 0x58,
	0x56, 0x63, 0x17, 0x8d, 0x6c, 0xfd, 0x38, 0xf9, 0xf3, 0x10, 0xf7, 0xfe, 0xc1, 0xa6, 0x1e, 0xbb,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x78, 0xbb, 0x89, 0x97, 0xb2, 0x01, 0x00, 0x00,
}