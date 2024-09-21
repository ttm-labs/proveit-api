/**
 * @fileoverview gRPC-Web generated client stub for job_interview_service
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.3
// source: job_interview.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as job_interview_pb from './job_interview_pb'; // proto import: "job_interview.proto"


export class JobInterviewServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorUnaryConversation = new grpcWeb.MethodDescriptor(
    '/job_interview_service.JobInterviewService/UnaryConversation',
    grpcWeb.MethodType.UNARY,
    job_interview_pb.JobInterviewRequest,
    job_interview_pb.JobInterviewResponse,
    (request: job_interview_pb.JobInterviewRequest) => {
      return request.serializeBinary();
    },
    job_interview_pb.JobInterviewResponse.deserializeBinary
  );

  unaryConversation(
    request: job_interview_pb.JobInterviewRequest,
    metadata?: grpcWeb.Metadata | null): Promise<job_interview_pb.JobInterviewResponse>;

  unaryConversation(
    request: job_interview_pb.JobInterviewRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: job_interview_pb.JobInterviewResponse) => void): grpcWeb.ClientReadableStream<job_interview_pb.JobInterviewResponse>;

  unaryConversation(
    request: job_interview_pb.JobInterviewRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: job_interview_pb.JobInterviewResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/job_interview_service.JobInterviewService/UnaryConversation',
        request,
        metadata || {},
        this.methodDescriptorUnaryConversation,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/job_interview_service.JobInterviewService/UnaryConversation',
    request,
    metadata || {},
    this.methodDescriptorUnaryConversation);
  }

}
