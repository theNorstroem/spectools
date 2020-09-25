// because we need the yaml, this file is overwritten by hand :-(
package specSpec

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/theNorstroem/spectools/pkg/orderedmap"
	furo "github.com/theNorstroem/spectools/pkg/specSpec/furo"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Defines a service
type Service struct {
	// Describe the rpcs or so
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	// The version number, use semver
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" yaml:"version"`
	// Describe the rpcs or so
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" yaml:"description"`
	//
	Lifecycle *Lifecycle `protobuf:"bytes,4,opt,name=lifecycle,proto3" json:"lifecycle" yaml:"lifecycle"`

	// information for the proto generator, should be removed for the client spec
	XProto *Typeproto `protobuf:"bytes,5,opt,name=__proto,json=Proto,proto3" json:"__proto" yaml:"__proto"`

	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RPCs for the service ==> will contains *Rpc
	Services   *orderedmap.OrderedMap `protobuf:"bytes,6,rep,name=services,proto3" json:"services" yaml:"services" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Extensions *orderedmap.OrderedMap `yaml:"extensions"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetXProto() *Typeproto {
	if x != nil {
		return x.XProto
	}
	return nil
}

func (x *Service) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Service) GetLifecycle() *Lifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetServices() *orderedmap.OrderedMap {
	if x != nil {
		return x.Services
	}
	return orderedmap.New()
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Repuest and response types for services, used in service.type
type Servicereqres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define the request type, leave this field empty if not needed
	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Define the response type, leave this field empty if not needed
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	// Define the body field in request the type
	// The name of the request field whose value is mapped to the HTTP request
	// body, or `*` for mapping all request fields not captured by the path
	// pattern to the HTTP body, or omitted for not having any HTTP request body.
	//
	// NOTE: the referred field must be present at the top-level of the request
	// message type.
	BodyField string `protobuf:"bytes,3,opt,name=response,proto3" json:"bodyfield" yaml:"bodyfield"`
}

func (x *Servicereqres) Reset() {
	*x = Servicereqres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Servicereqres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servicereqres) ProtoMessage() {}

func (x *Servicereqres) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Servicereqres.ProtoReflect.Descriptor instead.
func (*Servicereqres) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{1}
}

func (x *Servicereqres) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *Servicereqres) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// Lifecycle for a service
type Lifecycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Is this version deprecated
	Deprecated bool `protobuf:"varint,1,opt,name=deprecated,proto3" json:"deprecated"`
	// Inform about the replacement here, if you have one
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
}

func (x *Lifecycle) Reset() {
	*x = Lifecycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lifecycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lifecycle) ProtoMessage() {}

func (x *Lifecycle) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lifecycle.ProtoReflect.Descriptor instead.
func (*Lifecycle) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{2}
}

func (x *Lifecycle) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *Lifecycle) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// SpecCollection with repeated SpecEntity
type SpecCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains a spec.SpecEntity repeated
	Entities []*SpecEntity `protobuf:"bytes,4,rep,name=entities,proto3" json:"entities,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SpecCollection) Reset() {
	*x = SpecCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecCollection) ProtoMessage() {}

func (x *SpecCollection) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecCollection.ProtoReflect.Descriptor instead.
func (*SpecCollection) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{3}
}

func (x *SpecCollection) GetEntities() []*SpecEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *SpecCollection) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *SpecCollection) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// SpecEntity with Spec
type SpecEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// contains a spec.Spec
	Data *Type `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SpecEntity) Reset() {
	*x = SpecEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecEntity) ProtoMessage() {}

func (x *SpecEntity) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecEntity.ProtoReflect.Descriptor instead.
func (*SpecEntity) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{4}
}

func (x *SpecEntity) GetData() *Type {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SpecEntity) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *SpecEntity) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Proto options for a field
type Fieldproto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field numbers are used to identify your fields in the message binary format, and should not be changed once your message type is in use.
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number"`
	// Assign field to a protobuf oneof group.
	Oneof string `protobuf:"bytes,3,opt,name=oneof,proto3" json:"oneof"`
}

