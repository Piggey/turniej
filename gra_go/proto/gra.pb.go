// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: gra.proto

package proto

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

type KonfiguracjaGry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiczbaGraczy int32 `protobuf:"varint,1,opt,name=liczbaGraczy,proto3" json:"liczbaGraczy,omitempty"`
}

func (x *KonfiguracjaGry) Reset() {
	*x = KonfiguracjaGry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KonfiguracjaGry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KonfiguracjaGry) ProtoMessage() {}

func (x *KonfiguracjaGry) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KonfiguracjaGry.ProtoReflect.Descriptor instead.
func (*KonfiguracjaGry) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{0}
}

func (x *KonfiguracjaGry) GetLiczbaGraczy() int32 {
	if x != nil {
		return x.LiczbaGraczy
	}
	return 0
}

type NowaGraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID string `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	Opis  string `protobuf:"bytes,2,opt,name=opis,proto3" json:"opis,omitempty"`
}

func (x *NowaGraInfo) Reset() {
	*x = NowaGraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowaGraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowaGraInfo) ProtoMessage() {}

func (x *NowaGraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowaGraInfo.ProtoReflect.Descriptor instead.
func (*NowaGraInfo) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{1}
}

func (x *NowaGraInfo) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *NowaGraInfo) GetOpis() string {
	if x != nil {
		return x.Opis
	}
	return ""
}

type WizytowkaGracza struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nazwa string `protobuf:"bytes,1,opt,name=nazwa,proto3" json:"nazwa,omitempty"`
}

func (x *WizytowkaGracza) Reset() {
	*x = WizytowkaGracza{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WizytowkaGracza) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WizytowkaGracza) ProtoMessage() {}

func (x *WizytowkaGracza) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WizytowkaGracza.ProtoReflect.Descriptor instead.
func (*WizytowkaGracza) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{2}
}

func (x *WizytowkaGracza) GetNazwa() string {
	if x != nil {
		return x.Nazwa
	}
	return ""
}

type Dolaczanie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID     string           `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	Wizytowka *WizytowkaGracza `protobuf:"bytes,2,opt,name=wizytowka,proto3" json:"wizytowka,omitempty"`
}

func (x *Dolaczanie) Reset() {
	*x = Dolaczanie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dolaczanie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dolaczanie) ProtoMessage() {}

func (x *Dolaczanie) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dolaczanie.ProtoReflect.Descriptor instead.
func (*Dolaczanie) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{3}
}

func (x *Dolaczanie) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *Dolaczanie) GetWizytowka() *WizytowkaGracza {
	if x != nil {
		return x.Wizytowka
	}
	return nil
}

