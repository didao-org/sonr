// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: api.proto

package models

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

// Client Network Connection Status
type Connectivity int32

const (
	Connectivity_None   Connectivity = 0
	Connectivity_Mobile Connectivity = 1
	Connectivity_WiFi   Connectivity = 2
)

// Enum value maps for Connectivity.
var (
	Connectivity_name = map[int32]string{
		0: "None",
		1: "Mobile",
		2: "WiFi",
	}
	Connectivity_value = map[string]int32{
		"None":   0,
		"Mobile": 1,
		"WiFi":   2,
	}
)

func (x Connectivity) Enum() *Connectivity {
	p := new(Connectivity)
	*p = x
	return p
}

func (x Connectivity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Connectivity) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Connectivity) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Connectivity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Connectivity.Descriptor instead.
func (Connectivity) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

// Status is User Node Situation
type Status int32

const (
	Status_NONE         Status = 0 // Default Status on Launch
	Status_CONNECTED    Status = 1 // Status after starting Host
	Status_BOOTSTRAPPED Status = 2 // After Bootstrapping Host
	Status_AVAILABLE    Status = 3 // Connected and Visible on Local Lobby
	Status_SEARCHING    Status = 4 // File has been Processed ready to Invite
	Status_PENDING      Status = 5 // Awaiting Peer Authorization
	Status_INVITED      Status = 6 // Has received Invitation
	Status_INPROGRESS   Status = 7 // In Middle of Transfer
	Status_STANDBY      Status = 8 // Away from Sonr Application
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "NONE",
		1: "CONNECTED",
		2: "BOOTSTRAPPED",
		3: "AVAILABLE",
		4: "SEARCHING",
		5: "PENDING",
		6: "INVITED",
		7: "INPROGRESS",
		8: "STANDBY",
	}
	Status_value = map[string]int32{
		"NONE":         0,
		"CONNECTED":    1,
		"BOOTSTRAPPED": 2,
		"AVAILABLE":    3,
		"SEARCHING":    4,
		"PENDING":      5,
		"INVITED":      6,
		"INPROGRESS":   7,
		"STANDBY":      8,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

// Enum for Type of Transfer
type InviteRequest_TransferType int32

const (
	InviteRequest_Unknown    InviteRequest_TransferType = 0 // Default Value, Would not Transfer
	InviteRequest_File       InviteRequest_TransferType = 1 // Media or Document
	InviteRequest_MultiFiles InviteRequest_TransferType = 2 // Multiple Media or Documents
	InviteRequest_Contact    InviteRequest_TransferType = 3 // Users Contact Card
	InviteRequest_URL        InviteRequest_TransferType = 4 // Website URL Link
)

// Enum value maps for InviteRequest_TransferType.
var (
	InviteRequest_TransferType_name = map[int32]string{
		0: "Unknown",
		1: "File",
		2: "MultiFiles",
		3: "Contact",
		4: "URL",
	}
	InviteRequest_TransferType_value = map[string]int32{
		"Unknown":    0,
		"File":       1,
		"MultiFiles": 2,
		"Contact":    3,
		"URL":        4,
	}
)

func (x InviteRequest_TransferType) Enum() *InviteRequest_TransferType {
	p := new(InviteRequest_TransferType)
	*p = x
	return p
}

func (x InviteRequest_TransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InviteRequest_TransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[2].Descriptor()
}

func (InviteRequest_TransferType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[2]
}

func (x InviteRequest_TransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InviteRequest_TransferType.Descriptor instead.
func (InviteRequest_TransferType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2, 0}
}

type AuthReply_Type int32

const (
	AuthReply_None     AuthReply_Type = 0
	AuthReply_Transfer AuthReply_Type = 1
	AuthReply_Contact  AuthReply_Type = 2
	AuthReply_Cancel   AuthReply_Type = 3
)

// Enum value maps for AuthReply_Type.
var (
	AuthReply_Type_name = map[int32]string{
		0: "None",
		1: "Transfer",
		2: "Contact",
		3: "Cancel",
	}
	AuthReply_Type_value = map[string]int32{
		"None":     0,
		"Transfer": 1,
		"Contact":  2,
		"Cancel":   3,
	}
)

func (x AuthReply_Type) Enum() *AuthReply_Type {
	p := new(AuthReply_Type)
	*p = x
	return p
}

func (x AuthReply_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthReply_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[3].Descriptor()
}

func (AuthReply_Type) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[3]
}