func (x *Fieldproto) Reset() {
	*x = Fieldproto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fieldproto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fieldproto) ProtoMessage() {}

func (x *Fieldproto) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fieldproto.ProtoReflect.Descriptor instead.
func (*Fieldproto) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{5}
}

func (x *Fieldproto) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Fieldproto) GetOneof() string {
	if x != nil {
		return x.Oneof
	}
	return ""
}

// ui hints for a field
type Fielduiextension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// component hint for ui-builder
	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component"`
	// UI element flags like full, double, hidden,...
	Flags []string `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags"`
	// Skip adding this field on ui init
	NoInit bool `protobuf:"varint,3,opt,name=no_init,json=noInit,proto3" json:"no_init"`
	// do not skip this field, even it is in the default skip list
	NoSkip bool `protobuf:"varint,4,opt,name=no_skip,json=noSkip,proto3" json:"no_skip"`
}

func (x *Fielduiextension) Reset() {
	*x = Fielduiextension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fielduiextension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fielduiextension) ProtoMessage() {}

func (x *Fielduiextension) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fielduiextension.ProtoReflect.Descriptor instead.
func (*Fielduiextension) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{6}
}

func (x *Fielduiextension) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Fielduiextension) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Fielduiextension) GetNoInit() bool {
	if x != nil {
		return x.NoInit
	}
	return false
}

func (x *Fielduiextension) GetNoSkip() bool {
	if x != nil {
		return x.NoSkip
	}
	return false
}

// Defines a queryparam field (for rpc type)
type Queryparam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// constraints for a field, like min{}, max{}, step{}. Not used at the moment
	Constraints map[string]*furo.FieldConstraint `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the field description
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// meta information for the client, like label, default, repeated, options...
	Meta *furo.FieldMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	// the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Queryparam) Reset() {
	*x = Queryparam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queryparam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queryparam) ProtoMessage() {}

func (x *Queryparam) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queryparam.ProtoReflect.Descriptor instead.
func (*Queryparam) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{7}
}

func (x *Queryparam) GetConstraints() map[string]*furo.FieldConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *Queryparam) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Queryparam) GetMeta() *furo.FieldMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Queryparam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Defines a field in the furo spec
type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// the field description
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description" yaml:"description"`
	// information for the proto generator, like number, type
	XProto *Fieldproto       `protobuf:"bytes,6,opt,name=__proto,json=Proto,proto3" json:"__proto" yaml:"__proto"`
	XUi    *Fielduiextension `protobuf:"bytes,7,opt,name=__ui,json=Ui,proto3" json:"__ui" yaml:"__ui"`
	// meta information for the client, like label, default, repeated, options...
	Meta *furo.FieldMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta" yaml:"meta"`

	// constraints for a field, like min{}, max{}, step{}
	Constraints map[string]*furo.FieldConstraint `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints" json:"yaml" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{8}
}

func (x *Field) GetXProto() *Fieldproto {
	if x != nil {
		return x.XProto
	}
	return nil
}

func (x *Field) GetXUi() *Fielduiextension {
	if x != nil {
		return x.XUi
	}
	return nil
}

func (x *Field) GetConstraints() map[string]*furo.FieldConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *Field) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Field) GetMeta() *furo.FieldMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Field) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// SpecEntity with Spec
type SpecServiceEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// contains a spec.Spec
	Data *Service `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SpecServiceEntity) Reset() {
	*x = SpecServiceEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecServiceEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecServiceEntity) ProtoMessage() {}

func (x *SpecServiceEntity) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecServiceEntity.ProtoReflect.Descriptor instead.
func (*SpecServiceEntity) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{9}
}

