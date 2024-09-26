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
exports.job_interview_service = void 0;
/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 5.27.3
 * source: job_interview.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
const pb_1 = __importStar(require("google-protobuf"));
const grpc_1 = __importStar(require("@grpc/grpc-js"));
var job_interview_service;
(function (job_interview_service) {
    var _JobInterviewRequest_one_of_decls, _JobInterviewResponse_one_of_decls;
    class JobInterviewRequest extends pb_1.Message {
        constructor(data) {
            super();
            _JobInterviewRequest_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _JobInterviewRequest_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") {
                if ("job_application_id" in data && data.job_application_id != undefined) {
                    this.job_application_id = data.job_application_id;
                }
                if ("message" in data && data.message != undefined) {
                    this.message = data.message;
                }
            }
        }
        get job_application_id() {
            return pb_1.Message.getFieldWithDefault(this, 1, "");
        }
        set job_application_id(value) {
            pb_1.Message.setField(this, 1, value);
        }
        get message() {
            return pb_1.Message.getFieldWithDefault(this, 2, "");
        }
        set message(value) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data) {
            const message = new JobInterviewRequest({});
            if (data.job_application_id != null) {
                message.job_application_id = data.job_application_id;
            }
            if (data.message != null) {
                message.message = data.message;
            }
            return message;
        }
        toObject() {
            const data = {};
            if (this.job_application_id != null) {
                data.job_application_id = this.job_application_id;
            }
            if (this.message != null) {
                data.message = this.message;
            }
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (this.job_application_id.length)
                writer.writeString(1, this.job_application_id);
            if (this.message.length)
                writer.writeString(2, this.message);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new JobInterviewRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.job_application_id = reader.readString();
                        break;
                    case 2:
                        message.message = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return JobInterviewRequest.deserialize(bytes);
        }
    }
    _JobInterviewRequest_one_of_decls = new WeakMap();
    job_interview_service.JobInterviewRequest = JobInterviewRequest;
    class JobInterviewResponse extends pb_1.Message {
        constructor(data) {
            super();
            _JobInterviewResponse_one_of_decls.set(this, []);
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], __classPrivateFieldGet(this, _JobInterviewResponse_one_of_decls, "f"));
            if (!Array.isArray(data) && typeof data == "object") {
                if ("request_id" in data && data.request_id != undefined) {
                    this.request_id = data.request_id;
                }
                if ("message" in data && data.message != undefined) {
                    this.message = data.message;
                }
            }
        }
        get request_id() {
            return pb_1.Message.getFieldWithDefault(this, 1, "");
        }
        set request_id(value) {
            pb_1.Message.setField(this, 1, value);
        }
        get message() {
            return pb_1.Message.getFieldWithDefault(this, 2, "");
        }
        set message(value) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data) {
            const message = new JobInterviewResponse({});
            if (data.request_id != null) {
                message.request_id = data.request_id;
            }
            if (data.message != null) {
                message.message = data.message;
            }
            return message;
        }
        toObject() {
            const data = {};
            if (this.request_id != null) {
                data.request_id = this.request_id;
            }
            if (this.message != null) {
                data.message = this.message;
            }
            return data;
        }
        serialize(w) {
            const writer = w || new pb_1.BinaryWriter();
            if (this.request_id.length)
                writer.writeString(1, this.request_id);
            if (this.message.length)
                writer.writeString(2, this.message);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes) {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new JobInterviewResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.request_id = reader.readString();
                        break;
                    case 2:
                        message.message = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary() {
            return this.serialize();
        }
        static deserializeBinary(bytes) {
            return JobInterviewResponse.deserialize(bytes);
        }
    }
    _JobInterviewResponse_one_of_decls = new WeakMap();
    job_interview_service.JobInterviewResponse = JobInterviewResponse;
    class UnimplementedJobInterviewServiceService {
    }
    UnimplementedJobInterviewServiceService.definition = {
        UnaryConversation: {
            path: "/job_interview_service.JobInterviewService/UnaryConversation",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => JobInterviewRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => JobInterviewResponse.deserialize(new Uint8Array(bytes))
        },
        BidirectionalConversation: {
            path: "/job_interview_service.JobInterviewService/BidirectionalConversation",
            requestStream: true,
            responseStream: true,
            requestSerialize: (message) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes) => JobInterviewRequest.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes) => JobInterviewResponse.deserialize(new Uint8Array(bytes))
        }
    };
    job_interview_service.UnimplementedJobInterviewServiceService = UnimplementedJobInterviewServiceService;
    class JobInterviewServiceClient extends grpc_1.makeGenericClientConstructor(UnimplementedJobInterviewServiceService.definition, "JobInterviewService", {}) {
        constructor(address, credentials, options) {
            super(address, credentials, options);
            this.UnaryConversation = (message, metadata, options, callback) => {
                return super.UnaryConversation(message, metadata, options, callback);
            };
            this.BidirectionalConversation = (metadata, options) => {
                return super.BidirectionalConversation(metadata, options);
            };
        }
    }
    job_interview_service.JobInterviewServiceClient = JobInterviewServiceClient;
})(job_interview_service || (exports.job_interview_service = job_interview_service = {}));
