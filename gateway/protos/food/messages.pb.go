// Code generated by protoc-gen-go. DO NOT EDIT.
// source: food/messages.proto

package food

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

type FoodGenre struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodGenre) Reset()         { *m = FoodGenre{} }
func (m *FoodGenre) String() string { return proto.CompactTextString(m) }
func (*FoodGenre) ProtoMessage()    {}
func (*FoodGenre) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{0}
}

func (m *FoodGenre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGenre.Unmarshal(m, b)
}
func (m *FoodGenre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGenre.Marshal(b, m, deterministic)
}
func (m *FoodGenre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGenre.Merge(m, src)
}
func (m *FoodGenre) XXX_Size() int {
	return xxx_messageInfo_FoodGenre.Size(m)
}
func (m *FoodGenre) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGenre.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGenre proto.InternalMessageInfo

func (m *FoodGenre) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FoodGenre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FoodRecord struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *FoodData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodRecord) Reset()         { *m = FoodRecord{} }
func (m *FoodRecord) String() string { return proto.CompactTextString(m) }
func (*FoodRecord) ProtoMessage()    {}
func (*FoodRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{1}
}

func (m *FoodRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodRecord.Unmarshal(m, b)
}
func (m *FoodRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodRecord.Marshal(b, m, deterministic)
}
func (m *FoodRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodRecord.Merge(m, src)
}
func (m *FoodRecord) XXX_Size() int {
	return xxx_messageInfo_FoodRecord.Size(m)
}
func (m *FoodRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodRecord.DiscardUnknown(m)
}

var xxx_messageInfo_FoodRecord proto.InternalMessageInfo

func (m *FoodRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FoodRecord) GetData() *FoodData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FoodData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                uint64   `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	GenreId              uint64   `protobuf:"varint,3,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodData) Reset()         { *m = FoodData{} }
func (m *FoodData) String() string { return proto.CompactTextString(m) }
func (*FoodData) ProtoMessage()    {}
func (*FoodData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{2}
}

func (m *FoodData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodData.Unmarshal(m, b)
}
func (m *FoodData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodData.Marshal(b, m, deterministic)
}
func (m *FoodData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodData.Merge(m, src)
}
func (m *FoodData) XXX_Size() int {
	return xxx_messageInfo_FoodData.Size(m)
}
func (m *FoodData) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodData.DiscardUnknown(m)
}

var xxx_messageInfo_FoodData proto.InternalMessageInfo

func (m *FoodData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FoodData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *FoodData) GetGenreId() uint64 {
	if m != nil {
		return m.GenreId
	}
	return 0
}

type FoodCreateQuery struct {
	Data                 *FoodData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodCreateQuery) Reset()         { *m = FoodCreateQuery{} }
func (m *FoodCreateQuery) String() string { return proto.CompactTextString(m) }
func (*FoodCreateQuery) ProtoMessage()    {}
func (*FoodCreateQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{3}
}

func (m *FoodCreateQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodCreateQuery.Unmarshal(m, b)
}
func (m *FoodCreateQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodCreateQuery.Marshal(b, m, deterministic)
}
func (m *FoodCreateQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodCreateQuery.Merge(m, src)
}
func (m *FoodCreateQuery) XXX_Size() int {
	return xxx_messageInfo_FoodCreateQuery.Size(m)
}
func (m *FoodCreateQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodCreateQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodCreateQuery proto.InternalMessageInfo

func (m *FoodCreateQuery) GetData() *FoodData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FoodCreateResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodCreateResult) Reset()         { *m = FoodCreateResult{} }
func (m *FoodCreateResult) String() string { return proto.CompactTextString(m) }
func (*FoodCreateResult) ProtoMessage()    {}
func (*FoodCreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{4}
}

