//Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
//versions:
//- protoc-gen-go-fieldmask v1.0.20
//- protoc (unknown)
//source: proto/user/user.proto

package user

import (
	fieldmask "github.com/RyoJerryYu/protoc-gen-pluginx/pkg/fieldmask"
)

// IUserPathBuilder is the interface for the field path of User
type IUserPathBuilder interface {
	String() string
	Id() fieldmask.IEndPathBuilder
	Name() fieldmask.IEndPathBuilder
	Icon() IIconPathBuilder
	Password() IUserPasswordPathBuilder
	CreatedAt() fieldmask.IEndPathBuilder
}

// userPathBuilder is the implementation for the field path of User
type userPathBuilder struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewUserPathBuilder creates a new userPathBuilder
func NewUserPathBuilder(fieldPath string) IUserPathBuilder {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return userPathBuilder{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x userPathBuilder) String() string { return x.fieldPath }

func (x userPathBuilder) Id() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "id")
}
func (x userPathBuilder) Name() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "name")
}
func (x userPathBuilder) Icon() IIconPathBuilder { return NewIconPathBuilder(x.prefix + "icon") }
func (x userPathBuilder) Password() IUserPasswordPathBuilder {
	return NewUserPasswordPathBuilder(x.prefix + "password")
}
func (x userPathBuilder) CreatedAt() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "created_at")
}

// PathBuilder returns the field path for User
func (x *User) PathBuilder() IUserPathBuilder {
	return NewUserPathBuilder("")
}

// IIconPathBuilder is the interface for the field path of Icon
type IIconPathBuilder interface {
	String() string
	Id() fieldmask.IEndPathBuilder
	Url() fieldmask.IEndPathBuilder
	CreatedAt() fieldmask.ITimestampPathBuilder
	Nested() IIcon_NestedPathBuilder
	NestedAnother() fieldmask.IEndPathBuilder
}

// iconPathBuilder is the implementation for the field path of Icon
type iconPathBuilder struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewIconPathBuilder creates a new iconPathBuilder
func NewIconPathBuilder(fieldPath string) IIconPathBuilder {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return iconPathBuilder{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x iconPathBuilder) String() string { return x.fieldPath }

func (x iconPathBuilder) Id() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "id")
}
func (x iconPathBuilder) Url() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "url")
}
func (x iconPathBuilder) CreatedAt() fieldmask.ITimestampPathBuilder {
	return fieldmask.NewTimestampPathBuilder(x.prefix + "created_at")
}
func (x iconPathBuilder) Nested() IIcon_NestedPathBuilder {
	return NewIcon_NestedPathBuilder(x.prefix + "nested")
}
func (x iconPathBuilder) NestedAnother() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "nested_another")
}

// PathBuilder returns the field path for Icon
func (x *Icon) PathBuilder() IIconPathBuilder {
	return NewIconPathBuilder("")
}

// IIcon_NestedPathBuilder is the interface for the field path of Icon_Nested
type IIcon_NestedPathBuilder interface {
	String() string
	SomeField() fieldmask.IEndPathBuilder
}

// icon_NestedPathBuilder is the implementation for the field path of Icon_Nested
type icon_NestedPathBuilder struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewIcon_NestedPathBuilder creates a new icon_NestedPathBuilder
func NewIcon_NestedPathBuilder(fieldPath string) IIcon_NestedPathBuilder {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return icon_NestedPathBuilder{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x icon_NestedPathBuilder) String() string { return x.fieldPath }

func (x icon_NestedPathBuilder) SomeField() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "some_field")
}

// PathBuilder returns the field path for Icon_Nested
func (x *Icon_Nested) PathBuilder() IIcon_NestedPathBuilder {
	return NewIcon_NestedPathBuilder("")
}
