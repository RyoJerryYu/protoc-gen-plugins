import { Duration } from "../../google/protobuf/duration";
import { Empty } from "../../google/protobuf/empty";
import { StringValue } from "../../google/protobuf/wrappers";
import {
  CallOptions,
  ClientMiddleware,
  ClientMiddlewareCall,
} from "nice-grpc-common";
import {
  ABitOfEverything,
  ABitOfEverythingRepeated,
  ABitOfEverythingServiceClient,
  ABitOfEverything_Nested,
  AnotherServiceWithNoBindingsClient,
  Body,
  CheckStatusResponse,
  DeepPartial,
  MessageWithBody,
  RequiredMessageTypeRequest,
  UpdateV2Request,
} from "./a_bit_of_everything";
import { OneofEnumMessage, exampleEnumToJSON } from "../oneofenum/oneof_enum";
import {
  MessageWithNestedPathEnum,
  MessageWithPathEnum,
} from "../pathenum/path_enum";
import { IdMessage } from "../sub2/message";

function composeClientMiddleware<Ext1, Ext2, RequiredCallOptionsExt>(
  middleware1: ClientMiddleware<Ext1, RequiredCallOptionsExt>,
  middleware2: ClientMiddleware<Ext2, RequiredCallOptionsExt & Ext1>,
): ClientMiddleware<Ext1 & Ext2, RequiredCallOptionsExt> {
  return <Request, Response>(
    call: ClientMiddlewareCall<
      Request,
      Response,
      Ext1 & Ext2 & RequiredCallOptionsExt
    >,
    options: CallOptions & Partial<Ext1 & Ext2 & RequiredCallOptionsExt>,
  ) => {
    return middleware2<Request, Response>(
      {
        ...call,
        next: (request, options2) => {
          return middleware1<Request, Response>(
            { ...call, request } as any,
            options2,
          ) as any;
        },
      },
      options,
    );
  };
}

type Primitive = string | boolean | number;
type RequestPayload = Record<string, unknown>;
type FlattenedRequestPayload = Record<string, Primitive | Primitive[]>;

/**
 * Checks if given value is a plain object
 * Logic copied and adapted from below source:
 * https://github.com/char0n/ramda-adjunct/blob/master/src/isPlainObj.js
 */
function isPlainObject(value: unknown): boolean {
  const isObject =
    Object.prototype.toString.call(value).slice(8, -1) === "Object";
  const isObjLike = value !== null && isObject;

  if (!isObjLike || !isObject) {
    return false;
  }

  const proto: unknown = Object.getPrototypeOf(value);

  const hasObjectConstructor = !!(
    proto &&
    typeof proto === "object" &&
    proto.constructor === Object.prototype.constructor
  );

  return hasObjectConstructor;
}

/**
 * Checks if given value is of a primitive type
 */
function isPrimitive(value: unknown): boolean {
  return ["string", "number", "boolean"].some((t) => typeof value === t);
}

/**
 * Flattens a deeply nested request payload and returns an object
 * with only primitive values and non-empty array of primitive values
 * as per https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
 */
function flattenRequestPayload<T extends RequestPayload>(
  requestPayload: T,
  path = "",
): FlattenedRequestPayload {
  return Object.keys(requestPayload).reduce((acc: T, key: string): T => {
    const value = requestPayload[key];
    const newPath = path ? [path, key].join(".") : key;

    const isNonEmptyPrimitiveArray =
      Array.isArray(value) &&
      value.every((v) => isPrimitive(v)) &&
      value.length > 0;

    let objectToMerge = {};

    if (isPlainObject(value)) {
      objectToMerge = flattenRequestPayload(value as RequestPayload, newPath);
    } else if (isPrimitive(value) || isNonEmptyPrimitiveArray) {
      objectToMerge = { [newPath]: value };
    }

    return { ...acc, ...objectToMerge };
  }, {} as T) as FlattenedRequestPayload;
}

/**
 * Renders a deeply nested request payload into a string of URL search
 * parameters by first flattening the request payload and then removing keys
 * which are already present in the URL path.
 */
function renderURLSearchParams<T extends RequestPayload>(
  requestPayload: T,
  urlPathParams: string[] = [],
): string[][] {
  const flattenedRequestPayload = flattenRequestPayload(requestPayload);

  const urlSearchParams = Object.keys(flattenedRequestPayload).reduce(
    (acc: string[][], key: string): string[][] => {
      // key should not be present in the url path as a parameter
      const value = flattenedRequestPayload[key];
      if (urlPathParams.find((f) => f === key)) {
        return acc;
      }
      return Array.isArray(value)
        ? [...acc, ...value.map((m) => [key, m.toString()])]
        : (acc = [...acc, [key, value.toString()]]);
    },
    [] as string[][],
  );

  return urlSearchParams;
}

/**
 * must is a utility function that throws an error if the given value is null or undefined
 */
