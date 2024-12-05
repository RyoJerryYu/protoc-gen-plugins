//Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
//versions:
//- protoc-gen-go-fieldmask v1.0.18
//- protoc (unknown)
//source: proto/admin/admin.proto

package admin

import (
	fieldmask "github.com/RyoJerryYu/protoc-gen-pluginx/pkg/fieldmask"
	user "github.com/RyoJerryYu/protoc-gen-pluginx/tests/protoc-gen-go-fieldmask/e2e/proto/user"
)

// IAdminPathBuilder is the interface for the field path of Admin
type IAdminPathBuilder interface {
	String() string
	Id() fieldmask.IEndPathBuilder
	Name() fieldmask.IEndPathBuilder
	User() user.IUserPathBuilder
	CreatedAt() fieldmask.ITimestampPathBuilder
}

// adminPathBuilder is the implementation for the field path of Admin
type adminPathBuilder struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewAdminPathBuilder creates a new adminPathBuilder
func NewAdminPathBuilder(fieldPath string) IAdminPathBuilder {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return adminPathBuilder{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x adminPathBuilder) String() string { return x.fieldPath }

func (x adminPathBuilder) Id() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "id")
}
func (x adminPathBuilder) Name() fieldmask.IEndPathBuilder {
	return fieldmask.NewEndPathBuilder(x.prefix + "name")
}
func (x adminPathBuilder) User() user.IUserPathBuilder {
	return user.NewUserPathBuilder(x.prefix + "user")
}
func (x adminPathBuilder) CreatedAt() fieldmask.ITimestampPathBuilder {
	return fieldmask.NewTimestampPathBuilder(x.prefix + "created_at")
}

// PathBuilder returns the field path for Admin
func (x *Admin) PathBuilder() IAdminPathBuilder {
	return NewAdminPathBuilder("")
}
