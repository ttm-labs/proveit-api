/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 5.27.3
 * source: job_interview_analysis.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./google/protobuf/timestamp";
import * as pb_1 from "google-protobuf";
import * as grpc_1 from "@grpc/grpc-js";
export declare namespace job_interview_analysis_service {
    export class CreateJobInterviewAnalysisRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_interview_id?: string;
        });
        get job_interview_id(): string;
        set job_interview_id(value: string);
        static fromObject(data: {
            job_interview_id?: string;
        }): CreateJobInterviewAnalysisRequest;
        toObject(): {
            job_interview_id?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateJobInterviewAnalysisRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): CreateJobInterviewAnalysisRequest;
    }
    export class CreateJobInterviewAnalysisResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_interview_analysis_id?: string;
        });
        get job_interview_analysis_id(): string;
        set job_interview_analysis_id(value: string);
        static fromObject(data: {
            job_interview_analysis_id?: string;
        }): CreateJobInterviewAnalysisResponse;
        toObject(): {
            job_interview_analysis_id?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateJobInterviewAnalysisResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): CreateJobInterviewAnalysisResponse;
    }
    export class ReadJobInterviewAnalysisRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_interview_id?: string;
        });
        get job_interview_id(): string;
        set job_interview_id(value: string);
        static fromObject(data: {
            job_interview_id?: string;
        }): ReadJobInterviewAnalysisRequest;
        toObject(): {
            job_interview_id?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ReadJobInterviewAnalysisRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): ReadJobInterviewAnalysisRequest;
    }
    export class ReadJobInterviewAnalysisResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            analysis?: Analysis;
        });
        get analysis(): Analysis;
        set analysis(value: Analysis);
        get has_analysis(): boolean;
        static fromObject(data: {
            analysis?: ReturnType<typeof Analysis.prototype.toObject>;
        }): ReadJobInterviewAnalysisResponse;
        toObject(): {
            analysis?: ReturnType<typeof Analysis.prototype.toObject>;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ReadJobInterviewAnalysisResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): ReadJobInterviewAnalysisResponse;
    }
    export class UpdateJobInterviewAnalysisRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {});
        static fromObject(data: {}): UpdateJobInterviewAnalysisRequest;
        toObject(): {};
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): UpdateJobInterviewAnalysisRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): UpdateJobInterviewAnalysisRequest;
    }
    export class UpdateJobInterviewAnalysisResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {});
        static fromObject(data: {}): UpdateJobInterviewAnalysisResponse;
        toObject(): {};
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): UpdateJobInterviewAnalysisResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): UpdateJobInterviewAnalysisResponse;
    }
    export class DeleteJobInterviewAnalysisRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {});
        static fromObject(data: {}): DeleteJobInterviewAnalysisRequest;
        toObject(): {};
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): DeleteJobInterviewAnalysisRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): DeleteJobInterviewAnalysisRequest;
    }
    export class DeleteJobInterviewAnalysisResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {});
        static fromObject(data: {}): DeleteJobInterviewAnalysisResponse;
        toObject(): {};
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): DeleteJobInterviewAnalysisResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): DeleteJobInterviewAnalysisResponse;
    }
    export class GetAnalysesForInterviewsRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_application_id?: string;
        });
        get job_application_id(): string;
        set job_application_id(value: string);
        static fromObject(data: {
            job_application_id?: string;
        }): GetAnalysesForInterviewsRequest;
        toObject(): {
            job_application_id?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetAnalysesForInterviewsRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): GetAnalysesForInterviewsRequest;
    }
    export class GetAnalysesForInterviewsResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            analyses?: AnalysisForInterview[];
        });
        get analyses(): AnalysisForInterview[];
        set analyses(value: AnalysisForInterview[]);
        static fromObject(data: {
            analyses?: ReturnType<typeof AnalysisForInterview.prototype.toObject>[];
        }): GetAnalysesForInterviewsResponse;
        toObject(): {
            analyses?: ReturnType<typeof AnalysisForInterview.prototype.toObject>[];
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetAnalysesForInterviewsResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): GetAnalysesForInterviewsResponse;
    }
    export class GenerateAnalysesForInterviewsRequest extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_application_id?: string;
        });
        get job_application_id(): string;
        set job_application_id(value: string);
        static fromObject(data: {
            job_application_id?: string;
        }): GenerateAnalysesForInterviewsRequest;
        toObject(): {
            job_application_id?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GenerateAnalysesForInterviewsRequest;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): GenerateAnalysesForInterviewsRequest;
    }
    export class GenerateAnalysesForInterviewsResponse extends pb_1.Message {
        #private;
        constructor(data?: any[] | {});
        static fromObject(data: {}): GenerateAnalysesForInterviewsResponse;
        toObject(): {};
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GenerateAnalysesForInterviewsResponse;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): GenerateAnalysesForInterviewsResponse;
    }
    export class AnalysisForInterview extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            job_interview_id?: string;
            analysis?: Analysis;
        });
        get job_interview_id(): string;
        set job_interview_id(value: string);
        get analysis(): Analysis;
        set analysis(value: Analysis);
        get has_analysis(): boolean;
        static fromObject(data: {
            job_interview_id?: string;
            analysis?: ReturnType<typeof Analysis.prototype.toObject>;
        }): AnalysisForInterview;
        toObject(): {
            job_interview_id?: string;
            analysis?: ReturnType<typeof Analysis.prototype.toObject>;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AnalysisForInterview;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): AnalysisForInterview;
    }
    export class Analysis extends pb_1.Message {
        #private;
        constructor(data?: any[] | {
            last_message_id?: string;
            last_message_ts?: dependency_1.google.protobuf.Timestamp;
            analysis?: string;
        });
        get last_message_id(): string;
        set last_message_id(value: string);
        get last_message_ts(): dependency_1.google.protobuf.Timestamp;
        set last_message_ts(value: dependency_1.google.protobuf.Timestamp);
        get has_last_message_ts(): boolean;
        get analysis(): string;
        set analysis(value: string);
        static fromObject(data: {
            last_message_id?: string;
            last_message_ts?: ReturnType<typeof dependency_1.google.protobuf.Timestamp.prototype.toObject>;
            analysis?: string;
        }): Analysis;
        toObject(): {
            last_message_id?: string;
            last_message_ts?: ReturnType<typeof dependency_1.google.protobuf.Timestamp.prototype.toObject>;
            analysis?: string;
        };
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Analysis;
        serializeBinary(): Uint8Array;
        static deserializeBinary(bytes: Uint8Array): Analysis;
    }
    interface GrpcUnaryServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    }
    export abstract class UnimplementedJobInterviewAnalysisServiceService {
        static definition: {
            CreateJobInterviewAnalysis: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: CreateJobInterviewAnalysisRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => CreateJobInterviewAnalysisRequest;
                responseSerialize: (message: CreateJobInterviewAnalysisResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => CreateJobInterviewAnalysisResponse;
            };
            ReadJobInterviewAnalysis: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: ReadJobInterviewAnalysisRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => ReadJobInterviewAnalysisRequest;
                responseSerialize: (message: ReadJobInterviewAnalysisResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => ReadJobInterviewAnalysisResponse;
            };
            UpdateJobInterviewAnalysis: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: UpdateJobInterviewAnalysisRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => UpdateJobInterviewAnalysisRequest;
                responseSerialize: (message: UpdateJobInterviewAnalysisResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => UpdateJobInterviewAnalysisResponse;
            };
            DeleteJobInterviewAnalysis: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: DeleteJobInterviewAnalysisRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => DeleteJobInterviewAnalysisRequest;
                responseSerialize: (message: DeleteJobInterviewAnalysisResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => DeleteJobInterviewAnalysisResponse;
            };
            GenerateAnalysesForInterviews: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: GenerateAnalysesForInterviewsRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => GenerateAnalysesForInterviewsRequest;
                responseSerialize: (message: GenerateAnalysesForInterviewsResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => GenerateAnalysesForInterviewsResponse;
            };
            GetAnalysesForInterviews: {
                path: string;
                requestStream: boolean;
                responseStream: boolean;
                requestSerialize: (message: GetAnalysesForInterviewsRequest) => Buffer;
                requestDeserialize: (bytes: Buffer) => GetAnalysesForInterviewsRequest;
                responseSerialize: (message: GetAnalysesForInterviewsResponse) => Buffer;
                responseDeserialize: (bytes: Buffer) => GetAnalysesForInterviewsResponse;
            };
        };
        [method: string]: grpc_1.UntypedHandleCall;
        abstract CreateJobInterviewAnalysis(call: grpc_1.ServerUnaryCall<CreateJobInterviewAnalysisRequest, CreateJobInterviewAnalysisResponse>, callback: grpc_1.sendUnaryData<CreateJobInterviewAnalysisResponse>): void;
        abstract ReadJobInterviewAnalysis(call: grpc_1.ServerUnaryCall<ReadJobInterviewAnalysisRequest, ReadJobInterviewAnalysisResponse>, callback: grpc_1.sendUnaryData<ReadJobInterviewAnalysisResponse>): void;
        abstract UpdateJobInterviewAnalysis(call: grpc_1.ServerUnaryCall<UpdateJobInterviewAnalysisRequest, UpdateJobInterviewAnalysisResponse>, callback: grpc_1.sendUnaryData<UpdateJobInterviewAnalysisResponse>): void;
        abstract DeleteJobInterviewAnalysis(call: grpc_1.ServerUnaryCall<DeleteJobInterviewAnalysisRequest, DeleteJobInterviewAnalysisResponse>, callback: grpc_1.sendUnaryData<DeleteJobInterviewAnalysisResponse>): void;
        abstract GenerateAnalysesForInterviews(call: grpc_1.ServerUnaryCall<GenerateAnalysesForInterviewsRequest, GenerateAnalysesForInterviewsResponse>, callback: grpc_1.sendUnaryData<GenerateAnalysesForInterviewsResponse>): void;
        abstract GetAnalysesForInterviews(call: grpc_1.ServerUnaryCall<GetAnalysesForInterviewsRequest, GetAnalysesForInterviewsResponse>, callback: grpc_1.sendUnaryData<GetAnalysesForInterviewsResponse>): void;
    }
    const JobInterviewAnalysisServiceClient_base: grpc_1.ServiceClientConstructor;
    export class JobInterviewAnalysisServiceClient extends JobInterviewAnalysisServiceClient_base {
        constructor(address: string, credentials: grpc_1.ChannelCredentials, options?: Partial<grpc_1.ChannelOptions>);
        CreateJobInterviewAnalysis: GrpcUnaryServiceInterface<CreateJobInterviewAnalysisRequest, CreateJobInterviewAnalysisResponse>;
        ReadJobInterviewAnalysis: GrpcUnaryServiceInterface<ReadJobInterviewAnalysisRequest, ReadJobInterviewAnalysisResponse>;
        UpdateJobInterviewAnalysis: GrpcUnaryServiceInterface<UpdateJobInterviewAnalysisRequest, UpdateJobInterviewAnalysisResponse>;
        DeleteJobInterviewAnalysis: GrpcUnaryServiceInterface<DeleteJobInterviewAnalysisRequest, DeleteJobInterviewAnalysisResponse>;
        GenerateAnalysesForInterviews: GrpcUnaryServiceInterface<GenerateAnalysesForInterviewsRequest, GenerateAnalysesForInterviewsResponse>;
        GetAnalysesForInterviews: GrpcUnaryServiceInterface<GetAnalysesForInterviewsRequest, GetAnalysesForInterviewsResponse>;
    }
    export {};
}
