// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package go_micro_srv_attachment_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AttachmentService service

func NewAttachmentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttachmentService service

type AttachmentService interface {
	// 根据attachment_id获取attachment详情
	AttachmentDetailByIds(ctx context.Context, in *AttachmentDetailByIdsReq, opts ...client.CallOption) (*AttachmentDetailByIdsRep, error)
	// 增加附件
	AddAttachment(ctx context.Context, in *AddAttachmentReq, opts ...client.CallOption) (*AddAttachmentReqRep, error)
}

type attachmentService struct {
	c    client.Client
	name string
}

func NewAttachmentService(name string, c client.Client) AttachmentService {
	return &attachmentService{
		c:    c,
		name: name,
	}
}

func (c *attachmentService) AttachmentDetailByIds(ctx context.Context, in *AttachmentDetailByIdsReq, opts ...client.CallOption) (*AttachmentDetailByIdsRep, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.AttachmentDetailByIds", in)
	out := new(AttachmentDetailByIdsRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentService) AddAttachment(ctx context.Context, in *AddAttachmentReq, opts ...client.CallOption) (*AddAttachmentReqRep, error) {
	req := c.c.NewRequest(c.name, "AttachmentService.AddAttachment", in)
	out := new(AddAttachmentReqRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttachmentService service

type AttachmentServiceHandler interface {
	// 根据attachment_id获取attachment详情
	AttachmentDetailByIds(context.Context, *AttachmentDetailByIdsReq, *AttachmentDetailByIdsRep) error
	// 增加附件
	AddAttachment(context.Context, *AddAttachmentReq, *AddAttachmentReqRep) error
}

func RegisterAttachmentServiceHandler(s server.Server, hdlr AttachmentServiceHandler, opts ...server.HandlerOption) error {
	type attachmentService interface {
		AttachmentDetailByIds(ctx context.Context, in *AttachmentDetailByIdsReq, out *AttachmentDetailByIdsRep) error
		AddAttachment(ctx context.Context, in *AddAttachmentReq, out *AddAttachmentReqRep) error
	}
	type AttachmentService struct {
		attachmentService
	}
	h := &attachmentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttachmentService{h}, opts...))
}

type attachmentServiceHandler struct {
	AttachmentServiceHandler
}

func (h *attachmentServiceHandler) AttachmentDetailByIds(ctx context.Context, in *AttachmentDetailByIdsReq, out *AttachmentDetailByIdsRep) error {
	return h.AttachmentServiceHandler.AttachmentDetailByIds(ctx, in, out)
}

func (h *attachmentServiceHandler) AddAttachment(ctx context.Context, in *AddAttachmentReq, out *AddAttachmentReqRep) error {
	return h.AttachmentServiceHandler.AddAttachment(ctx, in, out)
}
