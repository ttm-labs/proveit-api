// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: job_interview_analysis.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateJobInterviewAnalysisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobInterviewId string `protobuf:"bytes,1,opt,name=job_interview_id,json=jobInterviewId,proto3" json:"job_interview_id,omitempty"`
}

func (x *CreateJobInterviewAnalysisRequest) Reset() {
	*x = CreateJobInterviewAnalysisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobInterviewAnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobInterviewAnalysisRequest) ProtoMessage() {}

func (x *CreateJobInterviewAnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobInterviewAnalysisRequest.ProtoReflect.Descriptor instead.
func (*CreateJobInterviewAnalysisRequest) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobInterviewAnalysisRequest) GetJobInterviewId() string {
	if x != nil {
		return x.JobInterviewId
	}
	return ""
}

type CreateJobInterviewAnalysisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobInterviewAnalysisId string `protobuf:"bytes,1,opt,name=job_interview_analysis_id,json=jobInterviewAnalysisId,proto3" json:"job_interview_analysis_id,omitempty"`
}

func (x *CreateJobInterviewAnalysisResponse) Reset() {
	*x = CreateJobInterviewAnalysisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobInterviewAnalysisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobInterviewAnalysisResponse) ProtoMessage() {}

func (x *CreateJobInterviewAnalysisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobInterviewAnalysisResponse.ProtoReflect.Descriptor instead.
func (*CreateJobInterviewAnalysisResponse) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobInterviewAnalysisResponse) GetJobInterviewAnalysisId() string {
	if x != nil {
		return x.JobInterviewAnalysisId
	}
	return ""
}

type ReadJobInterviewAnalysisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobInterviewId string `protobuf:"bytes,1,opt,name=job_interview_id,json=jobInterviewId,proto3" json:"job_interview_id,omitempty"`
}

func (x *ReadJobInterviewAnalysisRequest) Reset() {
	*x = ReadJobInterviewAnalysisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobInterviewAnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobInterviewAnalysisRequest) ProtoMessage() {}

func (x *ReadJobInterviewAnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobInterviewAnalysisRequest.ProtoReflect.Descriptor instead.
func (*ReadJobInterviewAnalysisRequest) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *ReadJobInterviewAnalysisRequest) GetJobInterviewId() string {
	if x != nil {
		return x.JobInterviewId
	}
	return ""
}

type ReadJobInterviewAnalysisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Analysis *Analysis `protobuf:"bytes,1,opt,name=analysis,proto3" json:"analysis,omitempty"`
}

func (x *ReadJobInterviewAnalysisResponse) Reset() {
	*x = ReadJobInterviewAnalysisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobInterviewAnalysisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobInterviewAnalysisResponse) ProtoMessage() {}

func (x *ReadJobInterviewAnalysisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobInterviewAnalysisResponse.ProtoReflect.Descriptor instead.
func (*ReadJobInterviewAnalysisResponse) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *ReadJobInterviewAnalysisResponse) GetAnalysis() *Analysis {
	if x != nil {
		return x.Analysis
	}
	return nil
}

type UpdateJobInterviewAnalysisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobInterviewAnalysisRequest) Reset() {
	*x = UpdateJobInterviewAnalysisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobInterviewAnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobInterviewAnalysisRequest) ProtoMessage() {}

func (x *UpdateJobInterviewAnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobInterviewAnalysisRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobInterviewAnalysisRequest) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{4}
}

type UpdateJobInterviewAnalysisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobInterviewAnalysisResponse) Reset() {
	*x = UpdateJobInterviewAnalysisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobInterviewAnalysisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobInterviewAnalysisResponse) ProtoMessage() {}

func (x *UpdateJobInterviewAnalysisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobInterviewAnalysisResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobInterviewAnalysisResponse) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{5}
}

type DeleteJobInterviewAnalysisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobInterviewAnalysisRequest) Reset() {
	*x = DeleteJobInterviewAnalysisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobInterviewAnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobInterviewAnalysisRequest) ProtoMessage() {}

func (x *DeleteJobInterviewAnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobInterviewAnalysisRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobInterviewAnalysisRequest) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{6}
}

type DeleteJobInterviewAnalysisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobInterviewAnalysisResponse) Reset() {
	*x = DeleteJobInterviewAnalysisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobInterviewAnalysisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobInterviewAnalysisResponse) ProtoMessage() {}

func (x *DeleteJobInterviewAnalysisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobInterviewAnalysisResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobInterviewAnalysisResponse) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{7}
}

type GetAnalysesForInterviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobApplicationId string `protobuf:"bytes,1,opt,name=job_application_id,json=jobApplicationId,proto3" json:"job_application_id,omitempty"`
}

func (x *GetAnalysesForInterviewsRequest) Reset() {
	*x = GetAnalysesForInterviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnalysesForInterviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnalysesForInterviewsRequest) ProtoMessage() {}