func (m *FoodCreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodCreateResult.Unmarshal(m, b)
}
func (m *FoodCreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodCreateResult.Marshal(b, m, deterministic)
}
func (m *FoodCreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodCreateResult.Merge(m, src)
}
func (m *FoodCreateResult) XXX_Size() int {
	return xxx_messageInfo_FoodCreateResult.Size(m)
}
func (m *FoodCreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodCreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodCreateResult proto.InternalMessageInfo

func (m *FoodCreateResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type FoodDeleteQuery struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodDeleteQuery) Reset()         { *m = FoodDeleteQuery{} }
func (m *FoodDeleteQuery) String() string { return proto.CompactTextString(m) }
func (*FoodDeleteQuery) ProtoMessage()    {}
func (*FoodDeleteQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{5}
}

func (m *FoodDeleteQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodDeleteQuery.Unmarshal(m, b)
}
func (m *FoodDeleteQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodDeleteQuery.Marshal(b, m, deterministic)
}
func (m *FoodDeleteQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodDeleteQuery.Merge(m, src)
}
func (m *FoodDeleteQuery) XXX_Size() int {
	return xxx_messageInfo_FoodDeleteQuery.Size(m)
}
func (m *FoodDeleteQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodDeleteQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodDeleteQuery proto.InternalMessageInfo

func (m *FoodDeleteQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FoodDeleteResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodDeleteResult) Reset()         { *m = FoodDeleteResult{} }
func (m *FoodDeleteResult) String() string { return proto.CompactTextString(m) }
func (*FoodDeleteResult) ProtoMessage()    {}
func (*FoodDeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{6}
}

func (m *FoodDeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodDeleteResult.Unmarshal(m, b)
}
func (m *FoodDeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodDeleteResult.Marshal(b, m, deterministic)
}
func (m *FoodDeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodDeleteResult.Merge(m, src)
}
func (m *FoodDeleteResult) XXX_Size() int {
	return xxx_messageInfo_FoodDeleteResult.Size(m)
}
func (m *FoodDeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodDeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodDeleteResult proto.InternalMessageInfo

func (m *FoodDeleteResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type FoodGetQuery struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodGetQuery) Reset()         { *m = FoodGetQuery{} }
func (m *FoodGetQuery) String() string { return proto.CompactTextString(m) }
func (*FoodGetQuery) ProtoMessage()    {}
func (*FoodGetQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{7}
}

func (m *FoodGetQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGetQuery.Unmarshal(m, b)
}
func (m *FoodGetQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGetQuery.Marshal(b, m, deterministic)
}
func (m *FoodGetQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGetQuery.Merge(m, src)
}
func (m *FoodGetQuery) XXX_Size() int {
	return xxx_messageInfo_FoodGetQuery.Size(m)
}
func (m *FoodGetQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGetQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGetQuery proto.InternalMessageInfo

func (m *FoodGetQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FoodGetResult struct {
	Record               *FoodRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FoodGetResult) Reset()         { *m = FoodGetResult{} }
func (m *FoodGetResult) String() string { return proto.CompactTextString(m) }
func (*FoodGetResult) ProtoMessage()    {}
func (*FoodGetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{8}
}

func (m *FoodGetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGetResult.Unmarshal(m, b)
}
func (m *FoodGetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGetResult.Marshal(b, m, deterministic)
}
func (m *FoodGetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGetResult.Merge(m, src)
}
func (m *FoodGetResult) XXX_Size() int {
	return xxx_messageInfo_FoodGetResult.Size(m)
}
func (m *FoodGetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGetResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGetResult proto.InternalMessageInfo

func (m *FoodGetResult) GetRecord() *FoodRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type FoodSearchQuery struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GenreId              uint64   `protobuf:"varint,4,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodSearchQuery) Reset()         { *m = FoodSearchQuery{} }
func (m *FoodSearchQuery) String() string { return proto.CompactTextString(m) }
func (*FoodSearchQuery) ProtoMessage()    {}
func (*FoodSearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{9}
}

func (m *FoodSearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodSearchQuery.Unmarshal(m, b)
}
func (m *FoodSearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodSearchQuery.Marshal(b, m, deterministic)
}
func (m *FoodSearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodSearchQuery.Merge(m, src)
}
func (m *FoodSearchQuery) XXX_Size() int {
	return xxx_messageInfo_FoodSearchQuery.Size(m)
}
func (m *FoodSearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodSearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodSearchQuery proto.InternalMessageInfo

func (m *FoodSearchQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FoodSearchQuery) GetGenreId() uint64 {
	if m != nil {
		return m.GenreId
	}
	return 0
}

type FoodSearchResult struct {
	Records              []*FoodRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FoodSearchResult) Reset()         { *m = FoodSearchResult{} }
func (m *FoodSearchResult) String() string { return proto.CompactTextString(m) }
func (*FoodSearchResult) ProtoMessage()    {}
func (*FoodSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{10}
}

func (m *FoodSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodSearchResult.Unmarshal(m, b)
}
func (m *FoodSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodSearchResult.Marshal(b, m, deterministic)
}
func (m *FoodSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodSearchResult.Merge(m, src)
}
func (m *FoodSearchResult) XXX_Size() int {
	return xxx_messageInfo_FoodSearchResult.Size(m)
}
func (m *FoodSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodSearchResult proto.InternalMessageInfo

func (m *FoodSearchResult) GetRecords() []*FoodRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type FoodEmptyQuery struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodEmptyQuery) Reset()         { *m = FoodEmptyQuery{} }
func (m *FoodEmptyQuery) String() string { return proto.CompactTextString(m) }
func (*FoodEmptyQuery) ProtoMessage()    {}
func (*FoodEmptyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{11}
}

func (m *FoodEmptyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodEmptyQuery.Unmarshal(m, b)
}
func (m *FoodEmptyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodEmptyQuery.Marshal(b, m, deterministic)
}
func (m *FoodEmptyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodEmptyQuery.Merge(m, src)
}
func (m *FoodEmptyQuery) XXX_Size() int {
	return xxx_messageInfo_FoodEmptyQuery.Size(m)
}
func (m *FoodEmptyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodEmptyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodEmptyQuery proto.InternalMessageInfo

type FoodGetGenresResult struct {
	Genres               []*FoodGenre `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FoodGetGenresResult) Reset()         { *m = FoodGetGenresResult{} }
func (m *FoodGetGenresResult) String() string { return proto.CompactTextString(m) }
func (*FoodGetGenresResult) ProtoMessage()    {}
func (*FoodGetGenresResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{12}
}

func (m *FoodGetGenresResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGetGenresResult.Unmarshal(m, b)
}
func (m *FoodGetGenresResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGetGenresResult.Marshal(b, m, deterministic)
}
func (m *FoodGetGenresResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGetGenresResult.Merge(m, src)
}
func (m *FoodGetGenresResult) XXX_Size() int {
	return xxx_messageInfo_FoodGetGenresResult.Size(m)
}
func (m *FoodGetGenresResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGetGenresResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGetGenresResult proto.InternalMessageInfo

func (m *FoodGetGenresResult) GetGenres() []*FoodGenre {
	if m != nil {
		return m.Genres
	}
	return nil
}

func init() {
	proto.RegisterType((*FoodGenre)(nil), "food.FoodGenre")
	proto.RegisterType((*FoodRecord)(nil), "food.FoodRecord")
	proto.RegisterType((*FoodData)(nil), "food.FoodData")
	proto.RegisterType((*FoodCreateQuery)(nil), "food.FoodCreateQuery")
	proto.RegisterType((*FoodCreateResult)(nil), "food.FoodCreateResult")
	proto.RegisterType((*FoodDeleteQuery)(nil), "food.FoodDeleteQuery")
	proto.RegisterType((*FoodDeleteResult)(nil), "food.FoodDeleteResult")
	proto.RegisterType((*FoodGetQuery)(nil), "food.FoodGetQuery")
	proto.RegisterType((*FoodGetResult)(nil), "food.FoodGetResult")
	proto.RegisterType((*FoodSearchQuery)(nil), "food.FoodSearchQuery")
	proto.RegisterType((*FoodSearchResult)(nil), "food.FoodSearchResult")
	proto.RegisterType((*FoodEmptyQuery)(nil), "food.FoodEmptyQuery")
	proto.RegisterType((*FoodGetGenresResult)(nil), "food.FoodGetGenresResult")
}

func init() { proto.RegisterFile("food/messages.proto", fileDescriptor_b52ab5d0a6e96024) }

var fileDescriptor_b52ab5d0a6e96024 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0xeb, 0xd3, 0x40,
	0x10, 0xc7, 0x49, 0x1b, 0xfb, 0x18, 0xb5, 0x2d, 0xa9, 0x87, 0x78, 0x91, 0xba, 0x17, 0x8b, 0xd8,
	0x04, 0x14, 0x91, 0x5e, 0x4a, 0xd1, 0xaa, 0x78, 0x12, 0xd7, 0x9b, 0x17, 0xd9, 0x66, 0xc7, 0x34,
	0x90, 0x64, 0xc3, 0xee, 0xa6, 0xd2, 0xff, 0x5e, 0xf6, 0x91, 0xc6, 0x47, 0xf1, 0x77, 0xcb, 0xbc,
	0x3e, 0xdf, 0xf9, 0x4e, 0x16, 0x96, 0x3f, 0x84, 0xe0, 0x69, 0x85, 0x4a, 0xb1, 0x1c, 0x55, 0xd2,
	0x48, 0xa1, 0x45, 0x14, 0x9a, 0x24, 0x49, 0x61, 0xfa, 0x41, 0x08, 0xfe, 0x11, 0x6b, 0x89, 0xd1,
	0x0c, 0x06, 0x05, 0x8f, 0x83, 0x55, 0xb0, 0x0e, 0xe9, 0xa0, 0xe0, 0x51, 0x04, 0x61, 0xcd, 0x2a,
	0x8c, 0x07, 0xab, 0x60, 0x3d, 0xa5, 0xf6, 0x9b, 0xec, 0x01, 0xcc, 0x00, 0xc5, 0x4c, 0x48, 0xfe,
	0xcf, 0x04, 0x81, 0x90, 0x33, 0xcd, 0xec, 0xc4, 0xfd, 0x97, 0xb3, 0xc4, 0x68, 0x24, 0xa6, 0xff,
	0xc0, 0x34, 0xa3, 0xb6, 0x46, 0x3e, 0xc3, 0xa4, 0xcb, 0x5c, 0x15, 0x82, 0x5e, 0x21, 0x7a, 0x04,
	0xf7, 0x1a, 0x59, 0x64, 0x4e, 0x36, 0xa4, 0x2e, 0x88, 0x1e, 0xc3, 0x24, 0x37, 0x4b, 0x7e, 0x2f,
	0x78, 0x3c, 0xb4, 0x85, 0xb1, 0x8d, 0x3f, 0x71, 0xf2, 0x1a, 0xe6, 0x06, 0xf8, 0x4e, 0x22, 0xd3,
	0xf8, 0xa5, 0x45, 0x79, 0xb9, 0xee, 0x11, 0xfc, 0x67, 0x8f, 0x17, 0xb0, 0xe8, 0xc7, 0x28, 0xaa,
	0xb6, 0xd4, 0x51, 0x0c, 0x63, 0xd5, 0x66, 0x19, 0x2a, 0x65, 0x47, 0x27, 0xb4, 0x0b, 0xc9, 0x53,
	0x27, 0x72, 0xc0, 0x12, 0x3b, 0x91, 0xbf, 0xcc, 0x77, 0x40, 0xd7, 0x72, 0x27, 0xf0, 0x09, 0x3c,
	0x70, 0x97, 0xd7, 0xb7, 0x69, 0x5b, 0x78, 0xe8, 0xeb, 0x1e, 0xb5, 0x86, 0x91, 0xb4, 0x57, 0xf7,
	0xae, 0x16, 0xbd, 0x2b, 0xf7, 0x37, 0xa8, 0xaf, 0x93, 0xbd, 0xdb, 0xf5, 0x2b, 0x32, 0x99, 0x9d,
	0x1c, 0xfd, 0xd6, 0xa1, 0x7f, 0x3f, 0x69, 0xf8, 0xe7, 0x49, 0x77, 0xce, 0x8a, 0x23, 0x78, 0xfd,
	0xe7, 0x30, 0x76, 0x7c, 0x63, 0x65, 0x78, 0x73, 0x81, 0xae, 0x81, 0x2c, 0x60, 0x66, 0xd2, 0xef,
	0xab, 0x46, 0x5f, 0xec, 0x02, 0x64, 0x07, 0x4b, 0x6f, 0xc7, 0xbe, 0x35, 0xe5, 0xa1, 0xcf, 0x60,
	0x64, 0x35, 0x3b, 0xe6, 0xbc, 0x67, 0xda, 0x3e, 0xea, 0xcb, 0x6f, 0xb7, 0xdf, 0xde, 0xe4, 0x85,
	0x3e, 0xb5, 0xc7, 0x24, 0x13, 0x55, 0xda, 0x48, 0x81, 0xe5, 0x51, 0xd7, 0xa9, 0xca, 0x4e, 0x42,
	0x94, 0x1b, 0x3c, 0xe3, 0xa6, 0x66, 0xe7, 0x22, 0xcd, 0x99, 0xc6, 0x9f, 0xec, 0x92, 0xda, 0x37,
	0xae, 0x52, 0x03, 0x3b, 0x8e, 0x6c, 0xf0, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb0,
	0x1e, 0x86, 0x07, 0x03, 0x00, 0x00,
}