func (x AuthReply_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthReply_Type.Descriptor instead.
func (AuthReply_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4, 0}
}

// Message for Status Update
type StatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value Status `protobuf:"varint,1,opt,name=value,proto3,enum=Status" json:"value,omitempty"`
}

func (x *StatusUpdate) Reset() {
	*x = StatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdate) ProtoMessage() {}

func (x *StatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdate.ProtoReflect.Descriptor instead.
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *StatusUpdate) GetValue() Status {
	if x != nil {
		return x.Value
	}
	return Status_NONE
}

// Initial Connection Message to Establish Sonr
type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude     float64      `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`                          // Geolocation for OLC
	Longitude    float64      `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`                        // Geolocation for OLC
	Username     string       `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`                            // General Peer Info
	Device       *Device      `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`                                // Users Device
	Directories  *Directories `protobuf:"bytes,5,opt,name=directories,proto3" json:"directories,omitempty"`                      // Users Available File Paths
	Contact      *Contact     `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`                              // Users Contact Card
	User         *User        `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`                                    // User Saved Data
	Connectivity Connectivity `protobuf:"varint,8,opt,name=connectivity,proto3,enum=Connectivity" json:"connectivity,omitempty"` // Client Network Type
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ConnectionRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ConnectionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ConnectionRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *ConnectionRequest) GetDirectories() *Directories {
	if x != nil {
		return x.Directories
	}
	return nil
}

func (x *ConnectionRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *ConnectionRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ConnectionRequest) GetConnectivity() Connectivity {
	if x != nil {
		return x.Connectivity
	}
	return Connectivity_None
}

// Processes Given File and Invites Specified Peer
type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Properties
	Type     InviteRequest_TransferType `protobuf:"varint,1,opt,name=type,proto3,enum=InviteRequest_TransferType" json:"type,omitempty"` // General Payload Type
	To       *Peer                      `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`                                      // User thats being Invited
	IsRemote bool                       `protobuf:"varint,3,opt,name=isRemote,proto3" json:"isRemote,omitempty"`                         // If Transfer is Remote
	Topic    string                     `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`                                // Remote Topic Name
	// Attached Data
	Url     string                    `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`         // URL to be Sent - Optional
	Contact *Contact                  `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"` // Contact Card to be Sent - Optional
	Files   []*InviteRequest_FileInfo `protobuf:"bytes,7,rep,name=files,proto3" json:"files,omitempty"`     // Files to be Sent - Optional
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest.ProtoReflect.Descriptor instead.
func (*InviteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *InviteRequest) GetType() InviteRequest_TransferType {
	if x != nil {
		return x.Type
	}
	return InviteRequest_Unknown
}

func (x *InviteRequest) GetTo() *Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteRequest) GetIsRemote() bool {
	if x != nil {
		return x.IsRemote
	}
	return false
}

func (x *InviteRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *InviteRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InviteRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *InviteRequest) GetFiles() []*InviteRequest_FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

