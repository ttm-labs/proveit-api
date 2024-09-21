/**
 * @fileoverview gRPC-Web generated client stub for job_posting_service
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.3
// source: job_posting.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.job_posting_service = require('./job_posting_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.job_posting_service.JobPostingServiceClient =
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
proto.job_posting_service.JobPostingServicePromiseClient =
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
 *   !proto.job_posting_service.CreateJobPostingRequest,
 *   !proto.job_posting_service.CreateJobPostingResponse>}
 */
const methodDescriptor_JobPostingService_CreateJobPosting = new grpc.web.MethodDescriptor(
  '/job_posting_service.JobPostingService/CreateJobPosting',
  grpc.web.MethodType.UNARY,
  proto.job_posting_service.CreateJobPostingRequest,
  proto.job_posting_service.CreateJobPostingResponse,
  /**
   * @param {!proto.job_posting_service.CreateJobPostingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.job_posting_service.CreateJobPostingResponse.deserializeBinary
);


/**
 * @param {!proto.job_posting_service.CreateJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.job_posting_service.CreateJobPostingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.job_posting_service.CreateJobPostingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.job_posting_service.JobPostingServiceClient.prototype.createJobPosting =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/job_posting_service.JobPostingService/CreateJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_CreateJobPosting,
      callback);
};


/**
 * @param {!proto.job_posting_service.CreateJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.job_posting_service.CreateJobPostingResponse>}
 *     Promise that resolves to the response
 */
proto.job_posting_service.JobPostingServicePromiseClient.prototype.createJobPosting =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/job_posting_service.JobPostingService/CreateJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_CreateJobPosting);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.job_posting_service.ReadJobPostingRequest,
 *   !proto.job_posting_service.ReadJobPostingResponse>}
 */
const methodDescriptor_JobPostingService_ReadJobPosting = new grpc.web.MethodDescriptor(
  '/job_posting_service.JobPostingService/ReadJobPosting',
  grpc.web.MethodType.UNARY,
  proto.job_posting_service.ReadJobPostingRequest,
  proto.job_posting_service.ReadJobPostingResponse,
  /**
   * @param {!proto.job_posting_service.ReadJobPostingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.job_posting_service.ReadJobPostingResponse.deserializeBinary
);


/**
 * @param {!proto.job_posting_service.ReadJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.job_posting_service.ReadJobPostingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.job_posting_service.ReadJobPostingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.job_posting_service.JobPostingServiceClient.prototype.readJobPosting =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/job_posting_service.JobPostingService/ReadJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_ReadJobPosting,
      callback);
};


/**
 * @param {!proto.job_posting_service.ReadJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.job_posting_service.ReadJobPostingResponse>}
 *     Promise that resolves to the response
 */
proto.job_posting_service.JobPostingServicePromiseClient.prototype.readJobPosting =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/job_posting_service.JobPostingService/ReadJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_ReadJobPosting);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.job_posting_service.UpdateJobPostingRequest,
 *   !proto.job_posting_service.UpdateJobPostingResponse>}
 */
const methodDescriptor_JobPostingService_UpdateJobPosting = new grpc.web.MethodDescriptor(
  '/job_posting_service.JobPostingService/UpdateJobPosting',
  grpc.web.MethodType.UNARY,
  proto.job_posting_service.UpdateJobPostingRequest,
  proto.job_posting_service.UpdateJobPostingResponse,
  /**
   * @param {!proto.job_posting_service.UpdateJobPostingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.job_posting_service.UpdateJobPostingResponse.deserializeBinary
);


/**
 * @param {!proto.job_posting_service.UpdateJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.job_posting_service.UpdateJobPostingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.job_posting_service.UpdateJobPostingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.job_posting_service.JobPostingServiceClient.prototype.updateJobPosting =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/job_posting_service.JobPostingService/UpdateJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_UpdateJobPosting,
      callback);
};


/**
 * @param {!proto.job_posting_service.UpdateJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.job_posting_service.UpdateJobPostingResponse>}
 *     Promise that resolves to the response
 */
proto.job_posting_service.JobPostingServicePromiseClient.prototype.updateJobPosting =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/job_posting_service.JobPostingService/UpdateJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_UpdateJobPosting);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.job_posting_service.DeleteJobPostingRequest,
 *   !proto.job_posting_service.DeleteJobPostingResponse>}
 */
const methodDescriptor_JobPostingService_DeleteJobPosting = new grpc.web.MethodDescriptor(
  '/job_posting_service.JobPostingService/DeleteJobPosting',
  grpc.web.MethodType.UNARY,
  proto.job_posting_service.DeleteJobPostingRequest,
  proto.job_posting_service.DeleteJobPostingResponse,
  /**
   * @param {!proto.job_posting_service.DeleteJobPostingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.job_posting_service.DeleteJobPostingResponse.deserializeBinary
);


/**
 * @param {!proto.job_posting_service.DeleteJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.job_posting_service.DeleteJobPostingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.job_posting_service.DeleteJobPostingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.job_posting_service.JobPostingServiceClient.prototype.deleteJobPosting =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/job_posting_service.JobPostingService/DeleteJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_DeleteJobPosting,
      callback);
};


/**
 * @param {!proto.job_posting_service.DeleteJobPostingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.job_posting_service.DeleteJobPostingResponse>}
 *     Promise that resolves to the response
 */
proto.job_posting_service.JobPostingServicePromiseClient.prototype.deleteJobPosting =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/job_posting_service.JobPostingService/DeleteJobPosting',
      request,
      metadata || {},
      methodDescriptor_JobPostingService_DeleteJobPosting);
};


module.exports = proto.job_posting_service;