func (x *GetAnalysesForInterviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnalysesForInterviewsRequest.ProtoReflect.Descriptor instead.
func (*GetAnalysesForInterviewsRequest) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{8}
}

func (x *GetAnalysesForInterviewsRequest) GetJobApplicationId() string {
	if x != nil {
		return x.JobApplicationId
	}
	return ""
}

type GetAnalysesForInterviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Analyses []*AnalysisForInterview `protobuf:"bytes,1,rep,name=analyses,proto3" json:"analyses,omitempty"`
}

func (x *GetAnalysesForInterviewsResponse) Reset() {
	*x = GetAnalysesForInterviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnalysesForInterviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnalysesForInterviewsResponse) ProtoMessage() {}

func (x *GetAnalysesForInterviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnalysesForInterviewsResponse.ProtoReflect.Descriptor instead.
func (*GetAnalysesForInterviewsResponse) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{9}
}

func (x *GetAnalysesForInterviewsResponse) GetAnalyses() []*AnalysisForInterview {
	if x != nil {
		return x.Analyses
	}
	return nil
}

type AnalysisForInterview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobInterviewId string    `protobuf:"bytes,1,opt,name=job_interview_id,json=jobInterviewId,proto3" json:"job_interview_id,omitempty"`
	Analysis       *Analysis `protobuf:"bytes,2,opt,name=analysis,proto3" json:"analysis,omitempty"`
}

func (x *AnalysisForInterview) Reset() {
	*x = AnalysisForInterview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisForInterview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisForInterview) ProtoMessage() {}

func (x *AnalysisForInterview) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisForInterview.ProtoReflect.Descriptor instead.
func (*AnalysisForInterview) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{10}
}

func (x *AnalysisForInterview) GetJobInterviewId() string {
	if x != nil {
		return x.JobInterviewId
	}
	return ""
}

func (x *AnalysisForInterview) GetAnalysis() *Analysis {
	if x != nil {
		return x.Analysis
	}
	return nil
}

type Analysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastMessageId string                 `protobuf:"bytes,1,opt,name=last_message_id,json=lastMessageId,proto3" json:"last_message_id,omitempty"`
	LastMessageTs *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_message_ts,json=lastMessageTs,proto3" json:"last_message_ts,omitempty"`
	Analysis      string                 `protobuf:"bytes,3,opt,name=analysis,proto3" json:"analysis,omitempty"`
}

func (x *Analysis) Reset() {
	*x = Analysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_interview_analysis_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Analysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Analysis) ProtoMessage() {}

func (x *Analysis) ProtoReflect() protoreflect.Message {
	mi := &file_job_interview_analysis_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Analysis.ProtoReflect.Descriptor instead.
func (*Analysis) Descriptor() ([]byte, []int) {
	return file_job_interview_analysis_proto_rawDescGZIP(), []int{11}
}

func (x *Analysis) GetLastMessageId() string {
	if x != nil {
		return x.LastMessageId
	}
	return ""
}

func (x *Analysis) GetLastMessageTs() *timestamppb.Timestamp {
	if x != nil {
		return x.LastMessageTs
	}
	return nil
}

func (x *Analysis) GetAnalysis() string {
	if x != nil {
		return x.Analysis
	}
	return ""
}

var File_job_interview_analysis_proto protoreflect.FileDescriptor

var file_job_interview_analysis_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4d, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x5f,
	0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x49, 0x64, 0x22,
	0x4b, 0x0a, 0x1f, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f,
	0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x20,
	0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x08, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x22, 0x23, 0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x22, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x22, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x6f, 0x62,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x74, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x46, 0x6f, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x46, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x28, 0x0a, 0x10,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x32, 0xd9, 0x06, 0x0a, 0x1b, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa5, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x12, 0x41, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x18, 0x52, 0x65,
	0x61, 0x64, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x3f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x41, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x12, 0x41, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x3f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_interview_analysis_proto_rawDescOnce sync.Once
	file_job_interview_analysis_proto_rawDescData = file_job_interview_analysis_proto_rawDesc
)

func file_job_interview_analysis_proto_rawDescGZIP() []byte {
	file_job_interview_analysis_proto_rawDescOnce.Do(func() {
		file_job_interview_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_interview_analysis_proto_rawDescData)
	})
	return file_job_interview_analysis_proto_rawDescData
}