type StanGry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID             string `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	GraczID           string `protobuf:"bytes,2,opt,name=graczID,proto3" json:"graczID,omitempty"`
	SytuacjaNaPlanszy string `protobuf:"bytes,3,opt,name=sytuacjaNaPlanszy,proto3" json:"sytuacjaNaPlanszy,omitempty"`
	TwojeKarty        string `protobuf:"bytes,4,opt,name=twojeKarty,proto3" json:"twojeKarty,omitempty"`
}

func (x *StanGry) Reset() {
	*x = StanGry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StanGry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StanGry) ProtoMessage() {}

func (x *StanGry) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StanGry.ProtoReflect.Descriptor instead.
func (*StanGry) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{4}
}

func (x *StanGry) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *StanGry) GetGraczID() string {
	if x != nil {
		return x.GraczID
	}
	return ""
}

func (x *StanGry) GetSytuacjaNaPlanszy() string {
	if x != nil {
		return x.SytuacjaNaPlanszy
	}
	return ""
}

func (x *StanGry) GetTwojeKarty() string {
	if x != nil {
		return x.TwojeKarty
	}
	return ""
}

type RuchGracza struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GraID        string `protobuf:"bytes,1,opt,name=graID,proto3" json:"graID,omitempty"`
	GraczID      string `protobuf:"bytes,2,opt,name=graczID,proto3" json:"graczID,omitempty"`
	ZagranaKarta string `protobuf:"bytes,3,opt,name=zagranaKarta,proto3" json:"zagranaKarta,omitempty"`
}

func (x *RuchGracza) Reset() {
	*x = RuchGracza{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuchGracza) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuchGracza) ProtoMessage() {}

func (x *RuchGracza) ProtoReflect() protoreflect.Message {
	mi := &file_gra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuchGracza.ProtoReflect.Descriptor instead.
func (*RuchGracza) Descriptor() ([]byte, []int) {
	return file_gra_proto_rawDescGZIP(), []int{5}
}

func (x *RuchGracza) GetGraID() string {
	if x != nil {
		return x.GraID
	}
	return ""
}

func (x *RuchGracza) GetGraczID() string {
	if x != nil {
		return x.GraczID
	}
	return ""
}

func (x *RuchGracza) GetZagranaKarta() string {
	if x != nil {
		return x.ZagranaKarta
	}
	return ""
}

var File_gra_proto protoreflect.FileDescriptor

var file_gra_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0f, 0x4b,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x63, 0x6a, 0x61, 0x47, 0x72, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x69, 0x63, 0x7a, 0x62, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x69, 0x63, 0x7a, 0x62, 0x61, 0x47, 0x72, 0x61, 0x63,
	0x7a, 0x79, 0x22, 0x37, 0x0a, 0x0b, 0x4e, 0x6f, 0x77, 0x61, 0x47, 0x72, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x69, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x69, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x57,
	0x69, 0x7a, 0x79, 0x74, 0x6f, 0x77, 0x6b, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x7a, 0x77, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x61, 0x7a, 0x77, 0x61, 0x22, 0x52, 0x0a, 0x0a, 0x44, 0x6f, 0x6c, 0x61, 0x63, 0x7a, 0x61, 0x6e,
	0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x09, 0x77, 0x69, 0x7a, 0x79,
	0x74, 0x6f, 0x77, 0x6b, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x57, 0x69,
	0x7a, 0x79, 0x74, 0x6f, 0x77, 0x6b, 0x61, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x52, 0x09, 0x77,
	0x69, 0x7a, 0x79, 0x74, 0x6f, 0x77, 0x6b, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x61,
	0x6e, 0x47, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72,
	0x61, 0x63, 0x7a, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61,
	0x63, 0x7a, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x79, 0x74, 0x75, 0x61, 0x63, 0x6a, 0x61,
	0x4e, 0x61, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x7a, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x73, 0x79, 0x74, 0x75, 0x61, 0x63, 0x6a, 0x61, 0x4e, 0x61, 0x50, 0x6c, 0x61, 0x6e, 0x73,
	0x7a, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x77, 0x6f, 0x6a, 0x65, 0x4b, 0x61, 0x72, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x77, 0x6f, 0x6a, 0x65, 0x4b, 0x61, 0x72,
	0x74, 0x79, 0x22, 0x60, 0x0a, 0x0a, 0x52, 0x75, 0x63, 0x68, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x61, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x63, 0x7a, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x63, 0x7a, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x7a, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x61, 0x4b, 0x61, 0x72, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x7a, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x61, 0x4b,
	0x61, 0x72, 0x74, 0x61, 0x32, 0x7f, 0x0a, 0x03, 0x47, 0x72, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x4e,
	0x6f, 0x77, 0x79, 0x4d, 0x65, 0x63, 0x7a, 0x12, 0x10, 0x2e, 0x4b, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x63, 0x6a, 0x61, 0x47, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x4e, 0x6f, 0x77, 0x61,
	0x47, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0b, 0x44, 0x6f, 0x6c,
	0x61, 0x63, 0x7a, 0x44, 0x6f, 0x47, 0x72, 0x79, 0x12, 0x0b, 0x2e, 0x44, 0x6f, 0x6c, 0x61, 0x63,
	0x7a, 0x61, 0x6e, 0x69, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x47, 0x72, 0x79, 0x22,
	0x00, 0x12, 0x22, 0x0a, 0x07, 0x4d, 0x6f, 0x6a, 0x52, 0x75, 0x63, 0x68, 0x12, 0x0b, 0x2e, 0x52,
	0x75, 0x63, 0x68, 0x47, 0x72, 0x61, 0x63, 0x7a, 0x61, 0x1a, 0x08, 0x2e, 0x53, 0x74, 0x61, 0x6e,
	0x47, 0x72, 0x79, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x61, 0x72, 0x61, 0x7a, 0x2f, 0x74, 0x75, 0x72, 0x6e, 0x69,
	0x65, 0x6a, 0x2f, 0x67, 0x72, 0x61, 0x5f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gra_proto_rawDescOnce sync.Once
	file_gra_proto_rawDescData = file_gra_proto_rawDesc
)

func file_gra_proto_rawDescGZIP() []byte {
	file_gra_proto_rawDescOnce.Do(func() {
		file_gra_proto_rawDescData = protoimpl.X.CompressGZIP(file_gra_proto_rawDescData)
	})
	return file_gra_proto_rawDescData
}

var file_gra_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gra_proto_goTypes = []interface{}{
	(*KonfiguracjaGry)(nil), // 0: KonfiguracjaGry
	(*NowaGraInfo)(nil),     // 1: NowaGraInfo
	(*WizytowkaGracza)(nil), // 2: WizytowkaGracza
	(*Dolaczanie)(nil),      // 3: Dolaczanie
	(*StanGry)(nil),         // 4: StanGry
	(*RuchGracza)(nil),      // 5: RuchGracza
}
var file_gra_proto_depIdxs = []int32{
	2, // 0: Dolaczanie.wizytowka:type_name -> WizytowkaGracza
	0, // 1: Gra.NowyMecz:input_type -> KonfiguracjaGry
	3, // 2: Gra.DolaczDoGry:input_type -> Dolaczanie
	5, // 3: Gra.MojRuch:input_type -> RuchGracza
	1, // 4: Gra.NowyMecz:output_type -> NowaGraInfo
	4, // 5: Gra.DolaczDoGry:output_type -> StanGry
	4, // 6: Gra.MojRuch:output_type -> StanGry
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gra_proto_init() }
func file_gra_proto_init() {
	if File_gra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KonfiguracjaGry); i {
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
		file_gra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowaGraInfo); i {
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
		file_gra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WizytowkaGracza); i {
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
		file_gra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dolaczanie); i {
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
		file_gra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StanGry); i {
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
		file_gra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuchGracza); i {
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
			RawDescriptor: file_gra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gra_proto_goTypes,
		DependencyIndexes: file_gra_proto_depIdxs,
		MessageInfos:      file_gra_proto_msgTypes,
	}.Build()
	File_gra_proto = out.File
	file_gra_proto_rawDesc = nil
	file_gra_proto_goTypes = nil
	file_gra_proto_depIdxs = nil
}