func (x *SpecServiceEntity) GetData() *Service {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SpecServiceEntity) GetLinks() []*furo.Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *SpecServiceEntity) GetMeta() *furo.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Main proto for a type
type Typeproto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the package this type belogs to
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package" yaml:"package"`
	// the target proto file for this type
	Targetfile string `protobuf:"bytes,3,opt,name=targetfile,proto3" json:"targetfile" yaml:"targetfile"`
	// needed imports like [ &#34;spec/spec.proto&#34;, &#34;google/protobuf/empty.proto&#34; ]
	Imports []string `protobuf:"bytes,2,rep,name=imports,proto3" json:"imports" yaml:"imports"`
	// Proto options Todo: find a solution for boolean options
	Options map[string]string `protobuf:"bytes,4,rep,name=options,proto3" json:"options" yaml:"options" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Typeproto) Reset() {
	*x = Typeproto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Typeproto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Typeproto) ProtoMessage() {}

func (x *Typeproto) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Typeproto.ProtoReflect.Descriptor instead.
func (*Typeproto) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{10}
}

func (x *Typeproto) GetImports() []string {
	if x != nil {
		return x.Imports
	}
	return nil
}

func (x *Typeproto) GetOptions() map[string]string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Typeproto) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Typeproto) GetTargetfile() string {
	if x != nil {
		return x.Targetfile
	}
	return ""
}

// Defines a rpc for a service
type Rpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// the service description
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`

	// Request and response types for the service
	Data *Servicereqres `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// This data is needed for...
	Deeplink *Servicedeeplink `protobuf:"bytes,5,opt,name=deeplink,proto3" json:"deeplink,omitempty"`

	// Query params, it is recomended to use string types
	// was a map[string]*Queryparam
	Query *orderedmap.OrderedMap `protobuf:"bytes,4,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// RPC name https://developers.google.com/protocol-buffers/docs/proto3#services
	RpcName string `protobuf:"bytes,2,opt,name=rpc_name,json=rpcName,proto3" json:"rpc_name,omitempty"`
}

func (x *Rpc) Reset() {
	*x = Rpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rpc) ProtoMessage() {}

func (x *Rpc) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Md.ProtoReflect.Descriptor instead.
func (*Rpc) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{11}
}

func (x *Rpc) GetData() *Servicereqres {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Rpc) GetDeeplink() *Servicedeeplink {
	if x != nil {
		return x.Deeplink
	}
	return nil
}

func (x *Rpc) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Rpc) GetQuery() *orderedmap.OrderedMap {
	if x != nil {
		return x.Query
	}
	return orderedmap.New()
}

func (x *Rpc) GetRpcName() string {
	if x != nil {
		return x.RpcName
	}
	return ""
}

// Defines a type in the furo spec
type Type struct {
	// Name of the type
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	// the type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// the type description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description" yaml:"description"`
	// information for the proto generator, should be removed for the client spec
	XProto *Typeproto `protobuf:"bytes,4,opt,name=__proto,json=Proto,proto3" json:"__proto" yaml:"__proto"`

	// fields of a type
	Fields *orderedmap.OrderedMap `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`

	Extensions *orderedmap.OrderedMap `yaml:"extensions"`

	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeSpec.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{12}
}

func (x *Type) GetXProto() *Typeproto {
	if x != nil {
		return x.XProto
	}
	return nil
}

func (x *Type) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Type) GetFields() *orderedmap.OrderedMap {
	if x != nil {
		return x.Fields
	}
	return orderedmap.New()
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// URL information for the service
type Servicedeeplink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describe the query params
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The link pattern, like /api/xxx/{qp}/yyy
	Href string `protobuf:"bytes,4,opt,name=href,proto3" json:"href,omitempty"`
	// method of curl
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	// the relationship
	Rel string `protobuf:"bytes,2,opt,name=rel,proto3" json:"rel,omitempty"`
}

func (x *Servicedeeplink) Reset() {
	*x = Servicedeeplink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_spec_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Servicedeeplink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servicedeeplink) ProtoMessage() {}

