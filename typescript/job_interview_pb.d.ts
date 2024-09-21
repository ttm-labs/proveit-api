import * as jspb from 'google-protobuf'



export class JobInterviewRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobInterviewRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JobInterviewRequest): JobInterviewRequest.AsObject;
  static serializeBinaryToWriter(message: JobInterviewRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobInterviewRequest;
  static deserializeBinaryFromReader(message: JobInterviewRequest, reader: jspb.BinaryReader): JobInterviewRequest;
}

export namespace JobInterviewRequest {
  export type AsObject = {
  }
}

export class JobInterviewResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobInterviewResponse.AsObject;
  static toObject(includeInstance: boolean, msg: JobInterviewResponse): JobInterviewResponse.AsObject;
  static serializeBinaryToWriter(message: JobInterviewResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobInterviewResponse;
  static deserializeBinaryFromReader(message: JobInterviewResponse, reader: jspb.BinaryReader): JobInterviewResponse;
}

export namespace JobInterviewResponse {
  export type AsObject = {
  }
}

