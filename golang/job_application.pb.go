// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: job_application.proto

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

type CreateJobApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId string   `protobuf:"bytes,1,opt,name=job_posting_id,json=jobPostingId,proto3" json:"job_posting_id,omitempty"`
	Name         *Name    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contact      *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	Info         *Info    `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateJobApplicationRequest) Reset() {
	*x = CreateJobApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobApplicationRequest) ProtoMessage() {}

func (x *CreateJobApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobApplicationRequest.ProtoReflect.Descriptor instead.
func (*CreateJobApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobApplicationRequest) GetJobPostingId() string {
	if x != nil {
		return x.JobPostingId
	}
	return ""
}

func (x *CreateJobApplicationRequest) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CreateJobApplicationRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *CreateJobApplicationRequest) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateJobApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobApplicationId string `protobuf:"bytes,1,opt,name=job_application_id,json=jobApplicationId,proto3" json:"job_application_id,omitempty"`
}

func (x *CreateJobApplicationResponse) Reset() {
	*x = CreateJobApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobApplicationResponse) ProtoMessage() {}

func (x *CreateJobApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobApplicationResponse.ProtoReflect.Descriptor instead.
func (*CreateJobApplicationResponse) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobApplicationResponse) GetJobApplicationId() string {
	if x != nil {
		return x.JobApplicationId
	}
	return ""
}

type ReadJobApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobApplicationId string `protobuf:"bytes,1,opt,name=job_application_id,json=jobApplicationId,proto3" json:"job_application_id,omitempty"`
}

func (x *ReadJobApplicationRequest) Reset() {
	*x = ReadJobApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobApplicationRequest) ProtoMessage() {}

func (x *ReadJobApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobApplicationRequest.ProtoReflect.Descriptor instead.
func (*ReadJobApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{2}
}

func (x *ReadJobApplicationRequest) GetJobApplicationId() string {
	if x != nil {
		return x.JobApplicationId
	}
	return ""
}

type ReadJobApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId string   `protobuf:"bytes,1,opt,name=job_posting_id,json=jobPostingId,proto3" json:"job_posting_id,omitempty"`
	Name         *Name    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contact      *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	Info         *Info    `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Status       *Status  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReadJobApplicationResponse) Reset() {
	*x = ReadJobApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadJobApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadJobApplicationResponse) ProtoMessage() {}

func (x *ReadJobApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadJobApplicationResponse.ProtoReflect.Descriptor instead.
func (*ReadJobApplicationResponse) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{3}
}

func (x *ReadJobApplicationResponse) GetJobPostingId() string {
	if x != nil {
		return x.JobPostingId
	}
	return ""
}

func (x *ReadJobApplicationResponse) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ReadJobApplicationResponse) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *ReadJobApplicationResponse) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReadJobApplicationResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type UpdateJobApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobApplicationId string   `protobuf:"bytes,1,opt,name=job_application_id,json=jobApplicationId,proto3" json:"job_application_id,omitempty"`
	Name             *Name    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contact          *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Status           *Status  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateJobApplicationRequest) Reset() {
	*x = UpdateJobApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobApplicationRequest) ProtoMessage() {}

func (x *UpdateJobApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobApplicationRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJobApplicationRequest) GetJobApplicationId() string {
	if x != nil {
		return x.JobApplicationId
	}
	return ""
}

func (x *UpdateJobApplicationRequest) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateJobApplicationRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *UpdateJobApplicationRequest) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *UpdateJobApplicationRequest) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type UpdateJobApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobApplicationResponse) Reset() {
	*x = UpdateJobApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobApplicationResponse) ProtoMessage() {}

func (x *UpdateJobApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobApplicationResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobApplicationResponse) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{5}
}

type DeleteJobApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobApplicationId string `protobuf:"bytes,1,opt,name=job_application_id,json=jobApplicationId,proto3" json:"job_application_id,omitempty"`
}

func (x *DeleteJobApplicationRequest) Reset() {
	*x = DeleteJobApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobApplicationRequest) ProtoMessage() {}

func (x *DeleteJobApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobApplicationRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteJobApplicationRequest) GetJobApplicationId() string {
	if x != nil {
		return x.JobApplicationId
	}
	return ""
}

type DeleteJobApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobApplicationResponse) Reset() {
	*x = DeleteJobApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobApplicationResponse) ProtoMessage() {}

func (x *DeleteJobApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobApplicationResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobApplicationResponse) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{7}
}

// Reused, nested messages
type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{8}
}