func (x *Servicedeeplink) ProtoReflect() protoreflect.Message {
	mi := &file_spec_spec_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Servicedeeplink.ProtoReflect.Descriptor instead.
func (*Servicedeeplink) Descriptor() ([]byte, []int) {
	return file_spec_spec_proto_rawDescGZIP(), []int{13}
}

func (x *Servicedeeplink) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Servicedeeplink) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Servicedeeplink) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Servicedeeplink) GetRel() string {
	if x != nil {
		return x.Rel
	}
	return ""
}

var File_spec_spec_proto protoreflect.FileDescriptor

var file_spec_spec_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x0f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x66, 0x75, 0x72, 0x6f, 0x2f, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x5f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x46, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x52, 0x70, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x71, 0x72, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x53, 0x70, 0x65, 0x63, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x0a, 0x53, 0x70, 0x65,
	0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0a, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x78, 0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x75, 0x69,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6e, 0x6f, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x73, 0x6b, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x53, 0x6b, 0x69, 0x70, 0x22,
	0x83, 0x02, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x43,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x55,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcd, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x28, 0x0a, 0x07, 0x5f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x0a, 0x04, 0x5f, 0x5f, 0x75,
	0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x75, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x02, 0x55, 0x69, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x55,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x75, 0x72, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x78, 0x0a, 0x11, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66,
	0x75, 0x72, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x66, 0x75, 0x72, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22,
	0xd3, 0x01, 0x0a, 0x09, 0x54, 0x79, 0x70, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x02, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x27, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x71, 0x72, 0x65, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x08, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x52,
	0x08, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x70, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x4a, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf1,
	0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x5f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x46, 0x0a, 0x0b, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x71, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x65, 0x65,
	0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x72, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_spec_proto_rawDescOnce sync.Once
	file_spec_spec_proto_rawDescData = file_spec_spec_proto_rawDesc
)

func file_spec_spec_proto_rawDescGZIP() []byte {
	file_spec_spec_proto_rawDescOnce.Do(func() {
		file_spec_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_spec_proto_rawDescData)
	})
	return file_spec_spec_proto_rawDescData
}

