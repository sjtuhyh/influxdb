// Code generated by protoc-gen-gogo.
// source: internal/meta.proto
// DO NOT EDIT!

/*
Package tsdb is a generated protocol buffer package.

It is generated from these files:
	internal/meta.proto

It has these top-level messages:
	Series
	Tag
	MeasurementFields
	Field
	MeasurementFieldSet
*/
package tsdb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Series struct {
	Key  string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Tags []*Tag `protobuf:"bytes,2,rep,name=Tags" json:"Tags,omitempty"`
}

func (m *Series) Reset()                    { *m = Series{} }
func (m *Series) String() string            { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()               {}
func (*Series) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{0} }

func (m *Series) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Series) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{1} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MeasurementFields struct {
	Name   string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Fields []*Field `protobuf:"bytes,2,rep,name=Fields" json:"Fields,omitempty"`
}

func (m *MeasurementFields) Reset()                    { *m = MeasurementFields{} }
func (m *MeasurementFields) String() string            { return proto.CompactTextString(m) }
func (*MeasurementFields) ProtoMessage()               {}
func (*MeasurementFields) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{2} }

func (m *MeasurementFields) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeasurementFields) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type int32  `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{3} }

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type MeasurementFieldSet struct {
	Measurements []*MeasurementFields `protobuf:"bytes,1,rep,name=Measurements" json:"Measurements,omitempty"`
}

func (m *MeasurementFieldSet) Reset()                    { *m = MeasurementFieldSet{} }
func (m *MeasurementFieldSet) String() string            { return proto.CompactTextString(m) }
func (*MeasurementFieldSet) ProtoMessage()               {}
func (*MeasurementFieldSet) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{4} }

func (m *MeasurementFieldSet) GetMeasurements() []*MeasurementFields {
	if m != nil {
		return m.Measurements
	}
	return nil
}

func init() {
	proto.RegisterType((*Series)(nil), "tsdb.Series")
	proto.RegisterType((*Tag)(nil), "tsdb.Tag")
	proto.RegisterType((*MeasurementFields)(nil), "tsdb.MeasurementFields")
	proto.RegisterType((*Field)(nil), "tsdb.Field")
	proto.RegisterType((*MeasurementFieldSet)(nil), "tsdb.MeasurementFieldSet")
}

func init() { proto.RegisterFile("internal/meta.proto", fileDescriptorMeta) }

var fileDescriptorMeta = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x2c, 0x1b, 0x72, 0xe9, 0xd0, 0x5e, 0x0a, 0xd5, 0x52, 0x08, 0xea, 0x92, 0xa5,
	0x0e, 0xb4, 0x53, 0xe9, 0xde, 0xa5, 0x3f, 0x06, 0x45, 0x74, 0xbf, 0x90, 0xc3, 0x18, 0x6c, 0x27,
	0x48, 0xca, 0x90, 0xff, 0xbe, 0xf8, 0xe4, 0xa1, 0x6d, 0xbc, 0x3d, 0x7d, 0xa7, 0xa7, 0x4f, 0x1c,
	0x2c, 0x9b, 0x3e, 0xb2, 0xef, 0xa9, 0xdd, 0x74, 0x1c, 0xa9, 0x3a, 0xfa, 0x43, 0x3c, 0xa0, 0x8a,
	0x61, 0xbf, 0x33, 0x2f, 0x50, 0x6e, 0xd9, 0x37, 0x1c, 0xf0, 0x1a, 0xf2, 0x77, 0x3e, 0xeb, 0x6c,
	0x95, 0xad, 0xe7, 0x76, 0x88, 0x78, 0x0f, 0xca, 0x51, 0x1d, 0xf4, 0x6c, 0x95, 0xaf, 0x17, 0x4f,
	0xf3, 0x6a, 0x28, 0x54, 0x8e, 0x6a, 0x2b, 0xd8, 0x3c, 0x42, 0xee, 0xa8, 0x9e, 0xe8, 0xdd, 0x42,
	0xf1, 0x4d, 0xed, 0x89, 0xf5, 0x4c, 0x58, 0x3a, 0x98, 0x0f, 0xb8, 0xf9, 0x64, 0x0a, 0x27, 0xcf,
	0x1d, 0xf7, 0xf1, 0xad, 0xe1, 0x76, 0x1f, 0x10, 0x41, 0x7d, 0x51, 0xc7, 0x63, 0x5b, 0x32, 0x3e,
	0x40, 0x99, 0xa6, 0xa3, 0x78, 0x91, 0xc4, 0xc2, 0xec, 0x38, 0x32, 0x1b, 0x28, 0x24, 0x4d, 0xbe,
	0x80, 0xa0, 0xdc, 0xf9, 0x98, 0xfc, 0x85, 0x95, 0x6c, 0x2c, 0x2c, 0xff, 0xeb, 0xb7, 0x1c, 0xf1,
	0x15, 0xae, 0x7e, 0xe1, 0xa0, 0x33, 0x51, 0xde, 0x25, 0xe5, 0xc5, 0x7f, 0xed, 0x9f, 0xcb, 0xbb,
	0x52, 0x36, 0xf9, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x31, 0x1f, 0xb9, 0x60, 0x01, 0x00,
	0x00,
}