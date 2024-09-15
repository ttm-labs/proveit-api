// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: job_posting.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateJobPostingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Industry       string   `protobuf:"bytes,1,opt,name=industry,proto3" json:"industry,omitempty"`
	CompanyName    string   `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Location       string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Title          string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	BaseSalary     *Salary  `protobuf:"bytes,5,opt,name=base_salary,json=baseSalary,proto3" json:"base_salary,omitempty"`
	BonusSalary    *Salary  `protobuf:"bytes,6,opt,name=bonus_salary,json=bonusSalary,proto3" json:"bonus_salary,omitempty"`
	Qualifications []string `protobuf:"bytes,7,rep,name=qualifications,proto3" json:"qualifications,omitempty"`
}

func (x *CreateJobPostingRequest) Reset() {
	*x = CreateJobPostingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobPostingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobPostingRequest) ProtoMessage() {}

func (x *CreateJobPostingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobPostingRequest.ProtoReflect.Descriptor instead.
func (*CreateJobPostingRequest) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobPostingRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *CreateJobPostingRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CreateJobPostingRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateJobPostingRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateJobPostingRequest) GetBaseSalary() *Salary {
	if x != nil {
		return x.BaseSalary
	}
	return nil
}

func (x *CreateJobPostingRequest) GetBonusSalary() *Salary {
	if x != nil {
		return x.BonusSalary
	}
	return nil
}

func (x *CreateJobPostingRequest) GetQualifications() []string {
	if x != nil {
		return x.Qualifications
	}
	return nil
}

type CreateJobPostingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId string `protobuf:"bytes,1,opt,name=job_posting_id,json=jobPostingId,proto3" json:"job_posting_id,omitempty"`
}

func (x *CreateJobPostingResponse) Reset() {
	*x = CreateJobPostingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobPostingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobPostingResponse) ProtoMessage() {}

func (x *CreateJobPostingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobPostingResponse.ProtoReflect.Descriptor instead.
func (*CreateJobPostingResponse) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobPostingResponse) GetJobPostingId() string {
	if x != nil {
		return x.JobPostingId
	}
	return ""
}

type ReadJobPostingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPosting_1 string `protobuf:"bytes,1,opt,name=job_posting_1,json=jobPosting1,proto3" json:"job_posting_1,omitempty"`
}

func (x *ReadJobPostingRequest) Reset() {
	*x = ReadJobPostingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobPostingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobPostingRequest) ProtoMessage() {}

func (x *ReadJobPostingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobPostingRequest.ProtoReflect.Descriptor instead.
func (*ReadJobPostingRequest) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{2}
}

func (x *ReadJobPostingRequest) GetJobPosting_1() string {
	if x != nil {
		return x.JobPosting_1
	}
	return ""
}

type ReadJobPostingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Industry       string   `protobuf:"bytes,1,opt,name=industry,proto3" json:"industry,omitempty"`
	CompanyName    string   `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Location       string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Title          string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	BaseSalary     *Salary  `protobuf:"bytes,5,opt,name=base_salary,json=baseSalary,proto3" json:"base_salary,omitempty"`
	BonusSalary    *Salary  `protobuf:"bytes,6,opt,name=bonus_salary,json=bonusSalary,proto3" json:"bonus_salary,omitempty"`
	Qualifications []string `protobuf:"bytes,7,rep,name=qualifications,proto3" json:"qualifications,omitempty"`
}

func (x *ReadJobPostingResponse) Reset() {
	*x = ReadJobPostingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobPostingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobPostingResponse) ProtoMessage() {}

func (x *ReadJobPostingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobPostingResponse.ProtoReflect.Descriptor instead.
func (*ReadJobPostingResponse) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{3}
}

func (x *ReadJobPostingResponse) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *ReadJobPostingResponse) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *ReadJobPostingResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ReadJobPostingResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReadJobPostingResponse) GetBaseSalary() *Salary {
	if x != nil {
		return x.BaseSalary
	}
	return nil
}

func (x *ReadJobPostingResponse) GetBonusSalary() *Salary {
	if x != nil {
		return x.BonusSalary
	}
	return nil
}

func (x *ReadJobPostingResponse) GetQualifications() []string {
	if x != nil {
		return x.Qualifications
	}
	return nil
}

type UpdateJobPostingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId   string   `protobuf:"bytes,1,opt,name=job_posting_id,json=jobPostingId,proto3" json:"job_posting_id,omitempty"`
	Industry       string   `protobuf:"bytes,2,opt,name=industry,proto3" json:"industry,omitempty"`
	CompanyName    string   `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Location       string   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Title          string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	BaseSalary     *Salary  `protobuf:"bytes,6,opt,name=base_salary,json=baseSalary,proto3" json:"base_salary,omitempty"`
	BonusSalary    *Salary  `protobuf:"bytes,7,opt,name=bonus_salary,json=bonusSalary,proto3" json:"bonus_salary,omitempty"`
	Qualifications []string `protobuf:"bytes,8,rep,name=qualifications,proto3" json:"qualifications,omitempty"`
}

func (x *UpdateJobPostingRequest) Reset() {
	*x = UpdateJobPostingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobPostingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobPostingRequest) ProtoMessage() {}

func (x *UpdateJobPostingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobPostingRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobPostingRequest) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJobPostingRequest) GetJobPostingId() string {
	if x != nil {
		return x.JobPostingId
	}
	return ""
}

func (x *UpdateJobPostingRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *UpdateJobPostingRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *UpdateJobPostingRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateJobPostingRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateJobPostingRequest) GetBaseSalary() *Salary {
	if x != nil {
		return x.BaseSalary
	}
	return nil
}

func (x *UpdateJobPostingRequest) GetBonusSalary() *Salary {
	if x != nil {
		return x.BonusSalary
	}
	return nil
}

func (x *UpdateJobPostingRequest) GetQualifications() []string {
	if x != nil {
		return x.Qualifications
	}
	return nil
}

type UpdateJobPostingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobPostingResponse) Reset() {
	*x = UpdateJobPostingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobPostingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobPostingResponse) ProtoMessage() {}

func (x *UpdateJobPostingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobPostingResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobPostingResponse) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{5}
}

type DeleteJobPostingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId string `protobuf:"bytes,1,opt,name=job_posting_id,json=jobPostingId,proto3" json:"job_posting_id,omitempty"`
}

func (x *DeleteJobPostingRequest) Reset() {
	*x = DeleteJobPostingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobPostingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobPostingRequest) ProtoMessage() {}

func (x *DeleteJobPostingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobPostingRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobPostingRequest) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteJobPostingRequest) GetJobPostingId() string {
	if x != nil {
		return x.JobPostingId
	}
	return ""
}

type DeleteJobPostingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobPostingResponse) Reset() {
	*x = DeleteJobPostingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobPostingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobPostingResponse) ProtoMessage() {}

func (x *DeleteJobPostingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobPostingResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobPostingResponse) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{7}
}

type Salary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinRange int32 `protobuf:"varint,1,opt,name=min_range,json=minRange,proto3" json:"min_range,omitempty"`
	MaxRange int32 `protobuf:"varint,2,opt,name=max_range,json=maxRange,proto3" json:"max_range,omitempty"`
}

func (x *Salary) Reset() {
	*x = Salary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_posting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Salary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Salary) ProtoMessage() {}

func (x *Salary) ProtoReflect() protoreflect.Message {
	mi := &file_job_posting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Salary.ProtoReflect.Descriptor instead.
func (*Salary) Descriptor() ([]byte, []int) {
	return file_job_posting_proto_rawDescGZIP(), []int{8}
}

func (x *Salary) GetMinRange() int32 {
	if x != nil {
		return x.MinRange
	}
	return 0
}

func (x *Salary) GetMaxRange() int32 {
	if x != nil {
		return x.MaxRange
	}
	return 0
}

var File_job_posting_proto protoreflect.FileDescriptor

var file_job_posting_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb0, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x0b, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x15, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a,
	0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x31, 0x22, 0xaf, 0x02, 0x0a, 0x16, 0x52,
	0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x0b, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd6, 0x02, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x3c, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x3e, 0x0a,
	0x0c, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x52, 0x0b, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a,
	0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42,
	0x0a, 0x06, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x32, 0xd9, 0x03, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0e, 0x52,
	0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x2c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_posting_proto_rawDescOnce sync.Once
	file_job_posting_proto_rawDescData = file_job_posting_proto_rawDesc
)

func file_job_posting_proto_rawDescGZIP() []byte {
	file_job_posting_proto_rawDescOnce.Do(func() {
		file_job_posting_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_posting_proto_rawDescData)
	})
	return file_job_posting_proto_rawDescData
}

var file_job_posting_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_job_posting_proto_goTypes = []any{
	(*CreateJobPostingRequest)(nil),  // 0: job_posting_service.CreateJobPostingRequest
	(*CreateJobPostingResponse)(nil), // 1: job_posting_service.CreateJobPostingResponse
	(*ReadJobPostingRequest)(nil),    // 2: job_posting_service.ReadJobPostingRequest
	(*ReadJobPostingResponse)(nil),   // 3: job_posting_service.ReadJobPostingResponse
	(*UpdateJobPostingRequest)(nil),  // 4: job_posting_service.UpdateJobPostingRequest
	(*UpdateJobPostingResponse)(nil), // 5: job_posting_service.UpdateJobPostingResponse
	(*DeleteJobPostingRequest)(nil),  // 6: job_posting_service.DeleteJobPostingRequest
	(*DeleteJobPostingResponse)(nil), // 7: job_posting_service.DeleteJobPostingResponse
	(*Salary)(nil),                   // 8: job_posting_service.Salary
}
var file_job_posting_proto_depIdxs = []int32{
	8,  // 0: job_posting_service.CreateJobPostingRequest.base_salary:type_name -> job_posting_service.Salary
	8,  // 1: job_posting_service.CreateJobPostingRequest.bonus_salary:type_name -> job_posting_service.Salary
	8,  // 2: job_posting_service.ReadJobPostingResponse.base_salary:type_name -> job_posting_service.Salary
	8,  // 3: job_posting_service.ReadJobPostingResponse.bonus_salary:type_name -> job_posting_service.Salary
	8,  // 4: job_posting_service.UpdateJobPostingRequest.base_salary:type_name -> job_posting_service.Salary
	8,  // 5: job_posting_service.UpdateJobPostingRequest.bonus_salary:type_name -> job_posting_service.Salary
	0,  // 6: job_posting_service.JobPostingService.CreateJobPosting:input_type -> job_posting_service.CreateJobPostingRequest
	2,  // 7: job_posting_service.JobPostingService.ReadJobPosting:input_type -> job_posting_service.ReadJobPostingRequest
	4,  // 8: job_posting_service.JobPostingService.UpdateJobPosting:input_type -> job_posting_service.UpdateJobPostingRequest
	6,  // 9: job_posting_service.JobPostingService.DeleteJobPosting:input_type -> job_posting_service.DeleteJobPostingRequest
	1,  // 10: job_posting_service.JobPostingService.CreateJobPosting:output_type -> job_posting_service.CreateJobPostingResponse
	3,  // 11: job_posting_service.JobPostingService.ReadJobPosting:output_type -> job_posting_service.ReadJobPostingResponse
	5,  // 12: job_posting_service.JobPostingService.UpdateJobPosting:output_type -> job_posting_service.UpdateJobPostingResponse
	7,  // 13: job_posting_service.JobPostingService.DeleteJobPosting:output_type -> job_posting_service.DeleteJobPostingResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_job_posting_proto_init() }
func file_job_posting_proto_init() {
	if File_job_posting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_posting_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobPostingRequest); i {
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
		file_job_posting_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobPostingResponse); i {
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
		file_job_posting_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobPostingRequest); i {
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
		file_job_posting_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobPostingResponse); i {
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
		file_job_posting_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobPostingRequest); i {
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
		file_job_posting_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobPostingResponse); i {
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
		file_job_posting_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobPostingRequest); i {
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
		file_job_posting_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobPostingResponse); i {
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
		file_job_posting_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Salary); i {
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
			RawDescriptor: file_job_posting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_posting_proto_goTypes,
		DependencyIndexes: file_job_posting_proto_depIdxs,
		MessageInfos:      file_job_posting_proto_msgTypes,
	}.Build()
	File_job_posting_proto = out.File
	file_job_posting_proto_rawDesc = nil
	file_job_posting_proto_goTypes = nil
	file_job_posting_proto_depIdxs = nil
}