var file_spec_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_spec_spec_proto_goTypes = []interface{}{
	(*Service)(nil),              // 0: spec.Service
	(*Servicereqres)(nil),        // 1: spec.Servicereqres
	(*Lifecycle)(nil),            // 2: spec.Lifecycle
	(*SpecCollection)(nil),       // 3: spec.SpecCollection
	(*SpecEntity)(nil),           // 4: spec.SpecEntity
	(*Fieldproto)(nil),           // 5: spec.Fieldproto
	(*Fielduiextension)(nil),     // 6: spec.Fielduiextension
	(*Queryparam)(nil),           // 7: spec.Queryparam
	(*Field)(nil),                // 8: spec.Field
	(*SpecServiceEntity)(nil),    // 9: spec.SpecServiceEntity
	(*Typeproto)(nil),            // 10: spec.Typeproto
	(*Rpc)(nil),                  // 11: spec.Md
	(*Type)(nil),                 // 12: spec.TypeSpec
	(*Servicedeeplink)(nil),      // 13: spec.Servicedeeplink
	nil,                          // 14: spec.Service.ServicesEntry
	nil,                          // 15: spec.Queryparam.ConstraintsEntry
	nil,                          // 16: spec.Field.ConstraintsEntry
	nil,                          // 17: spec.Typeproto.OptionsEntry
	nil,                          // 18: spec.Md.QueryEntry
	nil,                          // 19: spec.TypeSpec.FieldsEntry
	(*furo.Link)(nil),            // 20: furo.Link
	(*furo.Meta)(nil),            // 21: furo.Meta
	(*furo.FieldMeta)(nil),       // 22: furo.FieldMeta
	(*furo.FieldConstraint)(nil), // 23: furo.FieldConstraint
}
var file_spec_spec_proto_depIdxs = []int32{
	10, // 0: spec.Service.__proto:type_name -> spec.Typeproto
	2,  // 1: spec.Service.lifecycle:type_name -> spec.Lifecycle
	14, // 2: spec.Service.services:type_name -> spec.Service.ServicesEntry
	4,  // 3: spec.SpecCollection.entities:type_name -> spec.SpecEntity
	20, // 4: spec.SpecCollection.links:type_name -> furo.Link
	21, // 5: spec.SpecCollection.meta:type_name -> furo.Meta
	12, // 6: spec.SpecEntity.data:type_name -> spec.TypeSpec
	20, // 7: spec.SpecEntity.links:type_name -> furo.Link
	21, // 8: spec.SpecEntity.meta:type_name -> furo.Meta
	15, // 9: spec.Queryparam.constraints:type_name -> spec.Queryparam.ConstraintsEntry
	22, // 10: spec.Queryparam.meta:type_name -> furo.FieldMeta
	5,  // 11: spec.Field.__proto:type_name -> spec.Fieldproto
	6,  // 12: spec.Field.__ui:type_name -> spec.Fielduiextension
	16, // 13: spec.Field.constraints:type_name -> spec.Field.ConstraintsEntry
	22, // 14: spec.Field.meta:type_name -> furo.FieldMeta
	0,  // 15: spec.SpecServiceEntity.data:type_name -> spec.Service
	20, // 16: spec.SpecServiceEntity.links:type_name -> furo.Link
	21, // 17: spec.SpecServiceEntity.meta:type_name -> furo.Meta
	17, // 18: spec.Typeproto.options:type_name -> spec.Typeproto.OptionsEntry
	1,  // 19: spec.Md.data:type_name -> spec.Servicereqres
	13, // 20: spec.Md.deeplink:type_name -> spec.Servicedeeplink
	18, // 21: spec.Md.query:type_name -> spec.Md.QueryEntry
	10, // 22: spec.TypeSpec.__proto:type_name -> spec.Typeproto
	19, // 23: spec.TypeSpec.fields:type_name -> spec.TypeSpec.FieldsEntry
	11, // 24: spec.Service.ServicesEntry.value:type_name -> spec.Md
	23, // 25: spec.Queryparam.ConstraintsEntry.value:type_name -> furo.FieldConstraint
	23, // 26: spec.Field.ConstraintsEntry.value:type_name -> furo.FieldConstraint
	7,  // 27: spec.Md.QueryEntry.value:type_name -> spec.Queryparam
	8,  // 28: spec.TypeSpec.FieldsEntry.value:type_name -> spec.Field
	29, // [29:29] is the sub-list for method output_type
	29, // [29:29] is the sub-list for method input_type
	29, // [29:29] is the sub-list for extension type_name
	29, // [29:29] is the sub-list for extension extendee
	0,  // [0:29] is the sub-list for field type_name
}

func init() { file_spec_spec_proto_init() }
func file_spec_spec_proto_init() {
	if File_spec_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_spec_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Servicereqres); i {
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
		file_spec_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lifecycle); i {
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
		file_spec_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecCollection); i {
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
		file_spec_spec_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecEntity); i {
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
		file_spec_spec_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fieldproto); i {
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
		file_spec_spec_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fielduiextension); i {
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
		file_spec_spec_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queryparam); i {
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
		file_spec_spec_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_spec_spec_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecServiceEntity); i {
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
		file_spec_spec_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Typeproto); i {
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
		file_spec_spec_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rpc); i {
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
		file_spec_spec_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_spec_spec_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Servicedeeplink); i {
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
			RawDescriptor: file_spec_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spec_spec_proto_goTypes,
		DependencyIndexes: file_spec_spec_proto_depIdxs,
		MessageInfos:      file_spec_spec_proto_msgTypes,
	}.Build()
	File_spec_spec_proto = out.File
	file_spec_spec_proto_rawDesc = nil
	file_spec_spec_proto_goTypes = nil
	file_spec_spec_proto_depIdxs = nil
}
