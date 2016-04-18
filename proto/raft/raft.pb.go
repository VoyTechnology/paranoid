// Code generated by protoc-gen-go.
// source: raft/raft.proto
// DO NOT EDIT!

/*
Package raft is a generated protocol buffer package.

It is generated from these files:
	raft/raft.proto

It has these top-level messages:
	EmptyMessage
	EntryRequest
	AppendEntriesRequest
	StateMachineCommand
	KeyStateMessage
	Node
	Configuration
	DemoCommand
	Entry
	LogEntry
	AppendEntriesResponse
	RequestVoteRequest
	RequestVoteResponse
	SnapshotRequest
	SnapshotResponse
*/
package raft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Configuration_ConfigurationType int32

const (
	Configuration_CurrentConfiguration Configuration_ConfigurationType = 0
	Configuration_FutureConfiguration  Configuration_ConfigurationType = 1
)

var Configuration_ConfigurationType_name = map[int32]string{
	0: "CurrentConfiguration",
	1: "FutureConfiguration",
}
var Configuration_ConfigurationType_value = map[string]int32{
	"CurrentConfiguration": 0,
	"FutureConfiguration":  1,
}

func (x Configuration_ConfigurationType) String() string {
	return proto.EnumName(Configuration_ConfigurationType_name, int32(x))
}
func (Configuration_ConfigurationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type Entry_EntryType int32

const (
	Entry_StateMachineCommand Entry_EntryType = 0
	Entry_ConfigurationChange Entry_EntryType = 1
	Entry_Demo                Entry_EntryType = 2
	Entry_KeyStateMessage     Entry_EntryType = 3
)

var Entry_EntryType_name = map[int32]string{
	0: "StateMachineCommand",
	1: "ConfigurationChange",
	2: "Demo",
	3: "KeyStateMessage",
}
var Entry_EntryType_value = map[string]int32{
	"StateMachineCommand": 0,
	"ConfigurationChange": 1,
	"Demo":                2,
	"KeyStateMessage":     3,
}

func (x Entry_EntryType) String() string {
	return proto.EnumName(Entry_EntryType_name, int32(x))
}
func (Entry_EntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EntryRequest struct {
	SenderId string `protobuf:"bytes,1,opt,name=sender_id" json:"sender_id,omitempty"`
	Entry    *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *EntryRequest) Reset()                    { *m = EntryRequest{} }
func (m *EntryRequest) String() string            { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()               {}
func (*EntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EntryRequest) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AppendEntriesRequest struct {
	Term         uint64   `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	LeaderId     string   `protobuf:"bytes,2,opt,name=leader_id" json:"leader_id,omitempty"`
	PrevLogIndex uint64   `protobuf:"varint,3,opt,name=prev_log_index" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64   `protobuf:"varint,4,opt,name=prev_log_term" json:"prev_log_term,omitempty"`
	Entries      []*Entry `protobuf:"bytes,5,rep,name=entries" json:"entries,omitempty"`
	LeaderCommit uint64   `protobuf:"varint,6,opt,name=leader_commit" json:"leader_commit,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type StateMachineCommand struct {
	Type uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Used for Write command
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	// Used for Write and Truncate commands
	Length int64 `protobuf:"varint,5,opt,name=length" json:"length,omitempty"`
	// Used for Link and Rename commands
	OldPath string `protobuf:"bytes,6,opt,name=old_path" json:"old_path,omitempty"`
	NewPath string `protobuf:"bytes,7,opt,name=new_path" json:"new_path,omitempty"`
	// Used for Create, Chmod and Mkdir commands
	Mode uint32 `protobuf:"varint,8,opt,name=mode" json:"mode,omitempty"`
	// Used for Utimes command
	AccessSeconds     int64 `protobuf:"varint,9,opt,name=access_seconds" json:"access_seconds,omitempty"`
	AccessNanoseconds int64 `protobuf:"varint,10,opt,name=access_nanoseconds" json:"access_nanoseconds,omitempty"`
	ModifySeconds     int64 `protobuf:"varint,11,opt,name=modify_seconds" json:"modify_seconds,omitempty"`
	ModifyNanoseconds int64 `protobuf:"varint,12,opt,name=modify_nanoseconds" json:"modify_nanoseconds,omitempty"`
}

func (m *StateMachineCommand) Reset()                    { *m = StateMachineCommand{} }
func (m *StateMachineCommand) String() string            { return proto.CompactTextString(m) }
func (*StateMachineCommand) ProtoMessage()               {}
func (*StateMachineCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type KeyStateMessage struct {
	KeyOwner   *Node `protobuf:"bytes,1,opt,name=key_owner" json:"key_owner,omitempty"`
	KeyHolder  *Node `protobuf:"bytes,2,opt,name=key_holder" json:"key_holder,omitempty"`
	Generation int64 `protobuf:"varint,3,opt,name=generation" json:"generation,omitempty"`
}

func (m *KeyStateMessage) Reset()                    { *m = KeyStateMessage{} }
func (m *KeyStateMessage) String() string            { return proto.CompactTextString(m) }
func (*KeyStateMessage) ProtoMessage()               {}
func (*KeyStateMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KeyStateMessage) GetKeyOwner() *Node {
	if m != nil {
		return m.KeyOwner
	}
	return nil
}

func (m *KeyStateMessage) GetKeyHolder() *Node {
	if m != nil {
		return m.KeyHolder
	}
	return nil
}

type Node struct {
	Ip         string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port       string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name" json:"common_name,omitempty"`
	NodeId     string `protobuf:"bytes,4,opt,name=node_id" json:"node_id,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Configuration struct {
	Type  Configuration_ConfigurationType `protobuf:"varint,1,opt,name=type,enum=raft.Configuration_ConfigurationType" json:"type,omitempty"`
	Nodes []*Node                         `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Configuration) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type DemoCommand struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *DemoCommand) Reset()                    { *m = DemoCommand{} }
func (m *DemoCommand) String() string            { return proto.CompactTextString(m) }
func (*DemoCommand) ProtoMessage()               {}
func (*DemoCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Entry struct {
	Type      Entry_EntryType      `protobuf:"varint,1,opt,name=type,enum=raft.Entry_EntryType" json:"type,omitempty"`
	Uuid      string               `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Command   *StateMachineCommand `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
	Config    *Configuration       `protobuf:"bytes,4,opt,name=config" json:"config,omitempty"`
	Demo      *DemoCommand         `protobuf:"bytes,5,opt,name=demo" json:"demo,omitempty"`
	KeyChange *KeyStateMessage     `protobuf:"bytes,6,opt,name=key_change" json:"key_change,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Entry) GetCommand() *StateMachineCommand {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Entry) GetConfig() *Configuration {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Entry) GetDemo() *DemoCommand {
	if m != nil {
		return m.Demo
	}
	return nil
}

func (m *Entry) GetKeyChange() *KeyStateMessage {
	if m != nil {
		return m.KeyChange
	}
	return nil
}

type LogEntry struct {
	Term  uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Entry *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LogEntry) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AppendEntriesResponse struct {
	Term      uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	NextIndex uint64 `protobuf:"varint,2,opt,name=next_index" json:"next_index,omitempty"`
	Success   bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type RequestVoteRequest struct {
	Term         uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	CandidateId  string `protobuf:"bytes,2,opt,name=candidate_id" json:"candidate_id,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=last_log_index" json:"last_log_index,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=last_log_term" json:"last_log_term,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type RequestVoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=vote_granted" json:"vote_granted,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SnapshotRequest struct {
	Term              uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	LeaderId          string `protobuf:"bytes,2,opt,name=leader_id" json:"leader_id,omitempty"`
	LastIncludedIndex uint64 `protobuf:"varint,3,opt,name=last_included_index" json:"last_included_index,omitempty"`
	LastIncludedTerm  uint64 `protobuf:"varint,4,opt,name=last_included_term" json:"last_included_term,omitempty"`
	Offset            uint64 `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Data              []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Done              bool   `protobuf:"varint,7,opt,name=done" json:"done,omitempty"`
}

func (m *SnapshotRequest) Reset()                    { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()               {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SnapshotResponse struct {
	Term uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "raft.EmptyMessage")
	proto.RegisterType((*EntryRequest)(nil), "raft.EntryRequest")
	proto.RegisterType((*AppendEntriesRequest)(nil), "raft.AppendEntriesRequest")
	proto.RegisterType((*StateMachineCommand)(nil), "raft.StateMachineCommand")
	proto.RegisterType((*KeyStateMessage)(nil), "raft.KeyStateMessage")
	proto.RegisterType((*Node)(nil), "raft.Node")
	proto.RegisterType((*Configuration)(nil), "raft.Configuration")
	proto.RegisterType((*DemoCommand)(nil), "raft.DemoCommand")
	proto.RegisterType((*Entry)(nil), "raft.Entry")
	proto.RegisterType((*LogEntry)(nil), "raft.LogEntry")
	proto.RegisterType((*AppendEntriesResponse)(nil), "raft.AppendEntriesResponse")
	proto.RegisterType((*RequestVoteRequest)(nil), "raft.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "raft.RequestVoteResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "raft.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "raft.SnapshotResponse")
	proto.RegisterEnum("raft.Configuration_ConfigurationType", Configuration_ConfigurationType_name, Configuration_ConfigurationType_value)
	proto.RegisterEnum("raft.Entry_EntryType", Entry_EntryType_name, Entry_EntryType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for RaftNetwork service

type RaftNetworkClient interface {
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	ClientToLeaderRequest(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	InstallSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
}

type raftNetworkClient struct {
	cc *grpc.ClientConn
}

func NewRaftNetworkClient(cc *grpc.ClientConn) RaftNetworkClient {
	return &raftNetworkClient{cc}
}

func (c *raftNetworkClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) ClientToLeaderRequest(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/ClientToLeaderRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) InstallSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/InstallSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftNetwork service

type RaftNetworkServer interface {
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	ClientToLeaderRequest(context.Context, *EntryRequest) (*EmptyMessage, error)
	InstallSnapshot(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
}

func RegisterRaftNetworkServer(s *grpc.Server, srv RaftNetworkServer) {
	s.RegisterService(&_RaftNetwork_serviceDesc, srv)
}

func _RaftNetwork_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftNetworkServer).AppendEntries(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RaftNetwork_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftNetworkServer).RequestVote(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RaftNetwork_ClientToLeaderRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftNetworkServer).ClientToLeaderRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RaftNetwork_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RaftNetworkServer).InstallSnapshot(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RaftNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftNetwork",
	HandlerType: (*RaftNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _RaftNetwork_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftNetwork_RequestVote_Handler,
		},
		{
			MethodName: "ClientToLeaderRequest",
			Handler:    _RaftNetwork_ClientToLeaderRequest_Handler,
		},
		{
			MethodName: "InstallSnapshot",
			Handler:    _RaftNetwork_InstallSnapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 1065 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x5d, 0x6e, 0x23, 0x45,
	0x10, 0x5e, 0xff, 0xdb, 0x35, 0x71, 0xe2, 0xb4, 0x37, 0xcb, 0x24, 0x2b, 0x20, 0x3b, 0x10, 0x7e,
	0x04, 0x04, 0xc9, 0xcb, 0x0b, 0x8f, 0x21, 0xcb, 0x42, 0x20, 0x6b, 0x56, 0x93, 0x88, 0x27, 0x24,
	0x6b, 0x76, 0xa6, 0x6d, 0x8f, 0xd6, 0xee, 0x1e, 0x66, 0x7a, 0x36, 0xeb, 0x03, 0x70, 0x04, 0xae,
	0xc0, 0x09, 0xb8, 0x0e, 0x12, 0x12, 0x87, 0xe0, 0x95, 0xee, 0xaa, 0x1e, 0x7b, 0xfc, 0xc3, 0x0a,
	0xf1, 0x62, 0x75, 0x7f, 0xf5, 0xcd, 0xd7, 0x55, 0xd5, 0x55, 0x5d, 0x86, 0x83, 0x34, 0x18, 0xab,
	0xcf, 0xcd, 0xcf, 0x79, 0x92, 0x4a, 0x25, 0x59, 0xdd, 0xac, 0xbd, 0x7d, 0xd8, 0xfb, 0x7a, 0x9e,
	0xa8, 0xc5, 0x33, 0x9e, 0x65, 0xc1, 0x84, 0x7b, 0x43, 0xbd, 0x17, 0x2a, 0x5d, 0xf8, 0xfc, 0xe7,
	0x9c, 0x67, 0x8a, 0x3d, 0x84, 0x4e, 0xc6, 0x45, 0xc4, 0xd3, 0x51, 0x1c, 0xb9, 0x95, 0xd3, 0xca,
	0x47, 0x1d, 0xbf, 0x4d, 0xc0, 0x55, 0xc4, 0x1e, 0x41, 0x83, 0x1b, 0xb2, 0x5b, 0xd5, 0x06, 0x67,
	0xe0, 0x9c, 0xa3, 0x3c, 0x7d, 0x4f, 0x16, 0xef, 0x8f, 0x0a, 0xdc, 0xbf, 0x48, 0x12, 0xfd, 0x85,
	0x81, 0x63, 0x9e, 0x15, 0xc2, 0x0c, 0xea, 0x8a, 0xa7, 0x73, 0xd4, 0xac, 0xfb, 0xb8, 0x36, 0x87,
	0xcd, 0x78, 0x60, 0x0f, 0xab, 0xd2, 0x61, 0x04, 0xe8, 0xc3, 0xde, 0x87, 0xfd, 0x24, 0xe5, 0xaf,
	0x46, 0x33, 0x39, 0x19, 0xc5, 0xda, 0x81, 0xd7, 0x6e, 0x0d, 0x3f, 0xdd, 0x33, 0xe8, 0xb5, 0x9c,
	0x5c, 0x19, 0x8c, 0x79, 0xd0, 0x5d, 0xb2, 0x50, 0xbf, 0x8e, 0x24, 0xc7, 0x92, 0x6e, 0xcd, 0x31,
	0x67, 0xd0, 0xe2, 0xe4, 0x8c, 0xdb, 0x38, 0xad, 0x6d, 0x3a, 0x5e, 0xd8, 0xd8, 0x7b, 0xd0, 0xb5,
	0xde, 0x84, 0x72, 0x3e, 0x8f, 0x95, 0xdb, 0xa4, 0xf3, 0x08, 0xbc, 0x44, 0xcc, 0xfb, 0xbb, 0x0a,
	0xfd, 0x1b, 0x15, 0x28, 0xfe, 0x2c, 0x08, 0xa7, 0xb1, 0xe0, 0x06, 0x0e, 0x44, 0x84, 0xe1, 0x2d,
	0x12, 0x8e, 0xe1, 0x75, 0x7d, 0x5c, 0x1b, 0x2c, 0x09, 0xd4, 0xd4, 0x46, 0x86, 0x6b, 0x83, 0x45,
	0x81, 0x0a, 0x30, 0x96, 0x3d, 0x1f, 0xd7, 0xec, 0x01, 0x34, 0xe5, 0x78, 0x9c, 0x71, 0x85, 0xce,
	0xd7, 0x7c, 0xbb, 0x33, 0xf8, 0x8c, 0x8b, 0x89, 0x56, 0x68, 0x10, 0x4e, 0x3b, 0x76, 0x0c, 0x6d,
	0x39, 0x8b, 0x46, 0xa8, 0xdd, 0x44, 0xed, 0x96, 0xde, 0x3f, 0x0f, 0xc8, 0x24, 0xf8, 0x1d, 0x99,
	0x5a, 0x64, 0xd2, 0xfb, 0xe7, 0xf6, 0xe4, 0xb9, 0x8c, 0xb8, 0xdb, 0x26, 0x0f, 0xcd, 0x5a, 0x67,
	0x66, 0x3f, 0x08, 0x43, 0x5d, 0x0a, 0xa3, 0x8c, 0x87, 0x52, 0x44, 0x99, 0xdb, 0xc1, 0x93, 0xba,
	0x84, 0xde, 0x10, 0xc8, 0x3e, 0x03, 0x66, 0x69, 0x22, 0x10, 0xb2, 0xa0, 0x02, 0x52, 0x0f, 0xc9,
	0x32, 0x5c, 0x19, 0x8c, 0xaa, 0x56, 0x8f, 0xc7, 0x8b, 0xa5, 0xaa, 0x43, 0xaa, 0x84, 0x96, 0x54,
	0x2d, 0xad, 0xac, 0xba, 0x47, 0xaa, 0x64, 0x29, 0xa9, 0x7a, 0xbf, 0x54, 0xe0, 0xe0, 0x7b, 0xbe,
	0xa0, 0xe4, 0x53, 0xf5, 0xb2, 0x0f, 0xa1, 0xf3, 0x92, 0x2f, 0x46, 0xf2, 0x4e, 0xf0, 0x14, 0x53,
	0xef, 0x0c, 0x80, 0xee, 0x76, 0xa8, 0xc3, 0xf3, 0xdb, 0xda, 0xf8, 0x83, 0xb1, 0xb1, 0x8f, 0x01,
	0x0c, 0x71, 0xaa, 0xf3, 0xa4, 0x99, 0xd5, 0x2d, 0xa6, 0x91, 0xf9, 0x16, 0x8d, 0xec, 0x1d, 0x80,
	0x09, 0xd7, 0xdf, 0x04, 0x2a, 0x96, 0x02, 0xef, 0xa9, 0xe6, 0x97, 0x10, 0x2f, 0x82, 0xba, 0xf9,
	0x84, 0xed, 0x43, 0x35, 0x4e, 0x6c, 0x8b, 0xe8, 0x15, 0xde, 0xb6, 0x4c, 0xd5, 0xf2, 0xb6, 0xf5,
	0x9a, 0xbd, 0x0b, 0x8e, 0xa9, 0x25, 0x29, 0x74, 0x88, 0x73, 0x8e, 0x62, 0x1d, 0x1f, 0x08, 0x1a,
	0x6a, 0x84, 0xbd, 0x05, 0x2d, 0xa1, 0xc5, 0x4c, 0xfd, 0xd7, 0xd1, 0xd8, 0x34, 0xdb, 0xab, 0xc8,
	0xfb, 0xbd, 0x02, 0xdd, 0x4b, 0x29, 0xc6, 0xf1, 0x24, 0xa7, 0x73, 0xd9, 0x97, 0xa5, 0x0a, 0xdb,
	0x1f, 0x9c, 0x91, 0xf3, 0x6b, 0x94, 0xf5, 0xdd, 0xad, 0x26, 0xdb, 0x42, 0x3c, 0x85, 0x86, 0x91,
	0xcd, 0xb4, 0x6f, 0xb5, 0x8d, 0xc0, 0xc9, 0xe0, 0x3d, 0x85, 0xc3, 0xad, 0x8f, 0x99, 0x0b, 0xf7,
	0x2f, 0xf3, 0x34, 0xd5, 0xfd, 0xb1, 0x66, 0xeb, 0xdd, 0xd3, 0x6e, 0xf7, 0x9f, 0xe6, 0x2a, 0x4f,
	0xf9, 0xba, 0xa1, 0xe2, 0x9d, 0x81, 0xf3, 0x84, 0xcf, 0x65, 0xd1, 0x15, 0xba, 0x82, 0x45, 0x3e,
	0x7f, 0x61, 0x2f, 0xa7, 0xee, 0xdb, 0x9d, 0xf7, 0x57, 0x15, 0x1a, 0xd8, 0x7d, 0xfa, 0x62, 0xca,
	0x51, 0x1d, 0x95, 0x1a, 0x93, 0x7e, 0x4b, 0x51, 0xe8, 0x04, 0xe7, 0xf9, 0xf2, 0xa1, 0xc0, 0x35,
	0x7b, 0x0c, 0xad, 0x90, 0xce, 0xc2, 0xe4, 0x3a, 0x83, 0x63, 0x52, 0xd8, 0xd1, 0xa2, 0x7e, 0xc1,
	0x64, 0x9f, 0x40, 0x33, 0x44, 0xbf, 0x31, 0xe7, 0xce, 0xa0, 0xbf, 0x23, 0x97, 0xbe, 0xa5, 0xe8,
	0x62, 0xae, 0x47, 0x3a, 0x22, 0x6c, 0x41, 0x67, 0x70, 0x48, 0xd4, 0x52, 0x8c, 0x3e, 0x9a, 0xd9,
	0x17, 0x54, 0x60, 0xe1, 0x34, 0x10, 0x13, 0x8e, 0x5d, 0xe9, 0x14, 0xd1, 0x6c, 0x14, 0x2d, 0xd6,
	0xda, 0x25, 0xf2, 0xbc, 0x9f, 0xa0, 0xb3, 0x8c, 0xd2, 0x24, 0x75, 0x87, 0xdb, 0x94, 0xed, 0x35,
	0xdf, 0xe8, 0xe3, 0x5e, 0x85, 0xb5, 0xa1, 0x6e, 0x3c, 0xe9, 0x55, 0x59, 0x7f, 0xab, 0x37, 0x7a,
	0x35, 0xef, 0x02, 0xda, 0xfa, 0x09, 0xa4, 0x3c, 0xef, 0x7a, 0x7e, 0xff, 0xc3, 0x73, 0x1e, 0xc1,
	0xd1, 0xc6, 0x6b, 0x9e, 0x25, 0x52, 0x64, 0x7c, 0xa7, 0xde, 0xdb, 0x00, 0x82, 0xbf, 0x56, 0xf6,
	0xb5, 0xae, 0xa2, 0xa5, 0x63, 0x10, 0x7a, 0xaa, 0x5d, 0x68, 0x65, 0x39, 0x3e, 0x16, 0x78, 0x57,
	0x6d, 0xbf, 0xd8, 0x7a, 0xbf, 0x56, 0x80, 0xd9, 0x39, 0xf1, 0xa3, 0x54, 0xfc, 0x4d, 0x23, 0xe3,
	0x11, 0xec, 0x85, 0x3a, 0x2b, 0xb1, 0x7e, 0x38, 0xf9, 0x6a, 0x6a, 0x38, 0x4b, 0x8c, 0x06, 0xc7,
	0x2c, 0xc8, 0xd4, 0xf6, 0xe0, 0x30, 0x68, 0x79, 0x70, 0x2c, 0x59, 0xe5, 0xc1, 0x61, 0x49, 0x66,
	0x70, 0x78, 0xd7, 0xd0, 0x5f, 0x73, 0xeb, 0x0d, 0xb1, 0x6b, 0xbf, 0x5e, 0x69, 0xce, 0x68, 0x92,
	0x06, 0x42, 0x71, 0xf2, 0xab, 0xed, 0x3b, 0x06, 0xfb, 0x86, 0x20, 0xef, 0x4f, 0xfd, 0x80, 0xdd,
	0x88, 0x20, 0xc9, 0xa6, 0x52, 0xfd, 0xef, 0xa9, 0x78, 0x0e, 0x7d, 0x74, 0x3b, 0x16, 0xe1, 0x2c,
	0x8f, 0x78, 0xb4, 0x16, 0xe1, 0xa1, 0x31, 0x5d, 0x59, 0x0b, 0x85, 0xf9, 0x29, 0xb0, 0x75, 0x7e,
	0x29, 0xd6, 0x5e, 0x99, 0x8e, 0x93, 0x72, 0x35, 0x89, 0x1a, 0xd4, 0xaf, 0x76, 0x12, 0x15, 0x53,
	0xab, 0x59, 0x9a, 0x5a, 0x06, 0x93, 0x82, 0xe3, 0x98, 0x69, 0xfb, 0xb8, 0xf6, 0x3e, 0x80, 0xde,
	0x2a, 0xc2, 0x7f, 0xcf, 0xd6, 0xe0, 0xb7, 0x2a, 0x38, 0xbe, 0x2e, 0xb6, 0x21, 0x57, 0x77, 0x32,
	0x7d, 0xc9, 0xbe, 0x83, 0xee, 0x5a, 0x99, 0xb1, 0x13, 0xaa, 0xc5, 0x5d, 0xff, 0x24, 0x4e, 0x1e,
	0xee, 0xb4, 0xd1, 0x69, 0xde, 0x3d, 0xf6, 0x44, 0x4b, 0xaf, 0x2e, 0x8d, 0xb9, 0xc4, 0xde, 0x2e,
	0xaf, 0x93, 0xe3, 0x1d, 0x96, 0xa5, 0xca, 0x05, 0x1c, 0x5d, 0xce, 0x62, 0xdd, 0x04, 0xb7, 0xf2,
	0x1a, 0x73, 0xbf, 0xbc, 0xb1, 0x72, 0x97, 0x58, 0xa5, 0x02, 0x2b, 0xff, 0xb1, 0xba, 0xc7, 0xbe,
	0x82, 0x83, 0x2b, 0x91, 0xa9, 0x60, 0x36, 0x2b, 0x72, 0xc2, 0xec, 0x8b, 0xb0, 0x51, 0x05, 0x27,
	0x0f, 0x36, 0xe1, 0xc2, 0x8d, 0x17, 0x4d, 0xfc, 0xef, 0xf6, 0xf8, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x87, 0x23, 0xdc, 0xce, 0x09, 0x00, 0x00,
}
