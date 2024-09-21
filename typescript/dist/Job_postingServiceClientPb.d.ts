/**
 * @fileoverview gRPC-Web generated client stub for job_posting_service
 * @enhanceable
 * @public
 */
import * as grpcWeb from 'grpc-web';
import * as job_posting_pb from './job_posting_pb';
export declare class JobPostingServiceClient {
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
    methodDescriptorCreateJobPosting: any;
    createJobPosting(request: job_posting_pb.CreateJobPostingRequest, metadata?: grpcWeb.Metadata | null): Promise<job_posting_pb.CreateJobPostingResponse>;
    createJobPosting(request: job_posting_pb.CreateJobPostingRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_posting_pb.CreateJobPostingResponse) => void): grpcWeb.ClientReadableStream<job_posting_pb.CreateJobPostingResponse>;
    methodDescriptorReadJobPosting: any;
    readJobPosting(request: job_posting_pb.ReadJobPostingRequest, metadata?: grpcWeb.Metadata | null): Promise<job_posting_pb.ReadJobPostingResponse>;
    readJobPosting(request: job_posting_pb.ReadJobPostingRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_posting_pb.ReadJobPostingResponse) => void): grpcWeb.ClientReadableStream<job_posting_pb.ReadJobPostingResponse>;
    methodDescriptorUpdateJobPosting: any;
    updateJobPosting(request: job_posting_pb.UpdateJobPostingRequest, metadata?: grpcWeb.Metadata | null): Promise<job_posting_pb.UpdateJobPostingResponse>;
    updateJobPosting(request: job_posting_pb.UpdateJobPostingRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_posting_pb.UpdateJobPostingResponse) => void): grpcWeb.ClientReadableStream<job_posting_pb.UpdateJobPostingResponse>;
    methodDescriptorDeleteJobPosting: any;
    deleteJobPosting(request: job_posting_pb.DeleteJobPostingRequest, metadata?: grpcWeb.Metadata | null): Promise<job_posting_pb.DeleteJobPostingResponse>;
    deleteJobPosting(request: job_posting_pb.DeleteJobPostingRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_posting_pb.DeleteJobPostingResponse) => void): grpcWeb.ClientReadableStream<job_posting_pb.DeleteJobPostingResponse>;
}