var file_job_interview_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_job_interview_analysis_proto_goTypes = []any{
	(*CreateJobInterviewAnalysisRequest)(nil),  // 0: job_interview_analysis_service.CreateJobInterviewAnalysisRequest
	(*CreateJobInterviewAnalysisResponse)(nil), // 1: job_interview_analysis_service.CreateJobInterviewAnalysisResponse
	(*ReadJobInterviewAnalysisRequest)(nil),    // 2: job_interview_analysis_service.ReadJobInterviewAnalysisRequest
	(*ReadJobInterviewAnalysisResponse)(nil),   // 3: job_interview_analysis_service.ReadJobInterviewAnalysisResponse
	(*UpdateJobInterviewAnalysisRequest)(nil),  // 4: job_interview_analysis_service.UpdateJobInterviewAnalysisRequest
	(*UpdateJobInterviewAnalysisResponse)(nil), // 5: job_interview_analysis_service.UpdateJobInterviewAnalysisResponse
	(*DeleteJobInterviewAnalysisRequest)(nil),  // 6: job_interview_analysis_service.DeleteJobInterviewAnalysisRequest
	(*DeleteJobInterviewAnalysisResponse)(nil), // 7: job_interview_analysis_service.DeleteJobInterviewAnalysisResponse
	(*GetAnalysesForInterviewsRequest)(nil),    // 8: job_interview_analysis_service.GetAnalysesForInterviewsRequest
	(*GetAnalysesForInterviewsResponse)(nil),   // 9: job_interview_analysis_service.GetAnalysesForInterviewsResponse
	(*AnalysisForInterview)(nil),               // 10: job_interview_analysis_service.AnalysisForInterview
	(*Analysis)(nil),                           // 11: job_interview_analysis_service.Analysis
	(*timestamppb.Timestamp)(nil),              // 12: google.protobuf.Timestamp
}
var file_job_interview_analysis_proto_depIdxs = []int32{
	11, // 0: job_interview_analysis_service.ReadJobInterviewAnalysisResponse.analysis:type_name -> job_interview_analysis_service.Analysis
	10, // 1: job_interview_analysis_service.GetAnalysesForInterviewsResponse.analyses:type_name -> job_interview_analysis_service.AnalysisForInterview
	11, // 2: job_interview_analysis_service.AnalysisForInterview.analysis:type_name -> job_interview_analysis_service.Analysis
	12, // 3: job_interview_analysis_service.Analysis.last_message_ts:type_name -> google.protobuf.Timestamp
	0,  // 4: job_interview_analysis_service.JobInterviewAnalysisService.CreateJobInterviewAnalysis:input_type -> job_interview_analysis_service.CreateJobInterviewAnalysisRequest
	2,  // 5: job_interview_analysis_service.JobInterviewAnalysisService.ReadJobInterviewAnalysis:input_type -> job_interview_analysis_service.ReadJobInterviewAnalysisRequest
	4,  // 6: job_interview_analysis_service.JobInterviewAnalysisService.UpdateJobInterviewAnalysis:input_type -> job_interview_analysis_service.UpdateJobInterviewAnalysisRequest
	6,  // 7: job_interview_analysis_service.JobInterviewAnalysisService.DeleteJobInterviewAnalysis:input_type -> job_interview_analysis_service.DeleteJobInterviewAnalysisRequest
	8,  // 8: job_interview_analysis_service.JobInterviewAnalysisService.GetAnalysesForInterviews:input_type -> job_interview_analysis_service.GetAnalysesForInterviewsRequest
	1,  // 9: job_interview_analysis_service.JobInterviewAnalysisService.CreateJobInterviewAnalysis:output_type -> job_interview_analysis_service.CreateJobInterviewAnalysisResponse
	3,  // 10: job_interview_analysis_service.JobInterviewAnalysisService.ReadJobInterviewAnalysis:output_type -> job_interview_analysis_service.ReadJobInterviewAnalysisResponse
	5,  // 11: job_interview_analysis_service.JobInterviewAnalysisService.UpdateJobInterviewAnalysis:output_type -> job_interview_analysis_service.UpdateJobInterviewAnalysisResponse
	7,  // 12: job_interview_analysis_service.JobInterviewAnalysisService.DeleteJobInterviewAnalysis:output_type -> job_interview_analysis_service.DeleteJobInterviewAnalysisResponse
	9,  // 13: job_interview_analysis_service.JobInterviewAnalysisService.GetAnalysesForInterviews:output_type -> job_interview_analysis_service.GetAnalysesForInterviewsResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_job_interview_analysis_proto_init() }
func file_job_interview_analysis_proto_init() {
	if File_job_interview_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_interview_analysis_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobInterviewAnalysisRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobInterviewAnalysisResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobInterviewAnalysisRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobInterviewAnalysisResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobInterviewAnalysisRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobInterviewAnalysisResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobInterviewAnalysisRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobInterviewAnalysisResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAnalysesForInterviewsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetAnalysesForInterviewsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*AnalysisForInterview); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_interview_analysis_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Analysis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_interview_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_interview_analysis_proto_goTypes,
		DependencyIndexes: file_job_interview_analysis_proto_depIdxs,
		MessageInfos:      file_job_interview_analysis_proto_msgTypes,
	}.Build()
	File_job_interview_analysis_proto = out.File
	file_job_interview_analysis_proto_rawDesc = nil
	file_job_interview_analysis_proto_goTypes = nil
	file_job_interview_analysis_proto_depIdxs = nil
}