// Invitation Message sent on RPC
type AuthInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload Payload       `protobuf:"varint,1,opt,name=payload,proto3,enum=Payload" json:"payload,omitempty"` // Type of Transfer
	From    *Peer         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`                     // Users Peer Data
	Card    *TransferCard `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`                     // Card contains all Data Info, Transfer Info
}

func (x *AuthInvite) Reset() {
	*x = AuthInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInvite) ProtoMessage() {}

func (x *AuthInvite) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInvite.ProtoReflect.Descriptor instead.
func (*AuthInvite) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AuthInvite) GetPayload() Payload {
	if x != nil {
		return x.Payload
	}
	return Payload_UNDEFINED
}

func (x *AuthInvite) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthInvite) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

// Reply Message sent on RPC
type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     AuthReply_Type `protobuf:"varint,1,opt,name=type,proto3,enum=AuthReply_Type" json:"type,omitempty"` // Type of Transfer Reply
	From     *Peer          `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`                      // Users Peer Data
	Decision bool           `protobuf:"varint,3,opt,name=decision,proto3" json:"decision,omitempty"`             // Users Decision for the Invitation
	Card     *TransferCard  `protobuf:"bytes,4,opt,name=card,proto3" json:"card,omitempty"`                      // Card contains all Data Info, Transfer Info
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *AuthReply) GetType() AuthReply_Type {
	if x != nil {
		return x.Type
	}
	return AuthReply_None
}

func (x *AuthReply) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *AuthReply) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *AuthReply) GetCard() *TransferCard {
	if x != nil {
		return x.Card
	}
	return nil
}

// Error Message returned from Core Library
type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Generated Error Message
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`   // Method Error Occurred
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorMessage) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// FileInfo Message for Attached Files
type InviteRequest_FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`           // Location of File
	Thumbnail []byte `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"` // Provided Thumbnail Data
	Duration  int32  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`  // Duration of Video - Optional
	IsVideo   bool   `protobuf:"varint,4,opt,name=isVideo,proto3" json:"isVideo,omitempty"`    // If File is Video
}

func (x *InviteRequest_FileInfo) Reset() {
	*x = InviteRequest_FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest_FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest_FileInfo) ProtoMessage() {}

func (x *InviteRequest_FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest_FileInfo.ProtoReflect.Descriptor instead.
func (*InviteRequest_FileInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2, 0}
}

func (x *InviteRequest_FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *InviteRequest_FileInfo) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *InviteRequest_FileInfo) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *InviteRequest_FileInfo) GetIsVideo() bool {
	if x != nil {
		return x.IsVideo
	}
	return false
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x22, 0xaf, 0x03, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x72, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x10, 0x04, 0x22, 0x6e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64,
	0x22, 0x37, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0x03, 0x22, 0x40, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2a, 0x2e, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x69, 0x46, 0x69, 0x10, 0x02, 0x2a, 0x88, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x42, 0x4f, 0x4f, 0x54, 0x53, 0x54, 0x52, 0x41, 0x50, 0x50, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41,
	0x4e, 0x44, 0x42, 0x59, 0x10, 0x08, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(Connectivity)(0),               // 0: Connectivity
	(Status)(0),                     // 1: Status
	(InviteRequest_TransferType)(0), // 2: InviteRequest.TransferType
	(AuthReply_Type)(0),             // 3: AuthReply.Type
	(*StatusUpdate)(nil),            // 4: StatusUpdate
	(*ConnectionRequest)(nil),       // 5: ConnectionRequest
	(*InviteRequest)(nil),           // 6: InviteRequest
	(*AuthInvite)(nil),              // 7: AuthInvite
	(*AuthReply)(nil),               // 8: AuthReply
	(*ErrorMessage)(nil),            // 9: ErrorMessage
	(*InviteRequest_FileInfo)(nil),  // 10: InviteRequest.FileInfo
	(*Device)(nil),                  // 11: Device
	(*Directories)(nil),             // 12: Directories
	(*Contact)(nil),                 // 13: Contact
	(*User)(nil),                    // 14: User
	(*Peer)(nil),                    // 15: Peer
	(Payload)(0),                    // 16: Payload
	(*TransferCard)(nil),            // 17: TransferCard
}
var file_api_proto_depIdxs = []int32{
	1,  // 0: StatusUpdate.value:type_name -> Status
	11, // 1: ConnectionRequest.device:type_name -> Device
	12, // 2: ConnectionRequest.directories:type_name -> Directories
	13, // 3: ConnectionRequest.contact:type_name -> Contact
	14, // 4: ConnectionRequest.user:type_name -> User
	0,  // 5: ConnectionRequest.connectivity:type_name -> Connectivity
	2,  // 6: InviteRequest.type:type_name -> InviteRequest.TransferType
	15, // 7: InviteRequest.to:type_name -> Peer
	13, // 8: InviteRequest.contact:type_name -> Contact
	10, // 9: InviteRequest.files:type_name -> InviteRequest.FileInfo
	16, // 10: AuthInvite.payload:type_name -> Payload
	15, // 11: AuthInvite.from:type_name -> Peer
	17, // 12: AuthInvite.card:type_name -> TransferCard
	3,  // 13: AuthReply.type:type_name -> AuthReply.Type
	15, // 14: AuthReply.from:type_name -> Peer
	17, // 15: AuthReply.card:type_name -> TransferCard
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_data_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusUpdate); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInvite); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest_FileInfo); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
