package clientspec

import (
	"github.com/theNorstroem/spectools/pkg/domaintypes/furo"
	"github.com/theNorstroem/spectools/pkg/orderedmap"
	"github.com/theNorstroem/spectools/pkg/specSpec"
)

func CreateFromAstType(ast *specSpec.Type) (t *Type) {

	t = &Type{
		Name:   ast.Name,
		Type:   ast.Type,
		Fields: orderedmap.New(),
	}

	ast.Fields.Map(func(iKey interface{}, iValue interface{}) {
		astField := iValue.(*specSpec.Field)
		field := &Field{
			Type:        astField.Type,
			Meta:        astField.Meta,
			Constraints: astField.Constraints,
		}
		t.Fields.Set(iKey, field)
	})

	return t
}

// Defines a type in the furo spec
type Type struct {
	// Name of the type
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	// the type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// fields of a type
	Fields *orderedmap.OrderedMap `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
type Field struct {
	// the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// meta information for the client, like label, default, repeated, options...
	Meta *furo.FieldMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta" yaml:"meta"`
	// constraints for a field, like min{}, max{}, step{}
	Constraints map[string]*furo.FieldConstraint `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints" json:"yaml" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
