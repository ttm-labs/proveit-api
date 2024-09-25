"use strict";
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
var __classPrivateFieldGet = (this && this.__classPrivateFieldGet) || function (receiver, state, kind, f) {
    if (kind === "a" && !f) throw new TypeError("Private accessor was defined without a getter");
    if (typeof state === "function" ? receiver !== state || !f : !state.has(receiver)) throw new TypeError("Cannot read private member from an object whose class did not declare it");
    return kind === "m" ? f : kind === "a" ? f.call(receiver) : f ? f.value : state.get(receiver);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.job_application_service = void 0;
/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 5.27.3
 * source: job_application.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
const pb_1 = __importStar(require("google-protobuf"));
const grpc_1 = __importStar(require("@grpc/grpc-js"));
var job_application_service;
(function (job_application_service) {
    var _CreateJobApplicationRequest_one_of_decls, _CreateJobApplicationResponse_one_of_decls, _ReadJobApplicationRequest_one_of_decls, _ReadJobApplicationResponse_one_of_decls, _UpdateJobApplicationRequest_one_of_decls, _UpdateJobApplicationResponse_one_of_decls, _DeleteJobApplicationRequest_one_of_decls, _DeleteJobApplicationResponse_one_of_decls;
    class CreateJobApplicationRequest extends pb_1.Message {
        constructor(data) {
            super();
            _CreateJobApplicationRequest_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _CreateJobApplicationRequest_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new CreateJobApplicationRequest({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateJobApplicationRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return CreateJobApplicationRequest.deserialize(bytes);
        }
    }
    _CreateJobApplicationRequest_one_of_decls = new WeakMap();
    job_application_service.CreateJobApplicationRequest = CreateJobApplicationRequest;
    class CreateJobApplicationResponse extends pb_1.Message {
        constructor(data) {
            super();
            _CreateJobApplicationResponse_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _CreateJobApplicationResponse_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new CreateJobApplicationResponse({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateJobApplicationResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return CreateJobApplicationResponse.deserialize(bytes);
        }
    }
    _CreateJobApplicationResponse_one_of_decls = new WeakMap();
    job_application_service.CreateJobApplicationResponse = CreateJobApplicationResponse;
    class ReadJobApplicationRequest extends pb_1.Message {
        constructor(data) {
            super();
            _ReadJobApplicationRequest_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _ReadJobApplicationRequest_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new ReadJobApplicationRequest({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ReadJobApplicationRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return ReadJobApplicationRequest.deserialize(bytes);
        }
    }
    _ReadJobApplicationRequest_one_of_decls = new WeakMap();
    job_application_service.ReadJobApplicationRequest = ReadJobApplicationRequest;
    class ReadJobApplicationResponse extends pb_1.Message {
        constructor(data) {
            super();
            _ReadJobApplicationResponse_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _ReadJobApplicationResponse_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new ReadJobApplicationResponse({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ReadJobApplicationResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return ReadJobApplicationResponse.deserialize(bytes);
        }
    }
    _ReadJobApplicationResponse_one_of_decls = new WeakMap();
    job_application_service.ReadJobApplicationResponse = ReadJobApplicationResponse;
    class UpdateJobApplicationRequest extends pb_1.Message {
        constructor(data) {
            super();
            _UpdateJobApplicationRequest_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _UpdateJobApplicationRequest_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new UpdateJobApplicationRequest({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new UpdateJobApplicationRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return UpdateJobApplicationRequest.deserialize(bytes);
        }
    }
    _UpdateJobApplicationRequest_one_of_decls = new WeakMap();
    job_application_service.UpdateJobApplicationRequest = UpdateJobApplicationRequest;
    class UpdateJobApplicationResponse extends pb_1.Message {
        constructor(data) {
            super();
            _UpdateJobApplicationResponse_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _UpdateJobApplicationResponse_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new UpdateJobApplicationResponse({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new UpdateJobApplicationResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return UpdateJobApplicationResponse.deserialize(bytes);
        }
    }
    _UpdateJobApplicationResponse_one_of_decls = new WeakMap();
    job_application_service.UpdateJobApplicationResponse = UpdateJobApplicationResponse;
    class DeleteJobApplicationRequest extends pb_1.Message {
        constructor(data) {
            super();
            _DeleteJobApplicationRequest_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _DeleteJobApplicationRequest_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new DeleteJobApplicationRequest({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new DeleteJobApplicationRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return DeleteJobApplicationRequest.deserialize(bytes);
        }
    }
    _DeleteJobApplicationRequest_one_of_decls = new WeakMap();
    job_application_service.DeleteJobApplicationRequest = DeleteJobApplicationRequest;
    class DeleteJobApplicationResponse extends pb_1.Message {
        constructor(data) {
            super();
            _DeleteJobApplicationResponse_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _DeleteJobApplicationResponse_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") { }
        }
        static fromObject(data) {
            const message = new DeleteJobApplicationResponse({});
            return message;
        }
        toObject() {
            const data = {};
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new DeleteJobApplicationResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return DeleteJobApplicationResponse.deserialize(bytes);
        }
    }
    _DeleteJobApplicationResponse_one_of_decls = new WeakMap();
    job_application_service.DeleteJobApplicationResponse = DeleteJobApplicationResponse;
    class UnimplementedJobApplicationServiceService {
    }
    UnimplementedJobApplicationServiceService.definition = {
        CreateJobApplication: {
            path: "/job_application_service.JobApplicationService/CreateJobApplication",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => CreateJobApplicationRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => CreateJobApplicationResponse.deserialize(new Uint8Array(bytes))
        },
        ReadJobApplication: {
            path: "/job_application_service.JobApplicationService/ReadJobApplication",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => ReadJobApplicationRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => ReadJobApplicationResponse.deserialize(new Uint8Array(bytes))
        },
        UpdateJobApplication: {
            path: "/job_application_service.JobApplicationService/UpdateJobApplication",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => UpdateJobApplicationRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => UpdateJobApplicationResponse.deserialize(new Uint8Array(bytes))
        },
        DeleteJobApplication: {
            path: "/job_application_service.JobApplicationService/DeleteJobApplication",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => DeleteJobApplicationRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => DeleteJobApplicationResponse.deserialize(new Uint8Array(bytes))
        }
    };
    job_application_service.UnimplementedJobApplicationServiceService = UnimplementedJobApplicationServiceService;
    class JobApplicationServiceClient extends grpc_1.makeGenericClientConstructor(UnimplementedJobApplicationServiceService.definition, "JobApplicationService", {}) {
        constructor(address, credentials, options) {
            super(address, credentials, options);
            this.CreateJobApplication = (message, metadata, options, callback) => {
                return super.CreateJobApplication(message, metadata, options, callback);
            };
            this.ReadJobApplication = (message, metadata, options, callback) => {
                return super.ReadJobApplication(message, metadata, options, callback);
            };
            this.UpdateJobApplication = (message, metadata, options, callback) => {
                return super.UpdateJobApplication(message, metadata, options, callback);
            };
            this.DeleteJobApplication = (message, metadata, options, callback) => {
                return super.DeleteJobApplication(message, metadata, options, callback);
            };
        }
    }
    job_application_service.JobApplicationServiceClient = JobApplicationServiceClient;
})(job_application_service || (exports.job_application_service = job_application_service = {}));