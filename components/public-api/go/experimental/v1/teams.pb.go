// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/teams.proto

package v1

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

type TeamRole int32

const (
	// TEAM_ROLE_UNKNOWN is the unkwnon state.
	TeamRole_TEAM_ROLE_UNSPECIFIED TeamRole = 0
	// TEAM_ROLE_OWNER is the owner of the team.
	// A team can have multiple owners, but there must always be at least one owner.
	TeamRole_TEAM_ROLE_OWNER TeamRole = 1
	// TEAM_ROLE_MEMBER is a regular member of a team.
	TeamRole_TEAM_ROLE_MEMBER TeamRole = 2
)

// Enum value maps for TeamRole.
var (
	TeamRole_name = map[int32]string{
		0: "TEAM_ROLE_UNSPECIFIED",
		1: "TEAM_ROLE_OWNER",
		2: "TEAM_ROLE_MEMBER",
	}
	TeamRole_value = map[string]int32{
		"TEAM_ROLE_UNSPECIFIED": 0,
		"TEAM_ROLE_OWNER":       1,
		"TEAM_ROLE_MEMBER":      2,
	}
)

func (x TeamRole) Enum() *TeamRole {
	p := new(TeamRole)
	*p = x
	return p
}

func (x TeamRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeamRole) Descriptor() protoreflect.EnumDescriptor {
	return file_gitpod_experimental_v1_teams_proto_enumTypes[0].Descriptor()
}

func (TeamRole) Type() protoreflect.EnumType {
	return &file_gitpod_experimental_v1_teams_proto_enumTypes[0]
}

func (x TeamRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TeamRole.Descriptor instead.
func (TeamRole) EnumDescriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{0}
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a UUID of the Team
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name is the name of the Team
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// slug is the short version of the Team name
	Slug string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	// members are the team members of this Team
	Members []*TeamMember `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	// team_invitation is the team invitation.
	TeamInvitation *TeamInvitation `protobuf:"bytes,5,opt,name=team_invitation,json=teamInvitation,proto3" json:"team_invitation,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Team) GetMembers() []*TeamMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Team) GetTeamInvitation() *TeamInvitation {
	if x != nil {
		return x.TeamInvitation
	}
	return nil
}

type TeamMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id is the identifier of the user
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// role is the role this member is assigned
	Role TeamRole `protobuf:"varint,2,opt,name=role,proto3,enum=gitpod.experimental.v1.TeamRole" json:"role,omitempty"`
}

func (x *TeamMember) Reset() {
	*x = TeamMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMember) ProtoMessage() {}

func (x *TeamMember) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMember.ProtoReflect.Descriptor instead.
func (*TeamMember) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{1}
}

func (x *TeamMember) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TeamMember) GetRole() TeamRole {
	if x != nil {
		return x.Role
	}
	return TeamRole_TEAM_ROLE_UNSPECIFIED
}

type TeamInvitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the invitation ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamInvitation) Reset() {
	*x = TeamInvitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamInvitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamInvitation) ProtoMessage() {}

func (x *TeamInvitation) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamInvitation.ProtoReflect.Descriptor instead.
func (*TeamInvitation) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{2}
}

func (x *TeamInvitation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the team name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *CreateTeamResponse) Reset() {
	*x = CreateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamResponse) ProtoMessage() {}

func (x *CreateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_teams_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_teams_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTeamResponse) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

var File_gitpod_experimental_v1_teams_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_teams_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xcd, 0x01, 0x0a,
	0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x3c, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x0a,
	0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x65, 0x61,
	0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x2a, 0x50, 0x0a, 0x08,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x41, 0x4d,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x41, 0x4d,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x32, 0x75,
	0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x29, 0x2e, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69,
	0x74, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_teams_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_teams_proto_rawDescData = file_gitpod_experimental_v1_teams_proto_rawDesc
)

func file_gitpod_experimental_v1_teams_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_teams_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_teams_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_teams_proto_rawDescData
}

var file_gitpod_experimental_v1_teams_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gitpod_experimental_v1_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gitpod_experimental_v1_teams_proto_goTypes = []interface{}{
	(TeamRole)(0),              // 0: gitpod.experimental.v1.TeamRole
	(*Team)(nil),               // 1: gitpod.experimental.v1.Team
	(*TeamMember)(nil),         // 2: gitpod.experimental.v1.TeamMember
	(*TeamInvitation)(nil),     // 3: gitpod.experimental.v1.TeamInvitation
	(*CreateTeamRequest)(nil),  // 4: gitpod.experimental.v1.CreateTeamRequest
	(*CreateTeamResponse)(nil), // 5: gitpod.experimental.v1.CreateTeamResponse
}
var file_gitpod_experimental_v1_teams_proto_depIdxs = []int32{
	2, // 0: gitpod.experimental.v1.Team.members:type_name -> gitpod.experimental.v1.TeamMember
	3, // 1: gitpod.experimental.v1.Team.team_invitation:type_name -> gitpod.experimental.v1.TeamInvitation
	0, // 2: gitpod.experimental.v1.TeamMember.role:type_name -> gitpod.experimental.v1.TeamRole
	1, // 3: gitpod.experimental.v1.CreateTeamResponse.team:type_name -> gitpod.experimental.v1.Team
	4, // 4: gitpod.experimental.v1.TeamsService.CreateTeam:input_type -> gitpod.experimental.v1.CreateTeamRequest
	5, // 5: gitpod.experimental.v1.TeamsService.CreateTeam:output_type -> gitpod.experimental.v1.CreateTeamResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_teams_proto_init() }
func file_gitpod_experimental_v1_teams_proto_init() {
	if File_gitpod_experimental_v1_teams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_teams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_gitpod_experimental_v1_teams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMember); i {
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
		file_gitpod_experimental_v1_teams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamInvitation); i {
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
		file_gitpod_experimental_v1_teams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_gitpod_experimental_v1_teams_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamResponse); i {
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
			RawDescriptor: file_gitpod_experimental_v1_teams_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_teams_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_teams_proto_depIdxs,
		EnumInfos:         file_gitpod_experimental_v1_teams_proto_enumTypes,
		MessageInfos:      file_gitpod_experimental_v1_teams_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_teams_proto = out.File
	file_gitpod_experimental_v1_teams_proto_rawDesc = nil
	file_gitpod_experimental_v1_teams_proto_goTypes = nil
	file_gitpod_experimental_v1_teams_proto_depIdxs = nil
}