func (x *Name) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Name) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{9}
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Introduction string `protobuf:"bytes,1,opt,name=introduction,proto3" json:"introduction,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{10}
}

func (x *Info) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_application_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{11}
}

var File_job_application_proto protoreflect.FileDescriptor

var file_job_application_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0xe5, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f,
	0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x9d, 0x02, 0x0a, 0x1a, 0x52, 0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xa6, 0x02, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a,
	0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x31,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x1b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x08,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb0, 0x04, 0x0a, 0x15, 0x4a, 0x6f, 0x62,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x35, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x12, 0x52, 0x65,
	0x61, 0x64, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4a,
	0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_application_proto_rawDescOnce sync.Once
	file_job_application_proto_rawDescData = file_job_application_proto_rawDesc
)

func file_job_application_proto_rawDescGZIP() []byte {
	file_job_application_proto_rawDescOnce.Do(func() {
		file_job_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_application_proto_rawDescData)
	})
	return file_job_application_proto_rawDescData
}

var file_job_application_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_job_application_proto_goTypes = []any{
	(*CreateJobApplicationRequest)(nil),  // 0: job_application_service.CreateJobApplicationRequest
	(*CreateJobApplicationResponse)(nil), // 1: job_application_service.CreateJobApplicationResponse
	(*ReadJobApplicationRequest)(nil),    // 2: job_application_service.ReadJobApplicationRequest
	(*ReadJobApplicationResponse)(nil),   // 3: job_application_service.ReadJobApplicationResponse
	(*UpdateJobApplicationRequest)(nil),  // 4: job_application_service.UpdateJobApplicationRequest
	(*UpdateJobApplicationResponse)(nil), // 5: job_application_service.UpdateJobApplicationResponse
	(*DeleteJobApplicationRequest)(nil),  // 6: job_application_service.DeleteJobApplicationRequest
	(*DeleteJobApplicationResponse)(nil), // 7: job_application_service.DeleteJobApplicationResponse
	(*Name)(nil),                         // 8: job_application_service.Name
	(*Contact)(nil),                      // 9: job_application_service.Contact
	(*Info)(nil),                         // 10: job_application_service.Info
	(*Status)(nil),                       // 11: job_application_service.Status
}
var file_job_application_proto_depIdxs = []int32{
	8,  // 0: job_application_service.CreateJobApplicationRequest.name:type_name -> job_application_service.Name
	9,  // 1: job_application_service.CreateJobApplicationRequest.contact:type_name -> job_application_service.Contact
	10, // 2: job_application_service.CreateJobApplicationRequest.info:type_name -> job_application_service.Info
	8,  // 3: job_application_service.ReadJobApplicationResponse.name:type_name -> job_application_service.Name
	9,  // 4: job_application_service.ReadJobApplicationResponse.contact:type_name -> job_application_service.Contact
	10, // 5: job_application_service.ReadJobApplicationResponse.info:type_name -> job_application_service.Info
	11, // 6: job_application_service.ReadJobApplicationResponse.status:type_name -> job_application_service.Status
	8,  // 7: job_application_service.UpdateJobApplicationRequest.name:type_name -> job_application_service.Name
	9,  // 8: job_application_service.UpdateJobApplicationRequest.contact:type_name -> job_application_service.Contact
	10, // 9: job_application_service.UpdateJobApplicationRequest.info:type_name -> job_application_service.Info
	11, // 10: job_application_service.UpdateJobApplicationRequest.status:type_name -> job_application_service.Status
	0,  // 11: job_application_service.JobApplicationService.CreateJobApplication:input_type -> job_application_service.CreateJobApplicationRequest
	2,  // 12: job_application_service.JobApplicationService.ReadJobApplication:input_type -> job_application_service.ReadJobApplicationRequest
	4,  // 13: job_application_service.JobApplicationService.UpdateJobApplication:input_type -> job_application_service.UpdateJobApplicationRequest
	6,  // 14: job_application_service.JobApplicationService.DeleteJobApplication:input_type -> job_application_service.DeleteJobApplicationRequest
	1,  // 15: job_application_service.JobApplicationService.CreateJobApplication:output_type -> job_application_service.CreateJobApplicationResponse
	3,  // 16: job_application_service.JobApplicationService.ReadJobApplication:output_type -> job_application_service.ReadJobApplicationResponse
	5,  // 17: job_application_service.JobApplicationService.UpdateJobApplication:output_type -> job_application_service.UpdateJobApplicationResponse
	7,  // 18: job_application_service.JobApplicationService.DeleteJobApplication:output_type -> job_application_service.DeleteJobApplicationResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_job_application_proto_init() }
func file_job_application_proto_init() {
	if File_job_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_application_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobApplicationRequest); i {
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
		file_job_application_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateJobApplicationResponse); i {
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
		file_job_application_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobApplicationRequest); i {
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
		file_job_application_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadJobApplicationResponse); i {
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
		file_job_application_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobApplicationRequest); i {
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
		file_job_application_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateJobApplicationResponse); i {
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
		file_job_application_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobApplicationRequest); i {
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
		file_job_application_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteJobApplicationResponse); i {
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
		file_job_application_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Name); i {
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
		file_job_application_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Contact); i {
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
		file_job_application_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Info); i {
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
		file_job_application_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_job_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_application_proto_goTypes,
		DependencyIndexes: file_job_application_proto_depIdxs,
		MessageInfos:      file_job_application_proto_msgTypes,
	}.Build()
	File_job_application_proto = out.File
	file_job_application_proto_rawDesc = nil
	file_job_application_proto_goTypes = nil
	file_job_application_proto_depIdxs = nil
}
