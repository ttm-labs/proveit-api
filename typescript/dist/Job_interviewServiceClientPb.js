"use strict";
/**
 * @fileoverview gRPC-Web generated client stub for job_interview_service
 * @enhanceable
 * @public
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.JobInterviewServiceClient = void 0;
// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.3
// source: job_interview.proto
/* eslint-disable */
// @ts-nocheck
const grpcWeb = __importStar(require("grpc-web"));
const job_interview_pb = __importStar(require("./job_interview_pb")); // proto import: "job_interview.proto"
class JobInterviewServiceClient {
    constructor(hostname, credentials, options) {
        this.methodDescriptorUnaryConversation = new grpcWeb.MethodDescriptor('/job_interview_service.JobInterviewService/UnaryConversation', grpcWeb.MethodType.UNARY, job_interview_pb.JobInterviewRequest, job_interview_pb.JobInterviewResponse, (request) => {
            return request.serializeBinary();
        }, job_interview_pb.JobInterviewResponse.deserializeBinary);
        if (!options)
            options = {};
        if (!credentials)
            credentials = {};
        options['format'] = 'text';
        this.client_ = new grpcWeb.GrpcWebClientBase(options);
        this.hostname_ = hostname.replace(/\/+$/, '');
        this.credentials_ = credentials;
        this.options_ = options;
    }
    unaryConversation(request, metadata, callback) {
        if (callback !== undefined) {
            return this.client_.rpcCall(this.hostname_ +
                '/job_interview_service.JobInterviewService/UnaryConversation', request, metadata || {}, this.methodDescriptorUnaryConversation, callback);
        }
        return this.client_.unaryCall(this.hostname_ +
            '/job_interview_service.JobInterviewService/UnaryConversation', request, metadata || {}, this.methodDescriptorUnaryConversation);
    }
}
exports.JobInterviewServiceClient = JobInterviewServiceClient;
