// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: proto/oneofenum/oneof_enum.proto

/* eslint-disable */

export const protobufPackage = "proto.oneofenum";

export enum ExampleEnum {
  EXAMPLE_ENUM_UNSPECIFIED = 0,
  EXAMPLE_ENUM_FIRST = 1,
  UNRECOGNIZED = -1,
}

export function exampleEnumFromJSON(object: any): ExampleEnum {
  switch (object) {
    case 0:
    case "EXAMPLE_ENUM_UNSPECIFIED":
      return ExampleEnum.EXAMPLE_ENUM_UNSPECIFIED;
    case 1:
    case "EXAMPLE_ENUM_FIRST":
      return ExampleEnum.EXAMPLE_ENUM_FIRST;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ExampleEnum.UNRECOGNIZED;
  }
}

export function exampleEnumToJSON(object: ExampleEnum): string {
  switch (object) {
    case ExampleEnum.EXAMPLE_ENUM_UNSPECIFIED:
      return "EXAMPLE_ENUM_UNSPECIFIED";
    case ExampleEnum.EXAMPLE_ENUM_FIRST:
      return "EXAMPLE_ENUM_FIRST";
    case ExampleEnum.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface OneofEnumMessage {
  exampleEnum?: ExampleEnum | undefined;
}

function createBaseOneofEnumMessage(): OneofEnumMessage {
  return { exampleEnum: undefined };
}

export const OneofEnumMessage: MessageFns<OneofEnumMessage> = {
  fromJSON(object: any): OneofEnumMessage {
    return {
      exampleEnum: isSet(object.exampleEnum)
        ? exampleEnumFromJSON(object.exampleEnum)
        : undefined,
    };
  },

  toJSON(message: OneofEnumMessage): unknown {
    const obj: any = {};
    if (message.exampleEnum !== undefined) {
      obj.exampleEnum = exampleEnumToJSON(message.exampleEnum);
    }
    return obj;
  },

  create(base?: DeepPartial<OneofEnumMessage>): OneofEnumMessage {
    return OneofEnumMessage.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<OneofEnumMessage>): OneofEnumMessage {
    const message = createBaseOneofEnumMessage();
    message.exampleEnum = object.exampleEnum ?? undefined;
    return message;
  },
};

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends globalThis.Array<infer U>
    ? globalThis.Array<DeepPartial<U>>
    : T extends ReadonlyArray<infer U>
      ? ReadonlyArray<DeepPartial<U>>
      : T extends {}
        ? { [K in keyof T]?: DeepPartial<T[K]> }
        : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