function must<T>(value: T | null | undefined): T {
  if (value === null || value === undefined) {
    throw new Error("Value is null or undefined");
  }
  return value;
}

/**
 * CallParams is a type that represents the parameters that are passed to the transport's call method
 */
export type CallParams = {
  url: string;
  method: string;
  queryParams?: string[][];
  body?: BodyInit | null;
};

/**
 * Transport is a type that represents the interface of a transport object
 */
export type Transport = {
  call(params: CallParams): Promise<any>;
};

// ABitOfEverything service is used to validate that APIs with complicated
// proto messages and URL templates are still processed correctly.
export function newABitOfEverythingService(
  transport: Transport,
  middlewares?: ClientMiddleware[],
): ABitOfEverythingServiceClient {
  //Compose middleware
  let middleware: ClientMiddleware = (call, options) => {
    return call.next(call.request, options);
  };
  for (let i = 0; i < middlewares?.length || 0; i++) {
    middleware = composeClientMiddleware(middleware, middlewares[i]);
  }

  return {
    async createBody(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything`,
        method: "POST",
        body: JSON.stringify(ABitOfEverything.toJSON(fullReq)),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async lookup(
      req: DeepPartial<IdMessage>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = IdMessage.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/${must(fullReq.uuid)}`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["uuid"]),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async custom(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/${must(fullReq.uuid)}:custom`,
        method: "POST",
      });
      return ABitOfEverything.fromJSON(res);
    },

    async doubleColon(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/${must(fullReq.uuid)}:custom:custom`,
        method: "POST",
      });
      return ABitOfEverything.fromJSON(res);
    },

    async update(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/${must(fullReq.uuid)}`,
        method: "PUT",
        body: JSON.stringify(ABitOfEverything.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },

    async updateV2(
      req: DeepPartial<UpdateV2Request>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = UpdateV2Request.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/a_bit_of_everything/${must(fullReq.abe?.uuid)}`,
        method: "PUT",
        body: JSON.stringify(ABitOfEverything.toJSON(must(fullReq.abe))),
      });
      return Empty.fromJSON(res);
    },

    async delete(
      req: DeepPartial<IdMessage>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = IdMessage.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/${must(fullReq.uuid)}`,
        method: "DELETE",
        queryParams: renderURLSearchParams(req, ["uuid"]),
      });
      return Empty.fromJSON(res);
    },

    async getQuery(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/query/${must(fullReq.uuid)}`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["uuid"]),
      });
      return Empty.fromJSON(res);
    },

    async getRepeatedQuery(
      req: DeepPartial<ABitOfEverythingRepeated>,
      options?: CallOptions,
    ): Promise<ABitOfEverythingRepeated> {
      const fullReq = ABitOfEverythingRepeated.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything_repeated/${must(fullReq.pathRepeatedFloatValue)}/${must(fullReq.pathRepeatedDoubleValue)}/${must(fullReq.pathRepeatedInt64Value)}/${must(fullReq.pathRepeatedUint64Value)}/${must(fullReq.pathRepeatedInt32Value)}/${must(fullReq.pathRepeatedFixed64Value)}/${must(fullReq.pathRepeatedFixed32Value)}/${must(fullReq.pathRepeatedBoolValue)}/${must(fullReq.pathRepeatedStringValue)}/${must(fullReq.pathRepeatedBytesValue)}/${must(fullReq.pathRepeatedUint32Value)}/${must(fullReq.pathRepeatedEnumValue)}/${must(fullReq.pathRepeatedSfixed32Value)}/${must(fullReq.pathRepeatedSfixed64Value)}/${must(fullReq.pathRepeatedSint32Value)}/${must(fullReq.pathRepeatedSint64Value)}`,
        method: "GET",
        queryParams: renderURLSearchParams(req, [
          "pathRepeatedFloatValue",
          "pathRepeatedDoubleValue",
          "pathRepeatedInt64Value",
          "pathRepeatedUint64Value",
          "pathRepeatedInt32Value",
          "pathRepeatedFixed64Value",
          "pathRepeatedFixed32Value",
          "pathRepeatedBoolValue",
          "pathRepeatedStringValue",
          "pathRepeatedBytesValue",
          "pathRepeatedUint32Value",
          "pathRepeatedEnumValue",
          "pathRepeatedSfixed32Value",
          "pathRepeatedSfixed64Value",
          "pathRepeatedSint32Value",
          "pathRepeatedSint64Value",
        ]),
      });
      return ABitOfEverythingRepeated.fromJSON(res);
    },

    async deepPathEcho(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/deep_path/${must(fullReq.singleNested?.name)}`,
        method: "POST",
        body: JSON.stringify(ABitOfEverything.toJSON(fullReq)),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async noBindings(
      req: DeepPartial<Duration>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Duration.fromPartial(req);
      const res = await transport.call({
        url: `/proto.examplepb.ABitOfEverythingService/NoBindings`,
        method: "POST",
        body: JSON.stringify(Duration.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },

    async timeout(
      req: DeepPartial<Empty>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Empty.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/timeout`,
        method: "GET",
        queryParams: renderURLSearchParams(req, []),
      });
      return Empty.fromJSON(res);
    },

    async errorWithDetails(
      req: DeepPartial<Empty>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Empty.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/errorwithdetails`,
        method: "GET",
        queryParams: renderURLSearchParams(req, []),
      });
      return Empty.fromJSON(res);
    },

    async getMessageWithBody(
      req: DeepPartial<MessageWithBody>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = MessageWithBody.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/withbody/${must(fullReq.id)}`,
        method: "POST",
        body: JSON.stringify(Body.toJSON(must(fullReq.data))),
      });
      return Empty.fromJSON(res);
    },

    async postWithEmptyBody(
      req: DeepPartial<Body>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Body.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/postwithemptybody/${must(fullReq.name)}`,
        method: "POST",
        body: JSON.stringify(Body.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },

    async checkGetQueryParams(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/params/get/${must(fullReq.singleNested?.name)}`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["singleNested.name"]),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async checkNestedEnumGetQueryParams(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/params/get/nested_enum/${must(fullReq.singleNested?.ok)}`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["singleNested.ok"]),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async checkPostQueryParams(
      req: DeepPartial<ABitOfEverything>,
      options?: CallOptions,
    ): Promise<ABitOfEverything> {
      const fullReq = ABitOfEverything.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/a_bit_of_everything/params/post/${must(fullReq.stringValue)}`,
        method: "POST",
        body: JSON.stringify(
          ABitOfEverything_Nested.toJSON(must(fullReq.singleNested)),
        ),
      });
      return ABitOfEverything.fromJSON(res);
    },

    async overwriteRequestContentType(
      req: DeepPartial<Body>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Body.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/overwriterequestcontenttype`,
        method: "POST",
        body: JSON.stringify(Body.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },

    async overwriteResponseContentType(
      req: DeepPartial<Empty>,
      options?: CallOptions,
    ): Promise<StringValue> {
      const fullReq = Empty.fromPartial(req);
      const res = await transport.call({
        url: `/v2/example/overwriteresponsecontenttype`,
        method: "GET",
        queryParams: renderURLSearchParams(req, []),
      });
      return StringValue.fromJSON(res);
    },

    async checkExternalPathEnum(
      req: DeepPartial<MessageWithPathEnum>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = MessageWithPathEnum.fromPartial(req);
      const res = await transport.call({
        url: `/v2/${must(fullReq.value)}:check`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["value"]),
      });
      return Empty.fromJSON(res);
    },

    async checkExternalNestedPathEnum(
      req: DeepPartial<MessageWithNestedPathEnum>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = MessageWithNestedPathEnum.fromPartial(req);
      const res = await transport.call({
        url: `/v3/${must(fullReq.value)}:check`,
        method: "GET",
        queryParams: renderURLSearchParams(req, ["value"]),
      });
      return Empty.fromJSON(res);
    },

    async checkStatus(
      req: DeepPartial<Empty>,
      options?: CallOptions,
    ): Promise<CheckStatusResponse> {
      const fullReq = Empty.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/checkStatus`,
        method: "GET",
        queryParams: renderURLSearchParams(req, []),
      });
      return CheckStatusResponse.fromJSON(res);
    },

    async postOneofEnum(
      req: DeepPartial<OneofEnumMessage>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = OneofEnumMessage.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/oneofenum`,
        method: "POST",
        body: exampleEnumToJSON(must(fullReq.exampleEnum)),
      });
      return Empty.fromJSON(res);
    },

    async postRequiredMessageType(
      req: DeepPartial<RequiredMessageTypeRequest>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = RequiredMessageTypeRequest.fromPartial(req);
      const res = await transport.call({
        url: `/v1/example/requiredmessagetype`,
        method: "POST",
        body: JSON.stringify(RequiredMessageTypeRequest.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },
  };
}

export function newAnotherServiceWithNoBindings(
  transport: Transport,
  middlewares?: ClientMiddleware[],
): AnotherServiceWithNoBindingsClient {
  //Compose middleware
  let middleware: ClientMiddleware = (call, options) => {
    return call.next(call.request, options);
  };
  for (let i = 0; i < middlewares?.length || 0; i++) {
    middleware = composeClientMiddleware(middleware, middlewares[i]);
  }

  return {
    async noBindings(
      req: DeepPartial<Empty>,
      options?: CallOptions,
    ): Promise<Empty> {
      const fullReq = Empty.fromPartial(req);
      const res = await transport.call({
        url: `/proto.examplepb.AnotherServiceWithNoBindings/NoBindings`,
        method: "POST",
        body: JSON.stringify(Empty.toJSON(fullReq)),
      });
      return Empty.fromJSON(res);
    },
  };
}
