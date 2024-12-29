package tsutils

import (
	"fmt"
	"strings"

	"github.com/RyoJerryYu/protoc-gen-pluginx/pkg/protobufx"
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
	messageToJson := func(g *TSRegistry, in string) string {
		ident := g.QualifiedTSIdent(TSIdent_TSProto_Message(msgTyp))
		return ident + `.toJSON(` + in + `)`
	}
	if protobufx.IsWellKnownType(msgTyp.Desc) {
		switch msgTyp.Desc.Name() {
		case protobufx.Any_message_name,
			protobufx.Empty_message_name:
			// Any and Empty, jsonify as a normal message
			return messageToJson
		case protobufx.Struct_message_name,
			protobufx.Value_message_name,
			protobufx.ListValue_message_name:
			// Struct, Value, ListValue, jsonify as what they are
			// It presents as a JSON object, so we don't need to convert it
			return TSProtoScalarToJson()
		case protobufx.Timestamp_message_name:
			// Timestamp represents as Date in ts-proto
			// and need to convert to a ISO string
			return TSProtoTimestampToJson()
		case protobufx.Duration_message_name:
			// don't know why ts-proto treats Duration as a normal message
			// in the docs, it should be a string like "1.234s"
			return messageToJson
		case protobufx.FieldMask_message_name:
			// FieldMask represents as []string in ts-proto
			// and need to convert to strings joined by ","
			return TSProtoFieldMaskToJson(msgTyp)
		case protobufx.BoolValue_message_name,
			protobufx.StringValue_message_name,
			protobufx.DoubleValue_message_name,
			protobufx.FloatValue_message_name,
			protobufx.Int32Value_message_name,
			protobufx.Int64Value_message_name,
			protobufx.UInt32Value_message_name,
			protobufx.UInt64Value_message_name:
			// well-known scalar types,
			// represent as a scalar in ts-proto
			// and should be as it is in JSON
			return TSProtoScalarToJson()
		default:
			// other types for reflection or syntax types
			// should be treated as a normal message
			return TSProtoMessageToJson(msgTyp)
		}
	}
	return messageToJson
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
