/**
 * @fileoverview gRPC-Web generated client stub for job_application_service
 * @enhanceable
 * @public
 */
import * as grpcWeb from 'grpc-web';
import * as job_application_pb from './job_application_pb';
export declare class JobApplicationServiceClient {
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
    methodDescriptorCreateJobApplication: any;
    createJobApplication(request: job_application_pb.CreateJobApplicationRequest, metadata?: grpcWeb.Metadata | null): Promise<job_application_pb.CreateJobApplicationResponse>;
    createJobApplication(request: job_application_pb.CreateJobApplicationRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_application_pb.CreateJobApplicationResponse) => void): grpcWeb.ClientReadableStream<job_application_pb.CreateJobApplicationResponse>;
    methodDescriptorReadJobApplication: any;
    readJobApplication(request: job_application_pb.ReadJobApplicationRequest, metadata?: grpcWeb.Metadata | null): Promise<job_application_pb.ReadJobApplicationResponse>;
    readJobApplication(request: job_application_pb.ReadJobApplicationRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_application_pb.ReadJobApplicationResponse) => void): grpcWeb.ClientReadableStream<job_application_pb.ReadJobApplicationResponse>;
    methodDescriptorUpdateJobApplication: any;
    updateJobApplication(request: job_application_pb.UpdateJobApplicationRequest, metadata?: grpcWeb.Metadata | null): Promise<job_application_pb.UpdateJobApplicationResponse>;
    updateJobApplication(request: job_application_pb.UpdateJobApplicationRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_application_pb.UpdateJobApplicationResponse) => void): grpcWeb.ClientReadableStream<job_application_pb.UpdateJobApplicationResponse>;
    methodDescriptorDeleteJobApplication: any;
    deleteJobApplication(request: job_application_pb.DeleteJobApplicationRequest, metadata?: grpcWeb.Metadata | null): Promise<job_application_pb.DeleteJobApplicationResponse>;
    deleteJobApplication(request: job_application_pb.DeleteJobApplicationRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: job_application_pb.DeleteJobApplicationResponse) => void): grpcWeb.ClientReadableStream<job_application_pb.DeleteJobApplicationResponse>;
}
