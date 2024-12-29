package tsutils

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Work with protoc-gen-ts_proto: https://github.com/stephenh/ts-proto

func TSModule_TSProto(file protoreflect.FileDescriptor) TSModule {
	protoPath := file.Path()
	return TSModule{
		ModuleName: GetModuleName(file),
		Path:       strings.TrimSuffix(protoPath, ".proto") + ".ts",
		Relative:   true,
	}
}

func TSIdent_TSProto_Message(msg *protogen.Message) TSIdent {
	return TSModule_TSProto(msg.Desc.ParentFile()).Ident(msg.GoIdent.GoName)
}

func TSProto_EnumToJSONFuncName(g *TSRegistry, enum protoreflect.EnumDescriptor) string {
	return FunctionCase_TSProto(string(enum.Name())) + "ToJSON"
}

func TSProtoMessageToJson(msgTyp *protogen.Message) func(g *TSRegistry, in string) string {
	return func(g *TSRegistry, in string) string {
		ident := g.QualifiedTSIdent(TSIdent_TSProto_Message(msgTyp))
		return ident + `.toJSON(` + in + `)`
	}
}
func TSProtoEnumToJson(enumTyp *protogen.Enum) func(g *TSRegistry, in string) string {
	return func(g *TSRegistry, in string) string {
		enumModule := TSModule_TSProto(enumTyp.Desc.ParentFile())
		toJsonIdent := enumModule.Ident(TSProto_EnumToJSONFuncName(g, enumTyp.Desc))
		toJsonFunc := g.QualifiedTSIdent(toJsonIdent)
		return toJsonFunc + `(` + in + `)`
	}
}

func TSProtoScalarToJson() func(g *TSRegistry, in string) string {
	return func(g *TSRegistry, in string) string {
		return in
	}
}

func TSProtoTimestampToJson() func(g *TSRegistry, in string) string {
	return func(g *TSRegistry, in string) string {
		// in is type of Date
		return in + `.toISOString()`
	}
}

func TSProtoFieldMaskToJson(msgTyp *protogen.Message) func(g *TSRegistry, in string) string {
	return func(g *TSRegistry, in string) string {
		// in is type of FieldMask
		ident := g.QualifiedTSIdent(TSIdent_TSProto_Message(msgTyp))
		return fmt.Sprintf(`%s.toJSON(%s.wrap(%s))`, ident, ident, in)
	}
}
