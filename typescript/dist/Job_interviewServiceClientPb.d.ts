/**
 * @fileoverview gRPC-Web generated client stub for job_interview_service
 * @enhanceable
 * @public
 */
import * as grpcWeb from 'grpc-web';
import * as job_interview_pb from './job_interview_pb';
export declare class JobInterviewServiceClient {
    client_: grpcWeb.AbstractClientBase;
    hostname_: string;
    credentials_: null | {
        [index: string]: string;
    };
    options_: null | {
        [index: string]: any;
    };
    constructor(hostname: string, credentials?: null | {
        [index: string]: string;
    }, options?: null | {
        [index: string]: any;
    });
    methodDescriptorUnaryConversation: any;
    unaryConversation(request: job_interview_pb.JobInterviewRequest, metadata?: grpcWeb.Metadata | null): Promise<job_interview_pb.JobInterviewResponse>;
    unaryConversation(request: job_interview_pb.JobInterviewRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_interview_pb.JobInterviewResponse) => void): grpcWeb.ClientReadableStream<job_interview_pb.JobInterviewResponse>;
}
