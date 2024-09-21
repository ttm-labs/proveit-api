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



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.job_interview_service = require('./job_interview_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.job_interview_service.JobInterviewServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.job_interview_service.JobInterviewServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.job_interview_service.JobInterviewRequest,
 *   !proto.job_interview_service.JobInterviewResponse>}
 */
const methodDescriptor_JobInterviewService_UnaryConversation = new grpc.web.MethodDescriptor(
  '/job_interview_service.JobInterviewService/UnaryConversation',
  grpc.web.MethodType.UNARY,
  proto.job_interview_service.JobInterviewRequest,
  proto.job_interview_service.JobInterviewResponse,
  /**
   * @param {!proto.job_interview_service.JobInterviewRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.job_interview_service.JobInterviewResponse.deserializeBinary
);


/**
 * @param {!proto.job_interview_service.JobInterviewRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.job_interview_service.JobInterviewResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.job_interview_service.JobInterviewResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.job_interview_service.JobInterviewServiceClient.prototype.unaryConversation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/job_interview_service.JobInterviewService/UnaryConversation',
      request,
      metadata || {},
      methodDescriptor_JobInterviewService_UnaryConversation,
      callback);
};


/**
 * @param {!proto.job_interview_service.JobInterviewRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.job_interview_service.JobInterviewResponse>}
 *     Promise that resolves to the response
 */
proto.job_interview_service.JobInterviewServicePromiseClient.prototype.unaryConversation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/job_interview_service.JobInterviewService/UnaryConversation',
      request,
      metadata || {},
      methodDescriptor_JobInterviewService_UnaryConversation);
};


module.exports = proto.job_interview_service;
